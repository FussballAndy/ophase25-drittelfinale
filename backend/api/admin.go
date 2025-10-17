package api

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"ophase25/gelaendespiel/consts"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var ADMIN_TOKEN = func() string {
	if tok, ok := os.LookupEnv("ADMIN_TOKEN"); ok {
		return tok
	} else {
		return "admin"
	}
}()

var ResultsChannel = make(chan JSONStationResult, 5)

func HandleAdmin(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Admin: %s\n", ADMIN_TOKEN)
	fmt.Printf("%+v\n", r.Cookies())
	cookie := r.URL.Query().Get("session")
	if cookie == "" {
		WriteError(w)
		fmt.Println("No cookie")
		return
	}
	token := cookie
	if token != ADMIN_TOKEN {
		WriteError(w)
		fmt.Println("Wrong cookie")
		return
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	go func() {
		err := adminRoutine(conn)
		if err != nil {
			log.Println(err)
			conn.Close()
		}
	}()
}

var ConfigDone = false

func adminRoutine(conn *websocket.Conn) error {
	err := conn.WriteJSON(JSONAdmin{
		Kind: "num_stations",
		Data: NUM_STATIONS,
	})
	if err != nil {
		return err
	}
	if !ConfigDone {
		err = setupAdmin(conn)
		if err != nil {
			return err
		}
		err = conn.WriteJSON(JSONAdmin{
			Kind: "confirmation",
		})
		if err != nil {
			return err
		}
		var conf any
		err = conn.ReadJSON(&conf)
		if err != nil {
			return err
		}
		StoreDB()
		ConfigDone = true
	} else {
		type JSONFullDump struct {
			Tokens    map[string]uint8 `json:"tokens"`
			Stations  []Station        `json:"stations"`
			Groups    []Group          `json:"groups"`
			Questions []Question       `json:"questions"`
		}
		err = conn.WriteJSON(JSONAdmin{
			Kind: "fulldump",
			Data: JSONFullDump{
				Tokens:    DBTokens,
				Stations:  DBStations,
				Groups:    DBGroups,
				Questions: DBQuestions,
			},
		})
		if err != nil {
			return err
		}
	}

	err = conn.WriteJSON(JSONAdmin{
		Kind: "ingame",
		Data: DBScores,
	})
	if err != nil {
		return err
	}

	ingameDur := time.Until(consts.INGAME_END)
	if ingameDur > 0 {
		ingameEndTimer := time.NewTimer(ingameDur)
	ingameLoop:
		for {
			select {
			case result := <-ResultsChannel:
				err = conn.WriteJSON(JSONAdmin{
					Kind: "result",
					Data: result,
				})
				if err != nil {
					return err
				}
				fmt.Println("Sent one result")
			case <-ingameEndTimer.C:
				break ingameLoop
			}
		}
	}

	err = conn.WriteJSON(JSONAdmin{
		Kind: "final",
		Data: DBScores,
	})
	if err != nil {
		return err
	}

	//if time.Until(consts.DRITTEL_END) >= -time.Minute*30 {
	if true {
		finished := false

		for !finished {
			var msg JSONAdmin
			err = conn.ReadJSON(&msg)
			if err != nil {
				return err
			}
			switch msg.Kind {
			case "next":
				err = handleNextQuestion(conn, uint8(msg.Data.(float64)))
				if err != nil {
					return err
				}
			case "clear":
				clearQuestion()
			case "finished":
				finished = true
			case "front":
				clientR, ok := CookieMap.Load(msg.Data)
				if ok {
					client := clientR.(*DrittelClient)
					client.Front = !client.Front
				}
				err = conn.WriteJSON(JSONAdmin{
					Kind: "front",
					Data: ok,
				})
				if err != nil {
					return err
				}
			}
		}
	}

	err = sendResults(conn)
	if err != nil {
		return err
	}

	return nil
}

func setupAdmin(conn *websocket.Conn) error {
	var err error
	err = conn.WriteJSON(JSONAdmin{
		Kind: "stations",
		Data: DBStations,
	})
	if err != nil {
		return err
	}
	var stationSlice []Station
	err = conn.ReadJSON(&stationSlice)
	if err != nil {
		return err
	}
	DBStations = stationSlice

	err = conn.WriteJSON(JSONAdmin{
		Kind: "groups",
		Data: DBGroups,
	})
	if err != nil {
		return err
	}
	var groupSlice []Group
	err = conn.ReadJSON(&groupSlice)
	if err != nil {
		return err
	}
	DBGroups = groupSlice

	if !TokensInit {
		for k := range DBTokens {
			delete(DBTokens, k)
		}
		for k := range NUM_STATIONS {
			for range 4 {
				token := genCookie()
				for tokenExists(token) {
					token = genCookie()
				}
				DBTokens[token] = uint8(k)
			}
		}
		for range 100 {
			token := genCookie()
			for tokenExists(token) {
				token = genCookie()
			}
			DBTokens[token] = math.MaxUint8
		}
	}
	err = conn.WriteJSON(JSONAdmin{
		Kind: "tokens",
		Data: DBTokens,
	})
	if err != nil {
		return err
	}
	var empty any
	err = conn.ReadJSON(&empty)
	if err != nil {
		return err
	}
	TokensInit = true

	err = conn.WriteJSON(JSONAdmin{
		Kind: "questions",
		Data: DBQuestions,
	})
	if err != nil {
		return err
	}
	var questionSlice []Question
	err = conn.ReadJSON(&questionSlice)
	if err != nil {
		return err
	}
	DBQuestions = questionSlice
	fmt.Println("Finished setup")

	return nil
}

func handleNextQuestion(conn *websocket.Conn, nextQuestion uint8) error {
	questionsLen := len(DBQuestions)
	CurrentMutex.Lock()
	if nextQuestion >= uint8(questionsLen) {
		return conn.WriteJSON(JSONAdmin{
			Kind: "end",
		})
	}

	rawQuestion := DBQuestions[nextQuestion]

	BroadcastClients(JSONQuestion{
		Number:  nextQuestion,
		Prompt:  rawQuestion.Prompt,
		Answers: rawQuestion.Answers,
	})

	CurrentQuestion = nextQuestion
	CurrentSubmissionStart = time.Now()

	CurrentMutex.Unlock()

	go func() {
		timer := time.NewTimer(time.Second * 31)
		<-timer.C
		submissions := make(map[string]DrittelAnswer)
		SubmissionMap.Range(func(key, value any) bool {
			key1 := key.(DrittelSub)
			v1 := value.(DrittelAnswer)
			if key1.Question == nextQuestion {
				submissions[key1.Session] = v1
			}
			return true
		})
		conn.WriteJSON(JSONAdmin{
			Kind: "submissions",
			Data: submissions,
		})
		clearQuestion()
	}()

	return nil
}

func clearQuestion() {
	CurrentMutex.Lock()

	CurrentQuestion = 255

	BroadcastClients(JSONQuestion{
		Number:  255,
		Prompt:  "",
		Answers: nil,
	})

	CurrentMutex.Unlock()
}

func sendResults(conn *websocket.Conn) error {
	err := conn.WriteJSON(JSONAdmin{
		Kind: "scores",
		Data: DBScores,
	})
	if err != nil {
		return err
	}

	type ResultDrittel struct {
		Question uint8 `json:"question"`
		Answer   uint8 `json:"answer"`
		Group    bool  `json:"group"`
		Front    bool  `json:"front"`
	}

	resultMap := make([]ResultDrittel, 200)
	SubmissionMap.Range(func(key, value any) bool {
		k1 := key.(DrittelSub)
		v1 := value.(DrittelAnswer)
		resultMap = append(resultMap, ResultDrittel{
			Question: k1.Question,
			Answer:   v1.Answer,
			Group:    v1.Group,
			Front:    v1.Front,
		})
		return true
	})

	err = conn.WriteJSON(JSONAdmin{
		Kind: "drittel",
		Data: resultMap,
	})
	if err != nil {
		return err
	}

	return nil
}

func tokenExists(token string) bool {
	_, ok := DBTokens[token]
	return ok
}

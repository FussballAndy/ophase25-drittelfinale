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
		ConfigDone = true
	}

	err = conn.WriteJSON(JSONAdmin{
		Kind: "ingame",
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

	if time.Until(consts.DRITTEL_END) < -time.Minute*30 {
		return nil
	}

	finished := false

	for !finished {
		var msg JSONAdmin
		err = conn.ReadJSON(&msg)
		if err != nil {
			return err
		}
		switch msg.Kind {
		case "next":
			err = handleNextQuestion(conn)
			if err != nil {
				return err
			}
		case "clear":
			clearQuestion()
		case "finished":
			err = sendResults(conn)
			if err != nil {
				return err
			}
		case "front":
			clientR, ok := CookieMap.Load(msg.Data)
			client := clientR.(*DrittelClient)
			if ok {
				client.Front = !client.Front
			}
		}
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
	var questionSlice []JSONQuestion
	err = conn.ReadJSON(&questionSlice)
	if err != nil {
		return err
	}
	DBQuestions = questionSlice
	fmt.Println("Finished setup")

	return nil
}

func handleNextQuestion(conn *websocket.Conn) error {
	questionsLen := len(DBQuestions)
	CurrentMutex.Lock()
	if CurrentQuestion >= uint8(questionsLen) {
		return conn.WriteJSON(JSONAdmin{
			Kind: "end",
		})
	}

	BroadcastClients(DBQuestions[CurrentQuestion])

	CurrentQuestion++

	CurrentMutex.Unlock()

	return nil
}

func clearQuestion() {
	CurrentMutex.Lock()
	if CurrentQuestion > 0 {
		CurrentQuestion--
	}

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

	resultMap := make(map[DrittelSub]uint8, 200)
	SubmissionMap.Range(func(key, value any) bool {
		resultMap[key.(DrittelSub)] = value.(uint8)
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

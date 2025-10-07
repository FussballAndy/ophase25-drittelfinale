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

var ADMIN_TOKEN = os.Getenv("ADMIN_SESSION")
var ResultsChannel = make(chan JSONStationResult, 5)

func HandleAdmin(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		WriteError(w)
		return
	}
	token := cookie.Value
	if token != ADMIN_TOKEN {
		WriteError(w)
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

func adminRoutine(conn *websocket.Conn) error {
	err := conn.WriteJSON(JSONAdmin{
		Kind: "num_iterations",
		Data: NUM_ITERATIONS,
	})
	if err != nil {
		return err
	}
	finished := false
	for !finished {
		err = setupAdmin(conn)
		if err != nil {
			return err
		}
		err = conn.WriteJSON(JSONAdmin{
			Kind: MSG_REQUEST,
			Data: "confirmation",
		})
		if err != nil {
			return err
		}
		var response JSONAdmin
		err = conn.ReadJSON(&response)
		if err != nil {
			return err
		}
		if response.Kind == "confirm" {
			finished = true
		} else {
			switch response.Data.(string) {
			case "groups":
				GroupsInit = false
			case "stations":
				StationsInit = false
			case "questions":
				QuestionsInit = false
			default:
				fmt.Printf("Unknown type: %s\n", response.Data)
			}
		}
	}

	ingameDur := time.Until(consts.INGAME_END)
	if ingameDur > 0 {
		ingameEndTimer := time.NewTimer(ingameDur)
	ingameLoop:
		for {
			select {
			case result := <-ResultsChannel:
				conn.WriteJSON(JSONAdmin{
					Kind: "result",
					Data: result,
				})
			case <-ingameEndTimer.C:
				break ingameLoop
			}
		}
	}

	err = drittelConfig(conn)
	if err != nil {
		return err
	}

	if time.Until(consts.DRITTEL_END) < 0 {
		return nil
	}

	finished = false

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
		}
	}

	return nil
}

func setupAdmin(conn *websocket.Conn) error {
	var err error
	if !StationsInit {
		err = conn.WriteJSON(JSONAdmin{
			Kind: MSG_REQUEST,
			Data: "stations",
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
		StationsInit = true
	} else {
		err = conn.WriteJSON(JSONAdmin{
			Kind: "stations",
			Data: DBStations,
		})
		if err != nil {
			return err
		}
	}
	if !GroupsInit {
		err = conn.WriteJSON(JSONAdmin{
			Kind: MSG_REQUEST,
			Data: "groups",
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
		GroupsInit = true
	} else {
		err = conn.WriteJSON(JSONAdmin{
			Kind: "groups",
			Data: DBGroups,
		})
		if err != nil {
			return err
		}
	}
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
		err = conn.WriteJSON(JSONAdmin{
			Kind: "tokens",
			Data: DBTokens,
		})
		if err != nil {
			return err
		}
	}
	if !QuestionsInit {
		err = conn.WriteJSON(JSONAdmin{
			Kind: MSG_REQUEST,
			Data: "questions",
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
		QuestionsInit = true
	} else {
		err = conn.WriteJSON(JSONAdmin{
			Kind: "questions",
			Data: DBQuestions,
		})
		if err != nil {
			return err
		}
	}

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

func drittelConfig(conn *websocket.Conn) error {
	err := conn.WriteJSON(JSONAdmin{
		Kind: "drittel_config",
	})
	if err != nil {
		return err
	}
	var config JSONDrittelConfig
	err = conn.ReadJSON(&config)
	if err != nil {
		return err
	}
	return nil
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

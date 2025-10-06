package api

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type DrittelClient struct {
	Group   bool
	Front   bool
	Conn    *websocket.Conn
	Session string
}

var CookieMap sync.Map
var ClientMutex sync.Mutex

var SubmissionMap sync.Map

var CurrentSubmissionStart int64
var CurrentQuestion uint8 = 0
var CurrentMutex sync.RWMutex

const NUM_SUBMISSIONS = 20

func HandleDrittel(w http.ResponseWriter, r *http.Request) {
	client := &DrittelClient{
		Group: false,
		Front: false,
		Conn:  nil,
	}
	var sessionToken string
	cookie, err := r.Cookie("session")
	if err != nil {
		sessionToken = genCookie()
	} else {
		sessionToken = cookie.Value
		client2, ok := CookieMap.Load(sessionToken)
		if !ok {
			if _, ok := DBTokens[sessionToken]; ok {
				client.Group = true
			} else {
				sessionToken = genCookie()
			}
		} else {
			client3 := client2.(*DrittelClient)
			if client3.Conn != nil {
				client3.Conn.Close()
			}
			client = client3
		}
	}
	responseHeader := make(http.Header)
	cookie = &http.Cookie{
		Name:     "session",
		Value:    sessionToken,
		HttpOnly: true,
		Secure:   true,
	}
	responseHeader.Add("Set-Cookie", cookie.String())
	conn, err := upgrader.Upgrade(w, r, responseHeader)
	if err != nil {
		log.Println(err)
		return
	}
	client.Conn = conn
	client.Session = sessionToken
	CookieMap.Store(sessionToken, client)
	go client.ReadMessages()

	newQuestion := &JSONQuestion{
		Number:  0,
		Prompt:  "Hallo",
		Answers: []string{"A", "B", "C"},
	}
	conn.WriteJSON(newQuestion)
}

func (c *DrittelClient) ReadMessages() {
	for {
		var answer JSONSubmission
		err := c.Conn.ReadJSON(&answer)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(answer)
		timeNow := time.Now().UnixMilli()
		CurrentMutex.RLock()
		diff := timeNow - CurrentSubmissionStart
		if 0 < diff && diff < 30_500 && answer.Question == CurrentQuestion {
			submissions, ok := SubmissionMap.Load(c.Session)
			if !ok {
				submissions = &[NUM_SUBMISSIONS]uint8{}
				SubmissionMap.Store(c.Session, submissions)
			}
			submissions.(*[NUM_SUBMISSIONS]uint8)[answer.Question] = answer.Answer
			log.Println(submissions)
		}
		CurrentMutex.RUnlock()
	}
}

func genCookie() string {
	bytes := make([]byte, 4)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

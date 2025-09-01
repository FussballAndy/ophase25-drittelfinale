package main

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/quic-go/quic-go/http3"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var userInfo map[string]ClientLevel = make(map[string]ClientLevel)
var clients []*Client = make([]*Client, 0)

func main() {
	userInfo["ADMIN-OPHASE25"] = Admin
	dat, err := os.ReadFile("client.html")
	if err != nil {
		log.Println(err)
		return
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(dat))
	})
	mux.HandleFunc("/ws", HandleWebSocket)
	err = http3.ListenAndServeTLS("0.0.0.0:5000", "cert.pem", "key.pem", mux)
	if err != nil {
		log.Printf("Encountered the following error: %s", err)
	}
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket Upgrade failed for %s: %s", r.RemoteAddr, err)
		return
	}

	var client *Client = nil
	if sess, err := r.Cookie("session"); err == nil {
		sess := sess.Value
		log.Printf("Received cookie: %s", sess)
		level, ok := userInfo[sess]
		if ok {
			client = &Client{
				conn:  conn,
				level: level,
			}
		}
	}
	if client == nil {
		sess, err := GenerateSession()
		if err != nil {
			log.Printf("Failed to generate session: %s", err)
			return
		}
		userInfo[sess] = User
		client = &Client{
			conn:  conn,
			level: User,
		}
		data := JSONNewSession{
			Session: sess,
		}

		if err := conn.WriteJSON(data); err != nil {
			log.Printf("WebSocket send failed: %s", err)
			return
		}
	}

	clients = append(clients, client)

	client.HandleMessages()
}

var SessionLength = 16 // bytes

func GenerateSession() (string, error) {
	bytes := make([]byte, SessionLength)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

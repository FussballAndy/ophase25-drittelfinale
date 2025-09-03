package main

import (
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

var clientMgr *ClientManager = NewClientManager()
var teamMgr *TeamManager = NewTeamManager(25)

func main() {
	dat, err := os.ReadFile("client.html")
	if err != nil {
		log.Println(err)
		return
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(dat))
	})
	mux.HandleFunc("/ws", HandleWebSocket)
	err = http3.ListenAndServeTLS("127.0.0.1:5000", "cert.pem", "key.pem", mux)
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

	sessionToken := ""
	if sess, err := r.Cookie("session"); err == nil {
		sessionToken = sess.Value
	}
	client, err := clientMgr.Create(conn, sessionToken)
	if err != nil {
		log.Printf("Failed to generate session: %s", err)
		conn.Close()
		return
	}

	team := teamMgr.AllocateTeam(client.id)

	announcement := JSONData{
		DataType: TeamAnnouncementType,
		Data:     team,
	}
	if !client.SendMessage(announcement) {
		return
	}

	client.HandleMessages()
}

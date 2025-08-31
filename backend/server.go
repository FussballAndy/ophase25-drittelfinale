package main

import (
	"log"
	"net/http"

	"github.com/FussballAndy/webtransport-go"
	"github.com/quic-go/quic-go/http3"
)

var server *webtransport.Server

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("HTTP Version: %s", r.Proto)
		w.Write([]byte("Hello world!"))
	})
	http3.ListenAndServeTLS("0.0.0.0:5000", "cert.pem", "key.pem", mux)
	// mux.HandleFunc("/wt", HandleWebTransport)

	/* var err error
	server, err = CreateServer(":5000", "cert.pem", "key.pem", mux)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ListenAndServeTLS(server, "cert.pem", "key.pem")
	if err != nil {
		fmt.Println(err)
		return
	} */
}

func HandleWebTransport(w http.ResponseWriter, r *http.Request) {
	sess, err := server.Upgrade(w, r)
	if err != nil {
		log.Printf("upgrading failed: %s", err)
		w.WriteHeader(500)
		return
	}
	_ = sess
}

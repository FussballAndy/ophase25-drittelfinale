package main

import (
	"log"
	"net/http"
	"ophase25/gelaendespiel/api"
)

func main() {
	api.DBTokens["A"] = 0
	mux := http.NewServeMux()

	mux.Handle("/api/", GetAPIHandler())
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("true")) })

	s := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	log.Fatal(s.ListenAndServe())
}

func GetAPIHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/token", api.HandleToken)
	mux.HandleFunc("POST /api/winner", api.HandleWinner)
	return CorsMiddleware(mux)
}

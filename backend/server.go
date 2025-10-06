package main

import (
	"log"
	"net/http"
	"ophase25/gelaendespiel/api"
)

func main() {
	api.DBTokens["A"] = 0
	api.SetupResultCreator()
	mux := http.NewServeMux()

	mux.Handle("/api/", GetAPIHandler())
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("true")) })
	mux.HandleFunc("/results", api.HandleResults)

	s := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	log.Printf("Server starting at http://%s\n", s.Addr)
	log.Fatal(s.ListenAndServe())
}

func GetAPIHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/token", api.HandleToken)
	mux.HandleFunc("POST /api/winner", api.HandleWinner)
	mux.HandleFunc("GET /api/stations", api.HandleStations)
	mux.HandleFunc("GET /api/groups", api.HandleGroups)
	mux.HandleFunc("/api/drittel", api.HandleDrittel)
	return CorsMiddleware(mux)
}

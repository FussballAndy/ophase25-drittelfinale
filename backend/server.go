package main

import (
	"log"
	"net/http"
	"ophase25/gelaendespiel/api"
)

func main() {
	api.PopulateDB()
	api.SetupResultCreator()
	mux := http.NewServeMux()

	mux.Handle("/api/", GetAPIHandler())
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./assets/index.html")
	})
	mux.HandleFunc("/index.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./assets/index.css")
	})
	mux.HandleFunc("/index.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./assets/index.js")
	})
	mux.HandleFunc("/logo.png", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./assets/logo.png")
	})
	mux.HandleFunc("/results", api.HandleResults)

	s := &http.Server{
		Addr:    ":8080",
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
	mux.HandleFunc("/api/admin", api.HandleAdmin)
	return CorsMiddleware(mux)
}

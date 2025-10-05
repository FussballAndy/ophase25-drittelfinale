package api

import (
	"encoding/json"
	"net/http"
)

func HandleStations(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(DBStations)
}

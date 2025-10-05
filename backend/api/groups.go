package api

import (
	"encoding/json"
	"net/http"
)

func HandleGroups(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(DBGroups)
}

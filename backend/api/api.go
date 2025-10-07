package api

import (
	"encoding/json"
	"net/http"
)

func WriteError(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(JSONData{Status: false})
}

func WriteOkEmpty(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(JSONData{Status: true})
}

func WriteOkData(w http.ResponseWriter, data any) {
	json.NewEncoder(w).Encode(JSONData{Status: true, Data: data})
}

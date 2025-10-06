package api

import (
	"io"
	"log"
	"net/http"
)

func HandleToken(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	body_raw, err := io.ReadAll(r.Body)
	if err != nil {
		WriteError(w)
		return
	}
	log.Printf("Token: %s\n", body_raw)

	_, ok := DBTokens[string(body_raw)]

	if !ok {
		WriteError(w)
		return
	}

	cookie := http.Cookie{
		Name:     "session",
		Value:    string(body_raw),
		HttpOnly: true,
		Secure:   true,
	}

	http.SetCookie(w, &cookie)

	WriteOkEmpty(w)
}

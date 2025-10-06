package api

import (
	"encoding/json"
	"net/http"
)

func HandleWinner(w http.ResponseWriter, r *http.Request) {
	var req JSONWinner
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		WriteError(w)
		return
	}

	st, ok := DBTokens[req.Token]
	if !ok {
		WriteError(w)
		return
	}

	if req.Iteration >= 3 {
		WriteError(w)
		return
	}

	index := st*NUM_ITERATIONS + req.Iteration

	if DBScores[index] != SCORE_UNSET {
		WriteError(w)
		return
	}

	score := SCORE_STUDENT

	if req.Score {
		score = SCORE_TUTOR
	}

	DBScores[index] = score

	ResultsDirty.Store(true)

	WriteOkEmpty(w)
}

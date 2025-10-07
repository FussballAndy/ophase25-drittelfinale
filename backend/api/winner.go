package api

import (
	"encoding/json"
	"fmt"
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

	if req.Iteration >= NUM_SCORES {
		WriteError(w)
		return
	}

	scorePtr := GetScorePtr(st, req.Iteration)

	if *scorePtr != SCORE_UNSET {
		WriteError(w)
		return
	}

	*scorePtr = req.Score

	ResultsChannel <- JSONStationResult{
		Station:   st,
		Iteration: req.Iteration,
		Winner:    req.Score,
	}

	fmt.Println(DBScores)

	ResultsDirty.Store(true)

	WriteOkEmpty(w)
}

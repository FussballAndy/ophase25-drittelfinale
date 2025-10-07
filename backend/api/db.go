package api

import "math"

type Station struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

type Group struct {
	Stations [NUM_ITERATIONS]uint8 `json:"stations"`
}

// --- CONSTS ---
const NUM_STATIONS = 25
const NUM_ITERATIONS = 3
const NUM_SCORES = NUM_ITERATIONS + 1

// --- Static Info ---
var TokensInit = false
var DBTokens = make(map[string]uint8)
var DBStations = make([]Station, NUM_STATIONS)
var DBGroups = make([]Group, NUM_STATIONS)
var DBQuestions = make([]JSONQuestion, 0)

// --- Scores ---
const SCORE_UNSET uint8 = math.MaxUint8
const SCORE_STUDENT uint8 = 0
const SCORE_TUTOR uint8 = 1

// DBScores[station*NUM_IT + it]
var DBScores = makeScoreSlice()

func makeScoreSlice() []uint8 {
	tmp := make([]uint8, NUM_STATIONS*NUM_SCORES)
	for i := range tmp {
		tmp[i] = SCORE_UNSET
	}
	return tmp
}

func GetScorePtr(station uint8, iteration uint8) *uint8 {
	return &DBScores[station*NUM_SCORES+iteration]
}

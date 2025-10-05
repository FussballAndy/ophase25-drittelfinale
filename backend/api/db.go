package api

type Station struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

type Group struct {
	Stations [NUM_ITERATIONS]uint8 `json:"stations"`
}

const NUM_STATIONS = 25
const NUM_ITERATIONS = 3

var DBTokens = make(map[string]uint8)
var DBStations = make([]Station, NUM_STATIONS)

const SCORE_UNSET uint8 = 0
const SCORE_STUDENT uint8 = 1
const SCORE_TUTOR uint8 = 2

// DBScores[station*NUM_IT + it]
var DBScores = make([]uint8, NUM_STATIONS*NUM_ITERATIONS)

var DBGroups = make([]Group, NUM_STATIONS)

package api

type JSONData struct {
	Status bool `json:"status"`
	Data   any  `json:"data"`
}

const WINNER_STUDENT uint8 = 0
const WINNER_TUTOR uint8 = 1

type JSONWinner struct {
	Token     string `json:"token"`
	Iteration uint8  `json:"iteration"`
	Score     uint8  `json:"score"`
}

type JSONSubmission struct {
	Question uint8 `json:"question"`
	Answer   uint8 `json:"answer"`
}

type JSONQuestion struct {
	Number  uint8    `json:"num"`
	Prompt  string   `json:"prompt"`
	Answers []string `json:"answers"`
}

type AdminMessageKind string

const MSG_REQUEST AdminMessageKind = "request"
const MSG_ERROR AdminMessageKind = "error"
const MSG_SEND AdminMessageKind = "send"

type JSONAdmin struct {
	Kind AdminMessageKind `json:"kind"`
	Data any              `json:"data"`
}

type JSONStationResult struct {
	Station   uint8 `json:"station"`
	Winner    uint8 `json:"result"`
	Iteration uint8 `json:"iter"`
}

type JSONDrittelConfig struct {
	PointScaling uint8 `json:"scaling"`
}

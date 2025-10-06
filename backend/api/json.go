package api

type JSONData struct {
	Status bool `json:"status"`
	Data   any  `json:"data"`
}

type IterationWinner bool

const WINNER_STUDENT IterationWinner = false
const WINNER_TUTOR IterationWinner = true

type JSONWinner struct {
	Token     string          `json:"token"`
	Iteration uint8           `json:"iteration"`
	Score     IterationWinner `json:"score"`
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

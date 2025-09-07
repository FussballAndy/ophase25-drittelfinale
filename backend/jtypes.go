package main

type JSONData struct {
	DataType string `json:"type"`
	Data     any    `json:"data"`
}

type JSONTeams struct {
	Capacity uint16   `json:"capacity"`
	Member   []uint16 `json:"member"`
}

const (
	SessionType           string = "session"
	TeamAnnouncementType  string = "team_announcement"
	TeamSwitchType        string = "team_switch"
	TeamsType             string = "teams"
	ChangePhaseType       string = "change_phase"
	AddGameType           string = "game_add"
	RemoveGameType        string = "game_remove"
	PathAnnouncementType  string = "path"
	FinalQuestionType     string = "final_question"
	FinalAnswerType       string = "final_answer"
	GameWinnerType        string = "game_winner"
	SGameWinnerType       string = "special_game_winner"
	ConnectionUpgradeType string = "upgrade"
)

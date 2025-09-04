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
	SessionType          string = "session"
	TeamAnnouncementType string = "teaman"
	TeamSwitchType       string = "teamsw"
	TeamsType            string = "teams"
)

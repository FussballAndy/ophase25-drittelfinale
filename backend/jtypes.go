package main

type JSONData struct {
	DataType string `json:"type"`
	Data     any    `json:"data"`
}

const (
	SessionType          string = "session"
	TeamAnnouncementType string = "teaman"
	TeamSwitchType       string = "teamsw"
	TeamsType            string = "teams"
)

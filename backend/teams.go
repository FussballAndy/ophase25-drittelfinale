package main

import "sync"

type TeamManager struct {
	Teams    map[uint16]uint8
	Mutex    sync.Mutex
	NumTeams uint8
	NextTeam uint8
}

func NewTeamManager(capacity uint8) *TeamManager {
	teams := make(map[uint16]uint8)
	return &TeamManager{
		Teams:    teams,
		NumTeams: capacity,
		NextTeam: 0,
	}
}

func (t *TeamManager) AllocateTeam(id uint16) uint8 {
	t.Mutex.Lock()
	team := t.NextTeam
	t.NextTeam++
	t.Teams[id] = team
	t.Mutex.Unlock()
	return team
}

func (t *TeamManager) GetTeam(id uint16) uint8 {
	t.Mutex.Lock()
	team := t.Teams[id]
	t.Mutex.Unlock()
	return team
}

func (t *TeamManager) ChangeTeam(id uint16, to uint8) bool {
	t.Mutex.Lock()
	t.Teams[id] = to
	t.Mutex.Unlock()
	return true
}

package main

import (
	"sync"
)

type TeamManager struct {
	TeamMapping  map[uint16]uint8
	TeamMembers  []uint16
	Mutex        sync.Mutex
	NumTeams     uint8
	NextTeam     uint8
	TeamCapacity uint16
}

func NewTeamManager(capacity uint8) *TeamManager {
	teamMap := make(map[uint16]uint8)
	teamMem := make([]uint16, capacity)
	return &TeamManager{
		TeamMapping: teamMap,
		TeamMembers: teamMem,
		NumTeams:    capacity,
		NextTeam:    0,
	}
}

func (t *TeamManager) AllocateTeam(id uint16) uint8 {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	if team, ok := t.TeamMapping[id]; ok {
		return team
	} else {
		team := t.NextTeam
		t.NextTeam = (t.NextTeam + 1) % t.NumTeams
		t.TeamMapping[id] = team
		t.TeamMembers[team]++
		return team
	}
}

func (t *TeamManager) GetTeam(id uint16) uint8 {
	t.Mutex.Lock()
	team := t.TeamMapping[id]
	t.Mutex.Unlock()
	return team
}

func (t *TeamManager) ChangeTeam(id uint16, to uint8) bool {
	t.Mutex.Lock()
	prev := t.TeamMapping[id]
	t.TeamMembers[prev]--
	t.TeamMapping[id] = to
	t.TeamMembers[to]++
	t.Mutex.Unlock()
	return true
}

func (t *TeamManager) GetTeamJSON() JSONTeams {
	return JSONTeams{
		Capacity: t.TeamCapacity,
		Member:   t.TeamMembers,
	}
}

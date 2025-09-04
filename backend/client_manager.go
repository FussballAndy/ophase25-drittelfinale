package main

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"sync"

	"github.com/gorilla/websocket"
)

type SessionInfo struct {
	level ClientLevel
	id    uint16
}

type ClientManager struct {
	Clients  []*Client
	Sessions map[string]SessionInfo
	NextId   uint16
	mu       sync.Mutex
}

func NewClientManager() *ClientManager {
	sessions := make(map[string]SessionInfo)
	sessions["ADMIN-OPHASE25"] = SessionInfo{level: Admin, id: 0}
	return &ClientManager{
		Clients:  make([]*Client, 0),
		Sessions: sessions,
		NextId:   1,
	}
}

func (m *ClientManager) Create(conn *websocket.Conn, sess string) (*Client, error) {
	m.mu.Lock()
	id := m.NextId
	level := User
	if info, ok := m.Sessions[sess]; ok {
		id = info.id
		level = info.level
	} else {
		sess, err := GenerateSession()
		if err != nil {
			return nil, err
		}
		m.Sessions[sess] = SessionInfo{
			level: User,
			id:    id,
		}
		m.NextId++
	}
	client := &Client{
		id:    id,
		level: level,
		conn:  conn,
	}
	m.Clients = append(m.Clients, client)
	m.mu.Unlock()

	data := JSONData{
		DataType: SessionType,
		Data:     sess,
	}

	if !client.SendMessage(data) {
		client.conn.Close()
		return nil, errors.New("failed to send session to client")
	}
	return client, nil
}

func (m *ClientManager) CopyClients() []*Client {
	m.mu.Lock()
	tmp := make([]*Client, len(m.Clients))
	copy(tmp, m.Clients)
	m.mu.Unlock()
	return tmp
}

var SessionLength = 16 // bytes

func GenerateSession() (string, error) {
	bytes := make([]byte, SessionLength)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

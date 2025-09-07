package main

import (
	"log"
	"reflect"

	"github.com/gorilla/websocket"
)

type ClientLevel int

const (
	User ClientLevel = iota
	Tutor
	Admin
)

type Client struct {
	conn  *websocket.Conn
	level ClientLevel
	id    uint16
}

func (c *Client) HandleMessages() {
	for {
		var content JSONData
		err := c.conn.ReadJSON(&content)
		if err != nil {
			log.Printf("WebSocket receive failed: %s", err)
			return
		}
		c.HandleMessage(&content)
	}
}

func (c *Client) SendMessage(v any) bool {
	if err := c.conn.WriteJSON(v); err != nil {
		log.Printf("WebSocket send failed: %s", err)
		return false
	}
	return true
}

func (c *Client) HandleMessage(content *JSONData) {
	switch content.DataType {
	case TeamSwitchType:
		if to, ok := content.Data.(float64); ok {
			teamMgr.ChangeTeam(c.id, uint8(to))
			c.conn.WriteJSON(JSONData{
				DataType: TeamAnnouncementType,
				Data:     teamMgr.GetTeam(c.id),
			})
		} else {
			log.Printf("Wrong data type: %s", reflect.TypeOf(content.Data))
		}
	default:
		log.Printf("Unknown message type: %s", content.DataType)
	}
}

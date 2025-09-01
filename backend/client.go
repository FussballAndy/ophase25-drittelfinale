package main

import (
	"log"

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
}

func (c *Client) HandleMessages() {
	for {
		messageType, p, err := c.conn.ReadMessage()
		if err != nil {
			log.Printf("WebSocket receive failed: %s", err)
			return
		}
		if messageType == websocket.TextMessage {
			msg := string(p)
			_ = msg
		}
	}
}

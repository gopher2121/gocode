package main

import (
	"github.com/gorilla/websocket"
)

type client struct {
	// web socket for the client
	socket *websocket.Conn
	// channel to send messages
	send chan []byte
	// room where the client is
	room *room
}

// reading from socket
func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err != nil {
			break
		} else {
			c.room.forward <- msg
		}

	}
	c.socket.Close()
}

// writing
func (c *client) write() {
	for msg := range c.send {

		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}

// we will need to simulate a room where clients are chatting with one another

package main

import (
	"github.com/gorilla/websocket"
)

type client struct {
	// hold a reference to the websocket that will allow us to communicate with the client
	socket *websocket.Conn

	// send is a buffered channel on which messages are sent and the
	// received messages are queued ready to be forwarded to user's browser
	send chan []byte

	// room will store the information on the room to which the client is chatting
	room *room
}

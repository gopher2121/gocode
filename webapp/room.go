package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type room struct {
	// forward is a channel to hold incoming messages
	// that should be forwarded to all the clients
	forward chan []byte

	// channel for joining the room
	join chan *client

	// channel for leaving the room
	leave chan *client

	// map to hold the current clients in the room
	clients map[*client]bool
}

func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			// forward the message to all clients
			for client := range r.clients {
				select {
				case client.send <- msg:
					//send the message
				default:
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

// upgrade to use the modern websocket

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

// method to implement ServeHTTP

func (r *room) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(rw, req, nil)
	if err != nil {
		log.Fatal("ERROR CODE :", err)
		return
	}

	// create a client
	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}

	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}

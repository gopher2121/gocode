package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

// implementing ServeHTTP

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("serve http :", err)
		return
	}
	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}
	r.join <- client
	defer func() {
		r.leave <- client
	}()
	go client.write()
	client.read()
}

type room struct {
	// forward is a channel to hold incoming messages
	// and will forward to clients
	forward chan []byte

	//join channel is for clients to join the room
	join chan *client

	//leave channel is for clients to leave the room
	leave chan *client

	//clientsmap will hold all the current clients in the room
	clients map[*client]bool
}

func (r *room) run() {
	for {
		select {
		// if a message is received in join channel
		case client := <-r.join:
			r.clients[client] = true
		// if a message is received in leave channel
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
		// if a message is received in forward channel
		// for all the clients, send the message
		case msg := <-r.forward:
			for client := range r.clients {
				select {
				case client.send <- msg:
					// send the message
				default:
					delete(r.clients, client)
					close(client.send)

				}
			}
		}
	}
}

func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

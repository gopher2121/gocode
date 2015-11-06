package main

type room struct {
	// forward is a channel to hold incoming messages
	forward chan []byte
}

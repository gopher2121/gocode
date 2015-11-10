// webserver

package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HELLO EVERYONE, I am go server with superb concurrency power")
}

func main() {
	//set the router
	http.HandleFunc("/", sayHelloName)

	//start the server , listen to a port
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}

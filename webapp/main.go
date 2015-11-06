package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`
            <html>
            <head>
            <title> web server </title>
            <body>
            <h1> consistency is the key and one day I will be where I want to be </h1>
            </body>
            </html>
			`))
	})

	//  start the web server at port :8080 for now
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("error information:", err)
	}
}

//we will create a simple web server as of now

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
                 </head>
                 <body>
                    <h1> let's hope our simple web server works !!! </h1>
                 </body>
              </html>    			     

			`))
	})

	// start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(" error code :", err)
	}

}

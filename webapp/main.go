package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

// our own type for handling the template

type templateHandler struct {
	filename string             //html file
	once     sync.Once          // compile once
	templ    *template.Template //reference of the compiled template
}

// method to satisfy the ServeHTTP method so that we can pass this struct instance in our handle function

func (t *templateHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// no matter how many go routines call it , the function inside do will be executed once only
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(rw, nil)
}

func main() {

	http.Handle("/", &templateHandler{filename: "index.html"})

	/*
				// some old code to test whether server is working or not . Remove it once the final code is ready
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
	*/
	//  start the web server at port :8080 for now
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("error information:", err)
	}
}

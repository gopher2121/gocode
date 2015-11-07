package main

import (
	"flag"
	"github.com/gocode/trace"
	"log"
	"net/http"
	"os"
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
	t.templ.Execute(rw, req)
}

func main() {
	// command line flags to make the address configurable
	var addr = flag.String("addr", ":8080", "address of the application")
	flag.Parse()
	// create a new room
	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	http.Handle("/chat", MustAuth(&templateHandler{filename: "index.html"}))
	//	http.Handle("/", &templateHandler{filename: "index.html"})

	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.Handle("/room", r)
	// run the room as a separate go routine
	go r.run()
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
	log.Println("starting web server on ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("error information:", err)
	}
}

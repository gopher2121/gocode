package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

// templateHandler struct so that we can use the template and use html
type templateHandler struct {
	filename string
	once     sync.Once
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	r := newRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	go r.run()
	// some code to check whether server is running correctly or not
	/*	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`
				<html>
					<head>
						<title>Chatting Application</title>
					<head>
					<body>
						Let's chat !!!!!!!!!!!
					</body>
					</html>

				`))
		})
	*/
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Listen and serve", err)
	}
}

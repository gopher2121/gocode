// simple router

package main

import (
	"fmt"
	"net/http"
)

type myMux struct {
}

func (m *myMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		// route to sayHelloName
		sayHelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return

}
func sayHelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "yjgjg")
}

func main() {
	mux := &myMux{}
	http.ListenAndServe(":8080", mux)
}

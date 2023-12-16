package main

import (
	"fmt"
	"log"
	"net/http"
)

type membersHandler struct{}

func (m *membersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "members handler called")
}

type groupsHandler struct{}

func (g *groupsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "groups handler called")
}

func main() {
	mh := &membersHandler{}
	gh := &groupsHandler{}

	mux := http.NewServeMux()
	mux.Handle("/members", mh)
	mux.Handle("/groups", gh)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "welcome!!")
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}

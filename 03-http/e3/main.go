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
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "groups handler called")
}

type departmentPostHandler struct{}

func (g *departmentPostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "departments handler called With POST method.")
}

type departmentGetHandler struct{}

// curl  http://localhost:8080/departments/aaaa -v -X GET
func (g *departmentGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	departmentId := r.PathValue("id")
	fmt.Fprintf(w, "departments handler called With Get method.\n")
	fmt.Fprintf(w, "departments id : %s \n", departmentId)

	r.SetPathValue("id", "bbbb")
	fmt.Fprintf(w, "changed departments id : %s \n", r.PathValue("id"))
}

func main() {
	mh := &membersHandler{}
	gh := &groupsHandler{}

	mux := http.NewServeMux()
	mux.Handle("/members", mh)
	mux.Handle("/groups", gh)

	dh := &departmentPostHandler{}
	dgh := &departmentGetHandler{}

	mux.Handle("POST /departments", dh)
	mux.Handle("GET /departments/{id}", dgh)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

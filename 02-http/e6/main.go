package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/members/{member-id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintf(w, "your member id is %s", vars["member-id"])
	}).Methods(http.MethodGet)

	// curl http://localhost:8080/groups/1 -v
	groupsMux := r.PathPrefix("/groups").Subrouter()
	groupsMux.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Requested Group id : %s\n", mux.Vars(r)["id"])
	}).Methods(http.MethodGet)

	// curl http://localhost:8080/members/1/1111 -v
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found by Manty", http.StatusNotFound)
	})

	// curl http://localhost:8080/members/1 -v -X POST
	r.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Method Not Allowed by Manty", http.StatusMethodNotAllowed)
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}

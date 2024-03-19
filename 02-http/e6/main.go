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

	log.Fatal(http.ListenAndServe(":8080", r))
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// TODO 1 GET /members
	mux.HandleFunc("GET /members", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You reqeuted url=%s, method=%s \n", r.URL.Path, r.Method)
	})

	// TODO 2 GET /members/{id}

	// TODO 3 POST /members

	// TODO 4 DELETE /members/{id}

	log.Fatal(http.ListenAndServe(":8080", mux))
}

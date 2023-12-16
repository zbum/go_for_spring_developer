package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})

	log.Println("Listening on http://0.0.0.0:8080")
	err := http.ListenAndServe(":8080", h)
	if err != nil {
		return
	}
}

package main

import (
	"log"
	"net/http"
)

func main() {
	var messageHandler http.Handler

	mux := http.NewServeMux()
	mux.Handle("GET /messages/", messageHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

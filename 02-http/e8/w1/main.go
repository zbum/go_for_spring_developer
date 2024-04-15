package main

import (
	"log"
	"net/http"
)

func main() {
	messageHandler := newMessageHandler()

	mux := http.NewServeMux()
	mux.Handle("/messages/", messageHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

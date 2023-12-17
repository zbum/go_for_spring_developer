package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func cancellableGreetHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling greeting request")
	defer log.Println("Handled greeting request")

	completeAfter := time.After(5 * time.Second)
	ctx := r.Context()

	for {
		select {
		case <-completeAfter:
			fmt.Fprintln(w, "Hello Gopher!")
			return
		case <-ctx.Done():
			err := ctx.Err()
			log.Printf("Context Error: %s", err.Error())
			return
		default:
			time.Sleep(1 * time.Second)
			log.Println("Greetings are hard. Thinking...")
		}
	}
}

func main() {
	http.HandleFunc("/", cancellableGreetHandler)
	log.Println("Server listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

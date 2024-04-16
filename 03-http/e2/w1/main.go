package main

import (
	"log"
	"net/http"
)

// TODO RequestInfoHandler 를 구현하지 않고 HandlerFunc를 이용합니다.
func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}

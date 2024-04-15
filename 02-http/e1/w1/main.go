package main

import (
	"log"
	"net/http"
)

// TODO RequestInfoHandler 를 구현하여 ListenAndServe 의 두번째 인자로 설정합니다.
func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}

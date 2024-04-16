package main

import (
	"go_for_spring_developer/03-http/e7/handler"
	"log"
	"net/http"
)

func main() {

	mh := handler.NewIndexHandler()

	mux := http.NewServeMux()

	//http://localhost:8080/index?name=zbum
	mux.HandleFunc("GET /index", mh.IndexPage)

	//http://localhost:8080/index2?name=zbum
	mux.HandleFunc("GET /index2", mh.IndexPageWithTemplate)

	//http://localhost:8080/index3?name=zbum
	mux.HandleFunc("GET /index3", mh.IndexPageWithTemplateCache)

	//http://localhost:8080/index4?name=zbum
	mux.HandleFunc("GET /index4", mh.IndexPageWithTemplateCacheAndEmbed)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

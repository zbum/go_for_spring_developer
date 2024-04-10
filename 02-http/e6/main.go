package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// curl http://localhost:8080/members/1\?a\=1\&b\=\2\&a\=3 -v
	r.HandleFunc("/members/{member-id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintf(w, "your member id is %s\n", vars["member-id"])
		fmt.Fprintf(w, "%v\n", r.URL.Query())
	}).Methods(http.MethodGet)

	/*
		PathPrefix 를 이용한 서브라우터
	*/
	groupsSubRouter := r.PathPrefix("/groups").Subrouter()
	// curl http://localhost:8080/groups/1 -v
	groupsSubRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Requested Group id : %s\n", mux.Vars(r)["id"])
	}).Methods(http.MethodGet)

	// curl http://localhost:8080/groups -v
	groupsSubRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Requested All Groups : \n")
	}).Methods(http.MethodGet)

	//curl http://localhost:8080/groups -v -X POST -d'Body Contents'
	groupsSubRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error Occurred", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "Register Group!! : %s\n", string(body))

	}).Methods(http.MethodPost)

	/*
		Methods 를 이용한 서브라우터
	*/
	deleteSubRouter := r.Methods(http.MethodDelete).Subrouter()
	// curl http://localhost:8080/delete/1 -v -X DELETE
	deleteSubRouter.HandleFunc("/delete/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "DELETE Requested delete id : %s\n", mux.Vars(r)["id"])
	})

	/*
		NotFoundHandler 변수를 설정하여 매핑되지 않는 URL 요청을 처리
	*/
	// curl http://localhost:8080/members/1/1111 -v
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found by Manty", http.StatusNotFound)
	})

	/*
		MethodNotAllowedHandler 변수를 설정하여 매핑되지 않는 Method 요청을 처리
	*/
	// curl http://localhost:8080/members/1 -v -X POST
	r.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Method Not Allowed by Manty", http.StatusMethodNotAllowed)
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}

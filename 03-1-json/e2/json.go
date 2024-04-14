package main

import (
	"fmt"
	"github.com/go-json-experiment/json"
)

type StudentRequest struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
type StudentResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func main() {
	marshal()
	unmarshal()
}

func marshal() {
	var id int64 = 1
	response := StudentResponse{Id: id, Name: "Manty"}
	marshal, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	fmt.Println("[1] marshal:", string(marshal))
}

func unmarshal() {
	requestString := `{"id":10,"name":"Manty"}`
	var request StudentRequest
	err := json.Unmarshal([]byte(requestString), &request)
	if err != nil {
		panic(err)
	}
	fmt.Println("[2] unmarshal : ", request)
}

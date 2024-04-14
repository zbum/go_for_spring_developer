package main

import (
	"fmt"
	"github.com/go-json-experiment/json"
)

type StudentRequest struct {
	Id   int64
	Name string
}
type StudentResponse struct {
	Id   int64
	Name string
}

func main() {
	marshal()
	unmarshal()
}

func marshal() {
	response := StudentResponse{Id: 10, Name: "Manty"}
	marshal, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	fmt.Println("[1] marshal:", string(marshal))
}

func unmarshal() {
	requestString := `{"Id":10,"Name":"Manty"}`
	var request StudentRequest
	err := json.Unmarshal([]byte(requestString), &request)
	if err != nil {
		panic(err)
	}
	fmt.Println("[2] unmarshal : ", request)
}

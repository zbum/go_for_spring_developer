package main

import (
	"fmt"
	"github.com/go-json-experiment/json"
)

type StudentRequest struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}
type StudentResponse struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

func main() {
	marshal1()
	marshal2()
	unmarshal1()
	unmarshal2()
}

func marshal1() {
	var id int64 = 1
	var name string = "Manty"
	response := StudentResponse{Id: &id, Name: &name}
	marshal, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	fmt.Println("[1] marshal:", string(marshal))
}

func marshal2() {
	response := StudentResponse{}
	marshal, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	fmt.Println("[2] marshal:", string(marshal))
}

func unmarshal1() {
	requestString := `{"id":10,"name":"Manty"}`
	var request StudentRequest
	err := json.Unmarshal([]byte(requestString), &request)
	if err != nil {
		panic(err)
	}
	fmt.Println("[3] unmarshal : ", request)
}

func unmarshal2() {
	requestString := `{"id":10}`
	var request StudentRequest
	err := json.Unmarshal([]byte(requestString), &request)
	if err != nil {
		panic(err)
	}
	fmt.Println("[3] unmarshal : ", request)
}

package main

import "fmt"

type Student struct {
}

func Init() *Student {
	var student = Student{}
	return &student
}

func main() {
	student := Init()
	fmt.Println(student)
}

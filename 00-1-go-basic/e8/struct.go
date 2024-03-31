package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	Id   int64
	Name string
}

func main() {
	// Student 구조체 포인터 생성
	mantyPtr := new(Student)
	mantyPtr.Id = 1
	mantyPtr.Name = "manty"
	fmt.Println(mantyPtr)

	// Student 구조체 생성(zero value)
	zero := Student{}
	fmt.Println(zero)

	// Student 구조체 생성
	comtin := Student{
		Id:   2,
		Name: "comtin",
	}
	fmt.Println(comtin)

	// Student 구조체 포인터 생성
	comtinPtr := &Student{
		Id:   2,
		Name: "comtin",
	}
	fmt.Println(comtinPtr)

	// Student 구조체의 슬라이스
	var students []Student
	for inx := 0; inx < 10; inx++ {
		student := Student{Id: int64(inx + 1), Name: "Name" + strconv.Itoa(inx)}
		students = append(students, student)
	}

	for _, student := range students {
		fmt.Println(student)
	}

}

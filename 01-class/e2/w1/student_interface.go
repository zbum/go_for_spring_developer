package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Id   int32
	Name string
}

type SortableStudents []Student

// TODO : 함수를 완성하세요.
func (s SortableStudents) Len() int {
	return 0
}

// TODO : 함수를 완성하세요.
func (s SortableStudents) Less(i, j int) bool {
	return true
}

// TODO : 함수를 완성하세요.
func (s SortableStudents) Swap(i, j int) {

}
func main() {
	data1 := SortableStudents([]Student{
		{9, "third"},
		{8, "second"},
		{5, "first"},
		{10, "fourth"},
	})
	sort.Sort(data1)

	fmt.Println(data1)
}

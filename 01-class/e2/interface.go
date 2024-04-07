package main

import (
	"fmt"
	"sort"
)

func main() {
	// 인터페이스를 사용하여 Sort 수행
	doSort()

	typeAssertion1()
	typeAssertion2()

	switchType(Unknown1{})
	switchType(Unknown2{})
	switchType("Test")

}

// 인터페이스의 사용
func doSort() {
	data1 := SortableSlice([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 10})
	sort.Sort(data1)
	fmt.Println(data1)

	sort.Sort(sort.Reverse(data1))
	fmt.Println(data1)
}

type SortableSlice []int

func (s SortableSlice) Len() int {
	return len(s)
}

func (s SortableSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s SortableSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Type Assertion
func typeAssertion1() {
	anInt := returnNumber()
	number := anInt.(int)
	number++
	fmt.Println(number)
}

func typeAssertion2() {
	anInt := returnNumber()
	number, ok := anInt.(int)
	if ok {
		number++
		fmt.Println(number)
	} else {
		fmt.Println("Type assertion Failed")
	}
}
func returnNumber() interface{} {
	return 12
}

// 타입 스위치

func switchType(x interface{}) {

	switch T := x.(type) {
	case Unknown1:
		fmt.Println("Unknown type")
	case Unknown2:
		fmt.Println("Entry type")
	default:
		fmt.Printf("Not supported type: %T\n", T)
	}
}

type Unknown1 struct {
	UnknownValue string
}

type Unknown2 struct {
	F1 int
	F2 string
	F3 Unknown1
}

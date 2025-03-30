package main

import (
	"fmt"
	"go_for_spring_developer/01-go-basic/e6-1/set"
)

func main() {
	set1 := set.New[string]()
	set1.Add("1")
	set1.Add("2")
	set1.Add("3")
	set1.Add("4")
	set1.Add("5")
	set1.Add("6")

	for value1 := range set1.All() {
		fmt.Println(value1)
	}
}

package main

import "fmt"

func main() {
	aSlice := []string{"tesla", "nvidia", "apple", "microsoft"}

	for i, v := range aSlice {
		fmt.Printf("Index : %d, Value : %s\n", i, v)
	}

	for inx := 0; inx < len(aSlice); inx++ {
		fmt.Printf("Index : %d, Value : %s\n", inx, aSlice[inx])
	}
}

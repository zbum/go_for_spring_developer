package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a command line argument.")
		return
	}

	argument := os.Args[1]
	value, err := strconv.Atoi(argument)
	if err != nil {
		fmt.Println("Can not convert to int:", argument)
		return
	}

	switch {
	case value == 0:
		fmt.Println("영")
	case value > 0:
		fmt.Println("양의 정수")
	case value > 10:
		fmt.Println("10보다 큰 정수")
	case value < 0:
		fmt.Println("음의 정수")
	default:
		fmt.Println("이 조건에 올 수 없습니다.", value)
	}
}

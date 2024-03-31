package main

import (
	"fmt"
)

func main() {
	// 슬라이스 선언 방법 1
	aSlice := []string{"tesla", "nvidia", "apple", "microsoft"}

	fmt.Println(aSlice, len(aSlice), cap(aSlice))

	// make 를 이용한 슬라이스 선언, 데이터는 모두 빈 문자열("")이다.
	bSlice := make([]string, 4)
	fmt.Println(bSlice, len(bSlice), cap(bSlice))

	// slice 에 값 설정
	bSlice[0] = "One"
	bSlice[1] = "Two"
	bSlice[2] = "Three"
	bSlice[3] = "Four"

	// bSlice[4] = "Five" 는 index out of range 를 발생시킴
	bSlice = append(bSlice, "Five")
	fmt.Println(bSlice, len(bSlice), cap(bSlice))

	// make 를 이용한 슬라이스 선언, 데이터는 모두 nil 이다.
	cSlice := make([]*string, 4)
	fmt.Println(cSlice, len(cSlice), cap(cSlice))

}

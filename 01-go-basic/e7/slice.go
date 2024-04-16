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
	// slice 에 더이상 공간이 없을때는 append 로 새 원소를 추가해야 합니다.
	bSlice = append(bSlice, "Five")
	fmt.Println(bSlice, len(bSlice), cap(bSlice))

	// slice 에 값을 계속 append로 대입하는 경우, 슬라이스의 용량은 자동으로 증가합니다.
	//
	// https://github.com/golang/go/blob/master/src/runtime/slice.go
	cSlice := make([]int, 0)
	for inx := 0; inx < 1000; inx++ {
		oldCap := cap(cSlice)
		cSlice = append(cSlice, inx)
		if oldCap != cap(cSlice) {
			fmt.Println("Len: ", len(cSlice), " ,Cap: ", cap(cSlice))
		}
	}
}

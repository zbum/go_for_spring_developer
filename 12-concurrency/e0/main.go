package main

import (
	"fmt"
	"time"
)

func main() {
	// 언버퍼드 채널 생성
	unbuffChan := make(chan int)

	fmt.Println("=== Unbuffered Channel 예제 ===")

	go func() {
		fmt.Println("언버퍼드: 송신 시도")
		unbuffChan <- 1 // 여기서 블록됨 (수신자가 준비될 때까지 대기)
		fmt.Println("언버퍼드: 송신 완료")
	}()

	// 송신자가 수신자를 기다리도록 잠시 대기
	time.Sleep(time.Second)
	fmt.Println("언버퍼드: 수신 준비")
	value := <-unbuffChan
	fmt.Println("언버퍼드: 수신 값:", value)

	// 로그에 송신완료는 왜 나오지 않을까?

}

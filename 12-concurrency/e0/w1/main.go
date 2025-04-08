package main

import (
	"fmt"
	"time"
)

func main() {

	buffChan := make(chan int)

	fmt.Println("\n=== Buffered Channel 예제 ===")

	// 버퍼가 있으므로 블록되지 않고 즉시 송신 가능
	fmt.Println("버퍼드: 첫 번째 송신 시도")
	buffChan <- 1
	fmt.Println("버퍼드: 첫 번째 송신 완료")

	fmt.Println("버퍼드: 두 번째 송신 시도")
	buffChan <- 2
	fmt.Println("버퍼드: 두 번째 송신 완료")

	// 버퍼가 가득 차면 여기서 블록됨
	go func() {
		fmt.Println("버퍼드: 세 번째 송신 시도")
		buffChan <- 3
		fmt.Println("버퍼드: 세 번째 송신 완료")
	}()

	// 수신 전 대기
	time.Sleep(time.Second)
	fmt.Println("버퍼드: 수신 시작")
	fmt.Println("버퍼드: 수신 값:", <-buffChan)
	fmt.Println("버퍼드: 수신 값:", <-buffChan)

}

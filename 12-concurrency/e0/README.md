## 동시성
* Go 언어는 동시성(concurrency)을 지원합니다.
* Go 언어의 동시성은 고루틴(goroutine)과 채널(channel)을 통해 구현됩니다.

## goroutine
* Go 언어의 고루틴은 경량 스레드입니다.
* 쓰레드보다 매우 가벼우며 실무적으로 몇 천개 이상 동작해도 문제가 없습니다.
* 고루틴으로 실행하면 자체 콜 스택이 생성됩니다.
* 함수나 메서드 호출 앞에 `go` 키워드를 붙이면 고루틴으로 실행됩니다.
```go
  go list.Sort()
```

* 익명함수(func)으로 고루틴을 편리하게 사용할 수 있습니다. 
```go
func Announce(message string, delay time.Duration) {
    go func() {
        time.Sleep(delay)
        fmt.Println(message)
    }()  // Note the parentheses - must call the function.
}
```

## 데이터 공유(channel)
* Go 언어는 쓰레드간에 공유 메모리를 사용하지 않고 채널(channel)을 통해 데이터를 전송하는 것을 권장합니다.

> Do not communicate by sharing memory; instead, share memory by communicating.
> 
> 공유메모리로 통신하지 말고, 통신으로 공유메모리를 사용하라.
>
> - Rob Pike
>
* 채널은 make 키워드로 생성합니다.
```go
ci := make(chan int) //unbuffered channel
cj := make(chan int, 0) //unbuffered channel
cs := make(chan *os.File, 100) //buffered channel
```
### unbuffered channel
* make(chan int)로 생성
* 송신자와 수신자가 동시에 준비되어 있어야 함
* 송신 시 수신자가 없으면 송신 고루틴이 블록됨
* 동기화(synchronize) 에 유용

### buffered channel
* make(chan int, capacity)로 생성 (capacity는 버퍼 크기)
* 버퍼 크기만큼 값을 저장 가능
* 버퍼가 가득 차기 전까지는 송신이 블록되지 않음
* 비동기 통신에 유용

### 채널 데이터 전송
* 채널에 데이터를 전송할 때는 `<-` 연산자를 사용합니다.
```go
ci <- 1
```
* 채널에서 데이터를 수신할 때도 `<-` 연산자를 사용합니다.
```go
value := <-ci
```

### 채널 예제 (e0)
```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    channel := make(chan string)
    go func() {
        for i := 0; i < 5; i++ {
            channel <- fmt.Sprintf("Hello %d", i)
            time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
        }
    }()

    
    for i := 0; i < 5; i++ {
        fmt.Println(<-channel)
    }
}
```

## 실습(w1)
* buffered channel을 사용하여 deadlock을 피하도록 수정해 보세요.


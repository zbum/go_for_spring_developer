## go routine 커뮤니케이션
* 위 예제에서 main 함수는 go routine 에서 실행한 결과를 볼 수 없습니다.
* go routine 간 통신을 위해 Channel을 사용해 보겠습니다. 

### Channel 선언 및 할당
```go
var c chan int
c = make(chan int)
```

### 채널에 시그널 보내기
```go
c <- 1
```

### 채널에서 시그널 받기
* 화살표(<-)는 데이터의 흐름을 보여주는 것 같습니다.
```go
value = <-c
```
## channel 을 이용한 통신
```go
func main() {
	c := make(chan string)
	go boring("boring!", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c) // Receive expression is just a value.
	}
	fmt.Println("You're boring; I'm leaving.")
}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i) // Expression to be sent can be any suitable value.
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

```
* c 변수에 담긴 Channel 이 main 과 boring 고루틴을 연결하고 있어 커뮤니케이션 할 수 있게 되었습니다.

## 동기화(Synchronization)
* 위 코드의 main 함수에서 <-c 를 호출하면 메시지가 수신될때까지 블로킹 됩니다. 
* 비슷하게 boring 함수의 c <- value 도 수신측에서 준비될 때까지 블로킹 됩니다. 
* 발신자와 수신자가 모두 준비가 되어야 커뮤니케이션 할 수 있고 , 그렇지 않으면 대기상태로 빠지게됩니다. 

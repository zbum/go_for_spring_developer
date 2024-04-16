## 기본 함수
* 다음과 같은 무한루프를 돌면서 메시지를 1초 내에 1회 표시하는 boring 이라는 함수가 있습니다.  
```go
func main() {
    boring("boring!")
}

func boring(msg string) {
    for i := 0; ; i++ {
        fmt.Println(msg, i)
        time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
    }
}
```
## go routine
* go 라는 키워드를 함수 실행 앞에 붙이면 go routine 으로 함수가 실행합니다. 
```go
func main() {
    go boring("boring!")
    fmt.Println("I'm listening.")
    time.Sleep(2 * time.Second)
    fmt.Println("You're boring; I'm leaving.")
}
```
* go routine 으로 실행하면 자체 콜 스택이 생성됩니다. 
* 쓰레드보다 매우 가벼우며 실무적으로 몇 천개 이상 동작해도 문제가 없습니다. 






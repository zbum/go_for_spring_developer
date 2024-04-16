
## 흐름 제어
### if/else
* 조건문에 괄호를 사용하지 않습니다.
```go
err := anyFunctionCall()
if err != nil {
    // 에러 처리
}
```
>java 에서의 null 은 Go 에서 nil 입니다.

### switch
* Go의 switch 문은 두 가지 타입이 있습니다.
* Go는 Java와 달리 case의 조건이 맞으면 break 가 없어도 다음 조건을 실행하지 않습니다.
* 다음 조건을 계속 평가하고 싶다면  fallthrough 를 사용합니다.
1. switch 문이 평가할 표현식을 가진 경우
#### (E3)
```go
switch argument {
case "0":
    fmt.Println("Zero")
case "1":
    fmt.Println("One")
case "2", "3", "4":
    fmt.Println("2 or 3 or 4")
    fallthrough
default:
    fmt.Println("Value:", argument)
}
```

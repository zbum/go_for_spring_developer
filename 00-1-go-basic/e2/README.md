
## 변수 선언
* var 키워드를 사용합니다.
```go
var count int
```
* 초기 값이 있다면 데이터 타입은 생략할 수 있습니다.
```go
var count = 1
```
* var 키워드 대신 := 을 이용해서 변수를 선언할 수 있습니다. 데이터 타입은 추론합니다.
* 전역 변수에는 이 형식을 사용할 수 없습니다.
```go
// short assignment
count := 1
```

## 변수 출력
* 화면(stdout)에 데이터를 출력하려면 fmt.Println() 함수를 사용합니다. 자바의 System.out.println()와 비슷합니다.
* fmt.Printf()는 자바의 System.out.printf()와 비슷한 기능을 가지고 있습니다.

### (e2)
```go
package main

import (
    "fmt"
    "math"
)

// int 타입의 전역변수
var Global int = 1234
// 두번째 전역변수, 타입은 int로 추론, ANOTHER_GLOBAL로 이름을 짓지 말것
var AnotherGlobal = -5678

func main() {
    // int 타입의 지역변수 0으로 초기화
    var j int
    // i 도 int 가 된다.
    i := Global + AnotherGlobal
    fmt.Println("Initial j value:", j)
    j = Global

    // math.Abs() int64 매개변수가 필요합니다. 명시적으로 변환합니다.
    k := math.Abs(float64(AnotherGlobal))
    
    // 형식을 잘 보세요.
    fmt.Printf("Global=%d, i=%d, j=%d, k=%.2f.\n", Global, i, j, k)
}
```

* output
```
Initial j value: 0
Global=1234, i=-4444, j=1234, k=5678.00.
```
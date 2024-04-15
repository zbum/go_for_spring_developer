
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

## 상수
* 상수는 const 키워드를 사용합니다.
* 기본 선언 방법은 const 뒤에 상수명을 적고 그뒤에 상수 타입, 상수값을 기술하는 방식입니다. 
```go
package a 

const studentCount int = 10
const helloPrefix string = "Hi"
```
> Java 의 네이밍 컨벤션은 STUDENT_COUNT, HELLO_PREFIX 이겠지만 Go의 변수,상수, 패키지 함수 이름에 _ 네이밍을 사용하지 않습니다.
> https://gosudaweb.gitbooks.io/effective-go-in-korean/content/names.html

* 상수도 변수와 같이 타입추론이 동작합니다. 
* 상수를 묶어서 선언 할 수 있습니다. 
```go
package b

const (
    Visa = "Visa"
    Master = "MasterCard"
    Amex = "American Express"
)
```
* java의 ENUM 은 Go에서 제공되지 않습니다. 대신 iota 라는 키워드를 이용하여 적용할 수 있습니다. 
* 아래 예제와 같이 iota 는 상수 묶음에서 1씩 증가된 값으로 처리됩니다. 
```go
package c

const (
    Apple   = iota // 0
    Grape   = iota // 1
    Orange  = iota // 2
)
```
* 상수의 내용이 같은 경우는 생략할 수 있기 때문에 다음과 같이 코드를 수정할 수 있습니다. 
```go
package c

const (
    Apple   = iota // 0
    Grape          // 1
    Orange         // 2
)
```

* iota의 0을 제거하고 싶으면 _ 를 사용합니다.
```go
package c

const (
    _ = iota       // 0
    Apple   = iota // 1
    Grape          // 2
    Orange         // 3
)
```

## 실습 (e2/w1)
* humanReadable 함수는 입력받은 바이트 값을 사람이 읽기 좋은 단위로 변환한 문자열을 반환해 주는 코드 입니다. 
* 현재 구현은 KiB 까지만 변환하고 있습니다. MiB, GiB, TiB 도 변환할 수 있도록 코드를 수정해 주세요.
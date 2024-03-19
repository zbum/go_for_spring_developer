## HandlerFunc Helper
* 매번 아래와 같이 Handler 인터페이스를 구현하는 것은 type 을 선언해야 하고 ServeHTTP를 작성해야 하기 때문에 매우 번거롭습니다. 
```go
type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello")
}
```
* 그래서 간편하게 구현할 수 있는 HandlerFunc 를 사용할 수 있습니다. 
### HandlerFunc 정의
```go
// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler that calls f.
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
f(w, r)
}
```
* type을 함수형으로 선언하면 '함수 시그니쳐의 커스텀 타입' 으로 선언 됩니다.

### HandlerFunc 사용
```go
func main() {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})

	log.Println("Listening on http://0.0.0.0:8080")
	err := http.ListenAndServe(":8080", h)
	if err != nil {
		return
	}
}
```

## Go 문법 해석
### function type variable
* HandlerFunc 는 처음 Go를 접한 사람에게 다소 어려운 문법입니다. (당황할 수도..)
* `http.HandlerFunc()` 는 함수를 실행하는 것 처럼 보이지만 `http.Handler` 인터페이스를 구현하는 표현식입니다.
* Go 는 함수 타입으로 변수를 작성할 수 있습니다. 이것은 int 등의 타입으로 변수를 작성하는 것과 동일합니다.
```go
// `func()` 타입의 변수 선언
// (기본값 nil로 초기화)
var d func()

// 표준 출력을 하는 함수 할당
d = func() {
  fmt.Println("Ella guru")
}

// `func() error` 타입의 변수 선언
// (기본값 nil로 초기화)
var e func() error = func() error {
    return nil
}

// `func(x int) bool` 타입의 변수 선언
// ( 이면 x > 5 true를 반환하는 함수로 초기화).
f := func(x int) bool {
  return x > 5
}

// `func() string` 타입의 함수 선언.
// (기본값 nil로 초기화)
var g func() string
```
* 물론 다른 타입의 값( `func(x int) int` )을 아래와 같이 할당하는 경우 `컴파일 오류`가 발생합니다.
```go
d = func(x int) int {
  return x + 1
}
```
* 작성한 함수타입의 변수는 다음의 형태로 실행할 수 있습니다. 
```go
d()
err := e()
result := f(10)
```
### 커스텀 타입정의

* 기본 타입이나 구조체는 커스텀 타입을 정의할 수 있습니다. 함수 타입도 물론 가능합니다.
```go
// int 타입인 커스텀타입 myInt를 정의 합니다. 
type myInt int

// x int 필드를 가진 myStruct 커스텀 타입을 정의 합니다. (왠지 익숙!)
type myStruct struct {
  x int
}

// myInt 타입으로 a 변수 선언
// 20으로 초기화
var a myInt = 20

// myStruct 타입으로 b 변수 선언
// x 필드 값을 20으로 하는 스트럭트로 초기화
b := myStruct{
  x: 20,
}

// 함수 시그니쳐가 `func(x int) int` 인 myIntFunction 커스텀 타입 정의  
type myIntFunction func(x int) int

// myIntFunction 타입으로 c 변수 선언
// 입력값의 2배를 반환하는 함수로 초기화
var c myIntFunction = func(x int) int {
    return x * 2
}
```

### 함수호출 아닙니다!
* 얼핏 보면 HandlerFunc는 함수 처럼 보이지만 http.HandlerFunc 는 타입입니다. 
```go
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})
```
* http.HandlerFunc(...) 는 타입 변환(type conversion) 이라고 하는 형태입니다. 
* 마치 아래의 기본 타입을 변환하는 형태와 동일합니다.
```go
var a int32 = 12
var b int64 = int64(a) // convert from int32 to int64
```

* 동일한 시그니쳐가 있는 경우, 함수 타입간에도 변환이 가능합니다. 
```go
type myFunctionOne func(x int) int
type myFunctionTwo func(y int) int

var c myFunctionOne = func(x int) int {
	return x * 2
}

// myFunctionOne 에서 myFunctionTwo로 타입변환
var d myFunctionTwo = myFunctionTwo(c)
```
### http.Handler
* http.Handler 는 인터페이스 입니다.
```go
// inside net/http package

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```
* Go 에서는 `ServeHTTP(ResponseWriter, *Request)` 메소드를 가지면 이 인터페이스를 구현한 것입니다. 
* 구조체 타입에 메소드를 추가하는 것을 알고 있습니다.
```go
type myStruct struct {
    x int
}

// define doSomething method on type *myStruct.
func (m *myStruct) doSomething() {
    fmt.Println(m.x) // print the field value.
}
```
* 물론 함수 타입에도 메소드를 정의할 수 있습니다. 
```go
type myFunction func() int

 // define doSomething method on type myFunction
func(f myFunction) doSomething() {
	// call f (a value of type myFunction).
	result := f()
	// print the result of the function call.
	fmt.Println(result)
}
```

* 그러면 다음과 같이 호출할 수 있습니다. 
```go
var c myFunction = func() int {
	return 5
}

var d myFunction = func() int {
	return 12
}

c.doSomething() // prints: 5
d.doSomething() // prints: 12
```
* doSomething 내부에서 f 변수에 "값으로서의 함수" 가 포함되어 있음을 주목하세요.

## http.HandlerFunc 구현
* doSomething 과 동일한 형태입니다.
```go
// inside net/http package

type HandlerFunc func(ResponseWriter, *Request)

 // define ServeHTTP method on type HandlerFunc.
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	 // call f (a value of type HandlerFunc)
	 // with the provided parameters w and r.
	f(w, r)
}
```
* HandlerFunc 는 ServeHTTP(w ResponseWriter, r *Request) 를 포함하고 있어 http.Handler 인터페이스를 구현한 것으로 이해할 수 있습니다. 

### 출처
* https://www.willem.dev/articles/http-handler-func/
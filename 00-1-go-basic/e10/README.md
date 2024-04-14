


## 오류 처리
* Go 언어에는 Java 의 throws, throw, try-catch-finally 와 같은 키워드가 없습니다.
* 함수의 호출 결과에 오류가 발생했을 때는 error 타입으로 반환할 수 있습니다.
* 예를 들어 파일을 열때, Open 이라는 함수를 사용합니다.
```go
func Open(name string) (file * File, err error)
```
* 이 함수는 2개의 반환타입을 가집니다. 함수 호출결과가 정상이라면 file을 반환하고 err는 nil로 반환합니다.
* Open 함수 실행 중에 오류가 발생하면 err 에 에러 정보가 포함됩니다.
* go 언어에서 다음의 코드는 Java의 try-catch 문과 같이 자주 나타납니다.
```go
f, err := os.Open("filename.ext")
if err != nil {
  log.Fatal(err)
}
```

### error 타입
* error는 아래와 같이 정의된 내장 인터페이스 타입 입니다.
```go
type error interface {
	Error() string
}
```
* errors 패키지에는 단순한 에러 구현체가 제공됩니다. 
```go
// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
```
* 에러메시지를 errors.New 함수를 이용해서 에러 처리를 할 수 있습니다. 
```go
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	// implementation
}
```
* 다음 예제에서 에러를 처리하는 일반적인 코드를 확인 할 수 있습니다. 
```go
f, err := Sqrt(-1)
if err != nil {
    fmt.Println(err)
}
```


> 출처: https://github.com/astaxie/build-web-application-with-golang/blob/master/en/11.1.md


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
* type을 함수형으로 선언하면 '함수 시그니쳐의 별칭 타입' 으로 선언 됩니다.
* 

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
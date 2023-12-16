## Handler Interface
* 처음부터 이해해야 할 가장 중요한 개념은 "net/http"패키지의 인터페이스입니다. 
* 코드에서 "Handler"를 만든다는 것은 이 단일 메서드 인터페이스를 구현하는 사용자 지정 형식(일반적으로 구조체)을 만든다는 것을 의미합니다.
```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

* 이 인터페이스를 만족하기 위해서는 "http.ResponseWriter" 타입의 값과 "http.Request" 타입의 포인터를 받아들이는 메소드를 추가할 필요가 있습니다.
* "ResponseWriter" 는 응답을 처리하기 위해 사용합니다. 
* "Request" 는 http 메소드, path, header 등 http 요청의 모든 정보를 담고 있습니다. 

* 리스닝 주소와 핸들러 포인터를 받는 "http.ListenAndServe" 에서 우리가 정의한 핸들러를 사용할 수 있습니다.

```go
func ListenAndServe(addr string, handler Handler) error
```
* 간단한 handler 를 사용하는 예제는 다음과 같습니다. 
```go
package main

import (
	"fmt"
	"net/http"
)

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {
	http.ListenAndServe(":8080", &helloHandler{})
}
```
### 실습 1
1. 위 코드를 작성하고 실행합니다. 프로그램이 실행하는 동안 아무런 아웃풋은 없습니다.
2. 새로운 터미널에서 "http://localhost:8080" 을 호출해 봅니다. 결과로 "hello" 가 표시되는 것을 확인 합니다.


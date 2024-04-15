## 개요
* 이번 장에서는 스프링부트에서 Embedded Tomcat 으로 자동 구성된 HTTP 서비스를 Go 언어에서 구현하겠습니다. 
* Go 언어는 표준라이브러리에서 Http Server, Http Client를 포함하고 있습니다. 
* 개발의 편의나 성능 개선을 위한 3rd 파티 라이브러리도 사용할 수 있습니다. 

## HTTP 서버 시작하기
* 웹서버를 시작할때는 http.ListenAndServe() 함수를 사용합니다.
```go
func ListenAndServe(addr string, handler Handler) error
```
* 구현
```go
package main

import (
    "log"
    "net/http"
)

func main() {
    log.Fatal( http.ListenAndServe(":8080", nil) )
}
```
* http.ListenAndServe() 함수의 첫번째 인자는 TCP 네트워크 주소를 입력해야 하고 
* 두번째 인자는 사용자 요청을 처리할 Handler 인터페이스를 구현한 변수를 설정하면 됩니다. 

> 두번째 인자를 nil 로 처리하면 모든 요청에 404 응답을 제공합니다.

## 표준 라이브러리의 Handler Interface
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
### workshop(e1/w1)
1. Http 요청의 URL.Path와 HTTP Method를 응답하는 RequestInfoHandler 를 구현합시다. 
2. URL 은 r.URL.Path, Method는 r.Method 로 얻을 수 있습니다.
3. 다음과 같이 curl 로 호출했을때 결과를 확인 할 수 있어야 합니다.
```
$  curl localhost:8080/aaaa/bbbb
URL: /aaaa/bbbb, METHOD: GET 
$  curl localhost:8080/aaaa/bbbb -X POST
URL: /aaaa/bbbb, METHOD: POST
```

## 생각할 점
1. 매번 인터페이스를 구현하는 것은 보일러 플레이트(boiler plate)가 됩니다. 
2. 여러가지 URL을 처리하는 방법이 필요합니다.

## 요약 
* http.ListenAndServe 함수로 HTTP 서버를 기동하였습니다. 
* http.Handler 인터페이스를 구현하여 사용자 요청에 응답할 수 있습니다.
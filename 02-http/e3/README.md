## Multiplexers
* HTTP 서버와 애플리케이션을 구현할때, 일반적으로 다양한 엔드포인트에 대한 요청을 별도의 핸들러로 라우팅하려고 합니다. 
* 이 라우팅을 처리하는 구성을 일반적으로 줄여서 "멀티플렉서" 또는 "mux" 라고 합니다. Go 에서는 ServeMux 타입이 이 라우팅을 로직을 처리합니다. 
<br />
* 기본으로 특별히 지정하지 않으면 DefaultServeMux 를 사용합니다. 

```go
package main

import (
	"io"
	"net/http"
	"log"
)

// hello world 를 응답합니다.
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
    // DefaultServeMux 에 /hello 엔드포인트와 HelloServer를 매핑합니다.
	http.HandleFunc("/hello", HelloServer)
    // ListenAndServe 의 두번째 인자가 nil 일때 DefaultServeMux 를 사용합니다.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```
## Custom ServeMux
* 당연히, 애플리케이션의 라우팅을 처리하는 ServeMux를 정의할 수 있습니다. 다음은 애플리케이션에 대한 다양한 요청을 처리하기 위한 사용자 정의의 예입니다.
```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

type membersHandler struct{}

func (m *membersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "members handler called")
}

type groupsHandler struct{}

func (g *groupsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "groups handler called")
}

func main() {
	mh := &membersHandler{}
	gh := &groupsHandler{}

	mux := http.NewServeMux()
	mux.Handle("/members", mh)
	mux.Handle("/groups", gh)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "welcome!!")
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
```
* 클라이언트를 사용하여 위에 나열된 경로에 대해 여러 요청을 발송하고 각각에 대해 다른 응답을 확인합니다. 


## Custom ServerMux 의 Http Method 라우팅
### 1.21 이전 버전
* handler는 모든 Http Method 에 대해 응답하기 때문에 handler 내부에 처리할 수 없는 Method 는 405(Method Not Allowed) 처리를 하거나 분기를 해야 합니다.
```go
func (g *groupsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "groups handler called")
}
```

### 1.22 이후 
* func (mux *ServeMux) Handle(pattern string, handler Handler) 함수의 pattern 파라미터의 스팩이 다음의 Proposal 로 변경되어 Http Method 라우팅도 기술 할 수 있게 되었습니다. 
* https://github.com/golang/go/discussions/60227
* 따라서 다음과 같이 코드를 작성하여 Http Method 를 지정할 수 있습니다. 
```go
    dh := &departmentPostHandler{}
    mux.Handle("POST /departments", dh)
```
* GET 메소드는 HEAD 메소드 요청도 처리합니다.


## SpringMVC 의 @PathVariable 은?
* Go 언어에서 URL의 일부분을 와일드 카드로 두고 처리하는 것을 지원합니다.
* 와일드 카드는 / 사이의 전체 경로 이어야 합니다. 예를 들어 `/departments/qa_{seq}`와 같은 형태는 지원하지 않습니다.
* 전달한 와일드 카드를 얻어내기 위해 http.Request에 두 개의 메소드가 추가되었습니다. 
```go
package http

func (*Request) PathValue(wildcardName string) string
func (*Request) SetPathValue(name, value string)
``` 
* PathValue 타입 메소드는 와일드 카드의 키 이름으로 값을 찾아오는 기능을 제공합니다.
```go
departmentId := r.PathValue("id")
```
* SetPathValue 는 r.PathValue의 응답을 바꾸기 위해 존재합니다. 






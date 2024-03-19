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
* 라우팅에 지정되지 않은 엔드포인트(예 "/test" )로 요청을 보내보세요. 서버가 응답합니까?



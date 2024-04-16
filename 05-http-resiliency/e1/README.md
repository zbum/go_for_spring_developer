## Context
* 대부분의 전문가 수준의 go 프로그램은 context 패키지를 사용해야 합니다. 
* context를 사용하는 기본 전제는 클라이언트가 호출을 취소할 수 있고 타임아웃을 처리할 수 있다는 것입니다. 
* context패키지를 이해하기 위해 아래의 예제를 확인해 봅시다.

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling greeting request")
	defer log.Println("Handled greeting request")

	completeAfter := time.After(5 * time.Second)

	for {
		select {
		case <-completeAfter:
			fmt.Fprintln(w, "Hello Gopher!")
			return
		default:
			time.Sleep(1 * time.Second)
			log.Println("Greetings are hard. Thinking...")
		}
	}
}

func main() {
	http.HandleFunc("/", greetHandler)
	log.Println("Server listening on :8080...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
```
### 작업순서
1. 위 프로그램을 실행합니다. 
2. curl 을 이용해서 http://localhost:8080 을 호출합니다. 그러면 서버는 1초에 한번 thinking 로그를 생성합니다.
3. 서버가 응답하기 전에 curl 요청을 중단합니다. (CRTL+C)
4. 클라이언트가 취소한 후에도 서버는 thinking 로그를 계속 생성하는 것을 확인하세요.

* 위 작업을 성공적으로 수행했다면 클라이언트가 더이상 관심 없는 작업을 서버가 계속 처리하고 있다는 것을 알 수 있습니다. 
* 네트워크는 생각보다 자주 신뢰할 수 없습니다. 특히 모바일 환경에서는 더욱 그렇습니다.
* 스마트 클라이언트는 서버가 요청을 응답할때까지 네트워크 커넥션을 유지하기 위해 노력하겠지만, 언제든지 접속은 끊어지고 다시 붙을 수 있습니다.
* 서버는 이 사용 패턴에 대한 복원력이 있어야 하며 이것이 context 패키지가 필요한 부분입니다.

<br/>

* 위 프로그램과 거의 동일한 프로그램입니다.
* 이번에는 Request 로 부터 context를 얻어서 사용합니다.
```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling greeting request")
	defer log.Println("Handled greeting request")

	completeAfter := time.After(5 * time.Second)
	ctx := r.Context()

	for {
		select {
		case <-completeAfter:
			fmt.Fprintln(w, "Hello Gopher!")
			return
		case <-ctx.Done():
			err := ctx.Err()
			log.Printf("Context Error: %s", err.Error())
			return
		default:
			time.Sleep(1 * time.Second)
			log.Println("Greetings are hard. Thinking...")
		}
	}
}

func main() {
	http.HandleFunc("/", greetHandler)
	log.Println("Server listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### 작업순서
1. 위 프로그램을 실행합니다.
2. curl 을 이용해서 http://localhost:8080 을 호출합니다. 그러면 서버는 1초에 한번 thinking 로그를 생성합니다.
3. 서버가 응답하기 전에 curl 요청을 중단합니다. (CRTL+C)
4. 이번에는 서버는 context 가 취소되었기 때문에 즉시 처리를 멈춥니다.

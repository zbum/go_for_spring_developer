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
1. 위 프로그램을 복사해서 실행합니다. 
2. curl 을 이용해서 http://localhost:8080 을 호출합니다. 그러면 서버는 1초에 한번 thinking 로그를 생성합니다.
3. 서버가 응답하기 전에 curl 요청을 중단합니다. (CRTL+C)
4. 클라이언트가 취소한 후에도 서버는 thinking 로그를 계속 생성하는 것을 확인하세요.

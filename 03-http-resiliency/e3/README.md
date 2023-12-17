## Rate Limits, Retries, and Backoff
* 보통 서버는 엄청나게 많은 클라이언트 요청을 받게됩니다.
* 이런 경우, 서버는 유입 요청수를 제어할 수 있습니다.
* go 언어는 요청을 제어할수 있는 커뮤니티 패키지를 다수 찾을 수 있습니다. 
* 이번 실습에는 표준 라이브러리인 golang.org/x/time/rate 패키지를 사용하겠습니다.

### HTTP 서버에 Rate Limit 적용하기
```go
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/time/rate"
)

var limiter *rate.Limiter

// limit is middleware that rate limits requests
func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			log.Printf("Rate limit exceeded (Request ID: %v)", r.Header.Get("X-Request-Id"))
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

const (
	defaultPort  = 8080
	defaultRate  = 1
	defaultBurst = 3
)

func main() {
	port := fmt.Sprintf(":%d", *flag.Int("port", defaultPort, "port (int)"))
	r := flag.Float64("rate", defaultRate, "rate limit (float)")
	b := flag.Int("burst", defaultBurst, "burst limit (int)")
	flag.Parse()
	limiter = rate.NewLimiter(rate.Limit(*r), *b)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(http.StatusText(http.StatusOK)))
	})

	log.Printf("Server ready on %s with allowed rate of %v req/s and burst of %v reqs...", port, *r, *b)
	http.ListenAndServe(port, limit(mux))
}
```


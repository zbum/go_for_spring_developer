## Other Muxes - gin-gonic/gin
* [gin](https://github.com/gin-gonic/gin)

* martini 와 비슷한 API 를 가지고 있으며 [httprouter](https://github.com/julienschmidt/httprouter)를 활용하여 표준 라이브러리의 40배 빠른 고성능 라우터를 제공합니다.(by gin)
* 성능과 생산성을 원한다면...Gin을....
> Gorilla-Mux 는 표준 라이브러리의 http.Handler 를 구현했지만 Gin은 자체 인터페이스를 사용했기 때문에 Handler 가 호환되지 않습니다. 

## 기능
### 빠른 속도
* 기수 트리(Radix tree)를 기반으로 한 라우팅, 적은 메모리 사용량. 리플렉션 미사용. 예측 가능한 API 성능.

### 미들웨어 지원
* 수신된 HTTP 요청은 미들웨어 체인과 최종적인 액션을 통한 처리가 가능합니다. 예: Logger, 인증, GZIP 압축, DB에 메시지 전송.

### 충돌방지
* Gin은 HTTP 요청 중 발생한 panic 을 감지하고 recover 할 수 있습니다. 따라서, 서버는 항상 이용 가능한 상태를 유지합니다. 예를 들면, 발생한 panic 을 Sentry에 보고 하는 것도 가능합니다!

### JSON 유효성 검사
* Gin은 JSON 형식의 요청에 대해 파싱과 유효성 검사를 할 수 있습니다. 예를 들어, 필수값이 들어 있는지 확인 할 수 있습니다.

### 라우팅 그룹화
* 라우팅 경로를 더 좋게 정리 할 수 있습니다. 인증이 필요한지 아닌지, 버전이 다른 API인지 아닌지 등… 또한, 성능 저하 없이 그룹을 무제한 중첩할 수 있습니다.

### 에러관리
* Gin은 HTTP 요청 중에 발생한 모든 에러를 수집하는 편리한 방법을 제공합니다. 이것을 통해 미들웨어는 로그 파일, 데이터베이스에 기록하고 네트워크를 통해 전송할 수 있습니다.

### 렌더링 기능 내장
* Gin은 JSON, XML, HTML 렌더링을 위한 사용하기 쉬운 API를 제공합니다.

### 확장 가능
* 아주 쉽게 새로운 미들웨어를 만들 수 있습니다. 샘플 코드를 확인하세요.



## 설치
```
$ go get -u github.com/gin-gonic/gin
```

## 사용법
* 다음의 코드로 gin.Engine을 생성합니다. 표준 라이브러리가 아니므로 url 을 포함한 경로로 import 를 하여야 합니다.
```go
package main

import (
    "log"
    "fmt"
    "net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Handle(http.MethodGet, "/members/:member-id", func(ctx *gin.Context) {
		fmt.Fprintf(ctx.Writer, "your member id is %s", ctx.Param("member-id"))
	})
	log.Fatal(r.Run(":8080"))
}

```


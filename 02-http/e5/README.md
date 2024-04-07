## HTTP 서버 테스트
* Http 서버를 테스트하기 위해 testing 패키지를 사용할 수 있습니다. 
* Http 테스트를 위해 도움을 주는 net/http/httptest 패키지로 조금 더 쉽게 테스트를 할 수 있습니다.
 
## workshop
* e4의 워크샵에서 작성한 서버의 응답을 검증할 수 있도록, net/http/httptest 패키지를 사용해서 Http 서버의 응답 내용을 레코딩하는 테스트 코드를 작성하세요.
* 테스트 코드는 /02-http/e4/workshop/ 디렉토리에 작성하세요.

### 목표
* net/http/httptest 패키지를 사용해서 응답 http status code, 응답 body 를 레코딩해야 합니다.
* 테이블 기반 테스팅을 사용합니다.

### 알아야 할 것
* 사용법을 알기 위해 [net/http/httptest](https://pkg.go.dev/net/http/httptest) 문서를 살펴 보세요.
### Test Coverage
* 테스트를 해보면 여러분이 작성한 테스트의 커버리지를 확인할 수 있습니다. 
* Go 언어는 테스트를 실행한 후에 커버리지를 제공할 수 있는 기능을 제공합니다. 
* 커버리지를 알고 싶다면 다음의 명령어로 실행하세요.
```
$ go test ./02-http/e5  -v -cover
```

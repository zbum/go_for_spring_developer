## uber-go/FX
* uber 에서 제공하는 dependency-inject 프레임워크
* 설치
```shell
go get go.uber.org/fx@v1
```
## fx 애플리케이션 실행
```go
package main

import "go.uber.org/fx"

func main() {
  fx.New().Run()
}
```
* 애플리케이션을 실행합니다. 다음 로그가 표시되면 정상입니다.
```
[Fx] PROVIDE    fx.Lifecycle <= go.uber.org/fx.New.func1()
[Fx] PROVIDE    fx.Shutdowner <= go.uber.org/fx.(*App).shutdowner-fm()
[Fx] PROVIDE    fx.DotGraph <= go.uber.org/fx.(*App).dotGraph-fm()
[Fx] RUNNING
```

## HTTP Server 추가
* HTTP Server 를 생성하는 함수를 작성합니다.
```go
func NewHTTPServer(lc fx.Lifecycle) *http.Server {
	srv := &http.Server(Addr: ":8080")
	return srv
}
```


## 참고 자료
* https://uber-go.github.io/fx/get-started/

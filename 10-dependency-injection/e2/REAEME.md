## uber-go/FX
* uber 에서 제공하는 dependency-inject 프레임워크
* 설치
```shell
go get go.uber.org/fx@v1
```
## fx 애플리케이션 실행
* DI 컨테이너를 fx.New().Run() 으로 실행합니다.

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
## Invoke 옵션
* fx App 이 실행할때, Invoke 옵션이 있으면 즉시 실행합니다.

```go
package main

import (
	"fmt"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Invoke(PrintHello),
	).
		Run()
}

func PrintHello() {
	fmt.Println("Hello")
}
```
* 결과
```shell
[Fx] PROVIDE    fx.Lifecycle <= go.uber.org/fx.New.func1()
[Fx] PROVIDE    fx.Shutdowner <= go.uber.org/fx.(*App).shutdowner-fm()
[Fx] PROVIDE    fx.DotGraph <= go.uber.org/fx.(*App).dotGraph-fm()
[Fx] INVOKE             main.PrintHello()
Hello
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

## life cycle hook
* 생성된 http.Server를 어떻게 시작하고 종료할지 알려 주어야 합니다. 다음과 같이 수정합니다.
```go
func NewHTTPServer(lc fx.Lifecycle) *http.Server {
  srv := &http.Server{Addr: ":8080"}
  lc.Append(fx.Hook{
    OnStart: func(ctx context.Context) error {
      ln, err := net.Listen("tcp", srv.Addr)
      if err != nil {
        return err
      }
      fmt.Println("Starting HTTP server at", srv.Addr)
      go srv.Serve(ln)
      return nil
    },
    OnStop: func(ctx context.Context) error {
      return srv.Shutdown(ctx)
    },
  })
  return srv
}
 ```

## fx.Provide
* 생성자 함수를 fx 에 등록하기 위한 작업입니다. 이 작업으로는 http.Server 가 실행되지는 않습니다. 
```go
func main() {
  fx.New(
    fx.Provide(NewHTTPServer),
  ).Run()
}
```

## fx.Invoke
* http.Server 가 실행하도록 Invoke 를 사용합니다.
```go
fx.New(
    fx.Provide(NewHTTPServer),
    fx.Invoke(func(*http.Server) {}),
  ).Run()
```

## Handler 등록
* 단순한 Echo 핸들러를 작성하였습니다. 
```go
// EchoHandler is an http.Handler that copies its request body
// back to the response.
type EchoHandler struct {
	log *zap.Logger
}

// NewEchoHandler builds a new EchoHandler.
func NewEchoHandler(log *zap.Logger) *EchoHandler {
	return &EchoHandler{log: log}
}

// ServeHTTP handles an HTTP request to the /echo endpoint.
func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		h.log.Warn("Failed to handle request:", zap.Error(err))
	}
}
```

* 생성자 함수를 fx.Provide 로 등록합니다.
```go
    fx.Provide(
      NewHTTPServer,
      NewEchoHandler,
    ),
    fx.Invoke(func(*http.Server) {}),
```

* *http.ServeMux 를 생성하는 함수를 작성합니다. 
```go
// NewServeMux builds a ServeMux that will route requests
// to the given EchoHandler.
func NewServeMux(echo *EchoHandler) *http.ServeMux {
  mux := http.NewServeMux()
  mux.Handle("/echo", echo)
  return mux
}
```
* fx.Provide 에 NewServeMux 를 등록합니다. 
```go
    fx.Provide(
      NewHTTPServer,
      NewServeMux,
      NewEchoHandler,
    ),
 
```

* 마지막으로 NewHTTPServer 생성자 함수에 *ServeMux 를 의존하도록 수정하고 http.Server 의 매개변수에 설정합니다.
```go
func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux) *http.Server {
  srv := &http.Server{Addr: ":8080", Handler: mux}
  lc.Append(fx.Hook{
```

* 서버를 실행합니다.
```
[Fx] PROVIDE    *http.Server <= main.NewHTTPServer()
[Fx] PROVIDE    *http.ServeMux <= main.NewServeMux()
[Fx] PROVIDE    *main.EchoHandler <= main.NewEchoHandler()
[Fx] PROVIDE    fx.Lifecycle <= go.uber.org/fx.New.func1()
[Fx] PROVIDE    fx.Shutdowner <= go.uber.org/fx.(*App).shutdowner-fm()
[Fx] PROVIDE    fx.DotGraph <= go.uber.org/fx.(*App).dotGraph-fm()
[Fx] INVOKE             main.main.func1()
[Fx] HOOK OnStart               main.NewHTTPServer.func1() executing (caller: main.NewHTTPServer)
Starting HTTP server at :8080
[Fx] HOOK OnStart               main.NewHTTPServer.func1() called by main.NewHTTPServer ran successfully in 7.459µs
[Fx] RUNNING
```
* 서버에 요청을 전달하고 결과를 확인합니다.
```
$ curl -X POST -d 'hello' http://localhost:8080/echo
hello
```
## logger 추가
* uber 에서 개발한 zap 로거를 추가합니다. 개발용 로거를 만들어주는 생성자 함수는 zap.NewExample 입니다. (운영환경에서는 zap.NewProduction를 쓰세요.)
```go
    fx.Provide(
      NewHTTPServer,
      NewServeMux,
      NewEchoHandler,
      zap.NewExample,
    ),
```
* EchoHandler 필드에 logger 를 선언하고 NewEchoHander 생성자 함수도 logger 를 매개변수로 받도록 수정합니다.
```go
type EchoHandler struct {
  log *zap.Logger
}

func NewEchoHandler(log *zap.Logger) *EchoHandler {
  return &EchoHandler{log: log}
}
```
* EchoHandler.ServeHTTP 메소드에서 표준출력으로 처리한 것을 logger 를 사용하도록 수정하세요.
```go
func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  if _, err := io.Copy(w, r.Body); err != nil {
    h.log.Warn("Failed to handle request", zap.Error(err))
  }
}
```
* NewHTTPServer 도 동일하게 매개변수로 logger 를 받고 logger 를 사용하세요.
```go
func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux, log *zap.Logger) *http.Server {
  srv := &http.Server{Addr: ":8080", Handler: mux}
  lc.Append(fx.Hook{
    OnStart: func(ctx context.Context) error {
      ln, err := net.Listen("tcp", srv.Addr)
      if err != nil {
        return err
      }
      log.Info("Starting HTTP server", zap.String("addr", srv.Addr))
      go srv.Serve(ln)
 
```
* fx 자체도 zap logger 를 사용하도록 수정하세요.
```go
func main() {
  fx.New(
    fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
      return &fxevent.ZapLogger{Logger: log}
    }),
 
```

## 커플링 제거
* NewServeMux 는 EchoHandler에 명시적으로 의존합니다. 불필요한 의존성을 제거해 보겠습니다.
1. Route 인터페이스를 정의 합니다. http.Handler 의 Path 를 알려주는 용도입니다.
```go
// Route is an http.Handler that knows the mux pattern
// under which it will be registered.
type Route interface {
  http.Handler

  // Pattern reports the path at which this is registered.
  Pattern() string
}
```

2. EchoHandler 가 Route 인터페이스를 구현하도록 합니다.
```go
func (*EchoHandler) Pattern() string {
  return "/echo"
}
```

3. main() 함수에서 NewEchoHandler 는 Route로 제공되어야 함을 명시합니다.
```go
    fx.Provide(
      NewHTTPServer,
      NewServeMux,
      fx.Annotate(
        NewEchoHandler,
        fx.As(new(Route)),
      ),
      zap.NewExample,
    ),
```

4. NewServeMux 가 Route 를 매개변수로 받고 패턴을 사용하도록 수정합니다.
```go
// NewServeMux builds a ServeMux that will route requests
// to the given Route.
func NewServeMux(route Route) *http.ServeMux {
  mux := http.NewServeMux()
  mux.Handle(route.Pattern(), route)
  return mux
}
```

## 다른 핸들러 등록
1. HelloHandler 를 작성합니다.
```go
// HelloHandler is an HTTP handler that
// prints a greeting to the user.
type HelloHandler struct {
  log *zap.Logger
}

// NewHelloHandler builds a new HelloHandler.
func NewHelloHandler(log *zap.Logger) *HelloHandler {
  return &HelloHandler{log: log}
}
```
2. Route 인터페이스를 구현하고 기능도 구현합니다.
```go
func (*HelloHandler) Pattern() string {
  return "/hello"
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  body, err := io.ReadAll(r.Body)
  if err != nil {
    h.log.Error("Failed to read request", zap.Error(err))
    http.Error(w, "Internal server error", http.StatusInternalServerError)
    return
  }

  if _, err := fmt.Fprintf(w, "Hello, %s\n", body); err != nil {
    h.log.Error("Failed to write response", zap.Error(err))
    http.Error(w, "Internal server error", http.StatusInternalServerError)
    return
  }
}
 
```
3. NewEchoHandler 다음에 NewHelloHandler를 Route 로 등록합니다.
```go
      fx.Annotate(
        NewEchoHandler,
        fx.As(new(Route)),
      ),
      fx.Annotate(
        NewHelloHandler,
        fx.As(new(Route)),
      ),
```
4. 이 상태에서 실행하면 실행은 실패합니다. Fx 는 스프링과 같이 동일한 타입의 인스턴스를 애너테이션이 없는 상태로 허용하지 않습니다.
5. 각 핸들러에 애너테이션을 설정합니다.
```go
      fx.Annotate(
        NewEchoHandler,
        fx.As(new(Route)),
        fx.ResultTags(`name:"echo"`),
      ),
      fx.Annotate(
        NewHelloHandler,
        fx.As(new(Route)),
        fx.ResultTags(`name:"hello"`),
      ),
 
```
6. NewServeMux 를 다음과 같이 수정합니다.
```go
// NewServeMux builds a ServeMux that will route requests
// to the given routes.
func NewServeMux(route1, route2 Route) *http.ServeMux {
  mux := http.NewServeMux()
  mux.Handle(route1.Pattern(), route1)
  mux.Handle(route2.Pattern(), route2)
  return mux
} 
```
7. main 함수에서도 이 두 이름값을 설정합니다.
```go
    fx.Provide(
      NewHTTPServer,
      fx.Annotate(
        NewServeMux,
        fx.ParamTags(`name:"echo"`, `name:"hello"`),
      ),
```

## 더 많은 핸들러 등록
* 핸들러가 추가될때마다 NewServeMux 의 인자를 조정해야 하는 문제점이 있습니다. 
1. NewServeMux 를 Route 슬라이스를 받도록 수정합니다.
```go
func NewServeMux(routes []Route) *http.ServeMux {
  mux := http.NewServeMux()
  for _, route := range routes {
    mux.Handle(route.Pattern(), route)
  }
  return mux
}
```
2. main 함수에서 NewServeMux 가 route group 의 슬라이스를 받도록 애너테이션을 설정합니다.
```go
    fx.Provide(
      NewHTTPServer,
      fx.Annotate(
        NewServeMux,
        fx.ParamTags(`group:"routes"`),
      ),

```
3. 이 그룹에 Route 구현체를 추가하는 AsRoute 라는 함수를 작성합니다.
```go
// AsRoute annotates the given constructor to state that
// it provides a route to the "routes" group.
func AsRoute(f any) any {
  return fx.Annotate(
    f,
    fx.As(new(Route)),
    fx.ResultTags(`group:"routes"`),
  )
}
```
4. AsRoute 를 사용하여 NewEchoHandler, NewHelloHandler 생성자 함수를 route 그룹으로 설정합니다.
```go
    fx.Provide(
      AsRoute(NewEchoHandler),
      AsRoute(NewHelloHandler),
      zap.NewExample,
    ),
```
## 결론
* 튜토리얼 방식으로 따라하기 를 해 보았습니다. 
1. Fx 응용 프로그램을 처음부터 시작하는 방법
2. 새로운 종속성을 주입하고 기존 종속성을 수정하는 방법
3. 인터페이스를 사용하여 구성 요소를 분리하는 방법
4. 명명된 값을 사용하는 방법
5. 값 그룹을 사용하는 방법

## 자료 출처
* https://uber-go.github.io/fx/get-started/

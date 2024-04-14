## Logging
## 표준 라이브러리
* 표준 라이브러리의 "log" 패키지를 사용합니다.
```go
package main
 
import "log"
 
func main() {
    log.Println("Logging")
}
```
* 결과
```
2024/04/14 18:34:44 Logging
```
## logger 타입
* log 패키지는 logger 라는 구조체 타입을 제공합니다.
```go
type Logger struct {
	outMu sync.Mutex
	out   io.Writer // destination for output

	prefix    atomic.Pointer[string] // prefix on each line to identify the logger (but see Lmsgprefix)
	flag      atomic.Int32           // properties
	isDiscard atomic.Bool
}
```
* 표준 라이브러리에서 `log.Println("Logging")` 와 같이 최상위 함수를 직접 실행하는 방식은 `표준 로거`를 사용하는 방식입니다.
```go
var std = New(os.Stderr, "", LstdFlags)

func Println(v ...any) {
	std.output(0, 2, func(b []byte) []byte {
		return fmt.Appendln(b, v...)
	})
}
```

## logger 생성
* 새로운 로거(Logger)를 만들기 위해 log.New() 함수를 사용할 수 있습니다.
* log.New()는 3개의 파라미터를 받아들이는데, 
  * 첫번째는 io.Writer 인터페이스를 지원하는 타입으로 표준콘솔출력(os.Stdout), 표준에러(os.Stderr), 파일포인터 혹은 io.Writer를 지원하는 모든 타겟이 사용될 수 있습니다. 
  * 두번째 파라미터는 로그출력의 가장 처음에 적는 Prefix로서 프로그램명,카테고리 등을 기재할 수 있습니다. 
  * 세번째 파라미터는 로그플래그로 표준플래그(log.LstdFlags), 날짜플래그(log.Ldate), 시간플래그(log.Ltime), 파일위치플래그(log.Lshortfile, log.Llongfile) 등을 | (OR 연산자)로 묶어 지정할 수 있습니다.
* 예제
```go
package main
 
import (
    "log"
    "os"
)
 
var myLogger *log.Logger
 
func main() {
    myLogger = log.New(os.Stdout, "INFO: ", log.LstdFlags)
 
    //....
    run()
 
    myLogger.Println("End of Program")
}
 
func run() {
    myLogger.Print("Test")
}
```

## file logger
* 컨테이너 환경이 아닌 경우, 로그파일에 로그를 출력하기 원합니다.
* 로그파일을 오픈하고 파일포인터를 log.New()의 첫번째 파라미터에 넣어 주면 됩니다. (주: os.File은 io.Writer를 구현)
    * file 은 Write Only, Append Only 로 열어 주면 됩니다.
```go
	fpLog, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()

    // 로그플래그는 날짜/시간(log.Ldate|log.Ltime), 짧은 파일명/라인수(Lshortfile)를 함께 출력
	myFileLogger = log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
```
* 결과 (logfile.txt)
```
INFO: 2016/01/15 15:30:53 test.go:27: Test
INFO: 2016/01/15 15:30:53 test.go:23: End of Program
```

* logger 를 생성하지 않고 표준로거가 파일을 사용하도록 하려면 `log.SetOutput(fpLog)` 함수를 사용하면 됩니다.
```go
    fpLog, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        panic(err)
    }
    defer fpLog.Close()
 
    // 표준로거를 파일로그로 변경
    log.SetOutput(fpLog)
```

## 멀티 로깅
* 표준 출력과 파일에 모두 로그를 표시해야 한다면 io.MultiWriter 함수를 사용합니다.

## 구조화 로깅(structured logging)
* log/slog 패키지는 Go 1.21 부터 제공됩니다.
* slog 로 남기는 로그는 키,값 쌍으로 로그를 남기게 되어 필터링, 분석, 검색에 용이합니다.

## slog 사용법
* log/slog 를 임포트 하고 slog 패키지의 함수를 사용합니다.
* slog.Info 라는 최상위 함수를 사용합니다.
```go
package main

import (
    "log/slog"
    "os"
)


func main() {
    slog.Info("hello, world", "user", os.Getenv("USER"))
}
```
* 결과
```
2024/04/14 19:05:22 INFO hello world user=nhn
```

## slog TextHandler
* slog를 New() 함수를 이용해서 직접 생성해 보겠습니다.
* TextHandler 는 key=value 형태로 로그를 남기도록 처리합니다. 
```go
logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
logger.Info("hello, world", "user", os.Getenv("USER"))
```

## slog JSONHandler
* JSONHandler 는 json 형태로 로그를 남기도록 처리합니다.
```go
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("hello, world", "user", os.Getenv("USER"))
```
* 출력
```
{"time":"2024-04-14T19:13:41.523079+09:00","level":"INFO","msg":"hello world","user":"nhn"}
```

## 로그 속성(LogAttrs)
* 지금까지 사용한 Info 등의 함수는 사용하기 편리하지만, 자주 실행되는 로그 문의 경우에는 Attr 유형을 사용하고 LogAttrs 메서드를 호출하는 것이 더 효율적입니다. 
* 메모리 할당을 최소화합니다.
* LogAttrs에 대한 이 호출은 위와 동일한 출력을 생성하지만 더 빠릅니다.
```go
slog.LogAttrs(context.Background(), slog.LevelInfo, "hello, world",
    slog.String("user", os.Getenv("USER")))
```



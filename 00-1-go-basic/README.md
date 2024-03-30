## Go 설치
* Go 언어는 공식 사이트(https://go.dev/doc/install)에서 다운로드 받아 설치하는 것을 권장합니다.
  * 패키지 매니저는 버전 업데이트가 너무 느리거나 중단한 경우가 많습니다. 
1. https://go.dev/dl/ 에서 PC 플랫폼용 바이너리를 찾아 다운로드 합니다. 
2. Microsoft Windows, macOS는 설치 패키지로 제공됩니다. 따라서, 다운로드한 파일을 실행하여 설치합니다. 
3. Linux 배포본은 다음의 명령어를 사용합니다.
* 기존 go 바이너리 삭제 및 신규 패키지 설치
```shell
$ sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.22.1.linux-amd64.tar.gz
```
* 실행 경로 추가
```shell
export PATH=$PATH:/usr/local/go/bin
```
4. 설치후 다음의 명령어로 원하는 go version 설치를 확인 합니다.
```shell
$ go version
```
* 출처: https://go.dev/doc/install

## 프로젝트 생성/빌드/실행
### 프로젝트 생성 
1. 프로젝트를 개발할 디렉토리를 생성하고 해당 디렉토리로 이동합니다. 
```shell
$ cd 
$ mkdir hello
$ cd hello
```
2. 프로젝트를 `go mod`로 생성합니다. (mod는 모듈이라는 말입니다.)
```shell
$ go mod init example/hello
go: creating new go.mod: module example/hello
```
3. 작업한 디렉토리에 hello.go 파일을 아래와 같이 생성합니다. (Copy & Paste)
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```
### 프로젝트 빌드/실행
1. `go build`로 프로젝트를 빌드합니다. 빌드하면 hello (or hello.exe) 라는 실행파일이 생성되어 있습니다.
```shell
$ go build -o hello hello.go
```
2. hello 를 실행합니다.
```shell
$ ./hello
Hello, World!
```

### 컴파일 없이 실행하기
* python과 같은 인터프리터 언어의 사용자 경험을 제공합니다.
```shell
$ go run hello.go
Hello, World!
```
* `go run` 커멘드는 Go 패키지를 빌드하고 임시 실행 파일을 만든 후 해당파일을 실행하고 실행이 끝나면 지워버립니다.
* go의 컴파일 속도는 매우 빨라서 스크립트 언어로 오해할 정도입니다.

### 외부 패키지 호출
* 위 프로그램에서 출력하는 내용을 두레이메신저로 발송하는 작업을 해보겠습니다.
1. https://pkg.go.dev 사이트에서 dooray-go를 검색합니다.
2. 검색결과 상단에서 "github.com/dooray-go/dooray" 를 복사합니다.
3. 다음과 같이 hello.go 를 수정합니다. (배우지 않은 문법은 신경쓰지 않습니다.!!)
```go
package main

import "github.com/dooray-go/dooray"

func main() {
    dooray.PostWebhook(
        "https://hook.dooray.com/services/3036349505739914786/3770555218093552684/autJQopeRTiVWUNxrgfaFA",
        &dooray.WebhookMessage{
            BotName: "Manty",
            Text:    "Hello, World!"},
    )
}
```
4. 새로운 모듈을 go.sum, go.mod 에 추가합니다. 추가하기 위한 명령어는 아래와 같습니다.
```shell
$ go mod tidy
```
5. 수정한 코드를 실행합니다.
```shell
$ go run hello.go
```
* 출처 : https://go.dev/doc/tutorial/getting-started

## godoc 유틸리티

## 변수 선언


## 오류 처리
* exception 없음
  https://github.com/astaxie/build-web-application-with-golang/blob/master/en/11.1.md
## 데이터 모델
* Array, Slice, Map
## Array

## Slice
* ArrayList 와 매우 흡사함

## Map
* HashMap 과 흡사함
* ConcurrentHashMap 처럼 동시성을 지원하지 않음.

## type 정의

## 인터페이스

## 신비로운 키워드 
* defer
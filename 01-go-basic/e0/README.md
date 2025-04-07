## 실행 프로젝트 생성/빌드/실행
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

> 출처 : https://go.dev/doc/tutorial/getting-started
>


## 함수
* Go 언어의 함수는 func 키워드로 시작하고 함수의 이름, 시그니쳐, 구현 으로 구성됩니다.
* 함수와 변수의 첫 글자가 소문자이면 자바의 private와 같은 효과가 있어 같은 패키지에서만 접근이 가능합니다.
* 패키지 이름은 대소문자에 따른 영향이 없지만 관례상 소문자만 사용합니다.

```go
func test() string {
    return "success!!"
}
```

* Go 언어의 함수는 2개 이상의 값을 반환 할 수 있습니다.
```go
func test() (string, error) {
    var err = someFunction()
    if err != nil {
        return "", err
    }   
    return "success!!", nil
}
```
* 함수의 인자는 타입을 명시해야 합니다.
```go

```
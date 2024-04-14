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
> 출처: https://go.dev/doc/install
## IDE
### IDE - GoLand
* https://www.jetbrains.com/ko-kr/go/
* 하지만 유료..

### IDE - intellij ultimate
* 다음 플러그인을 설치 합니다.
* 설정 > 플러그인 > go 검색 후 설치
* https://plugins.jetbrains.com/plugin/9568-go

### IDE - Visual Studio Code
* https://marketplace.visualstudio.com/items?itemName=golang.go

### IDE 참고자료
* https://go.dev/wiki/IDEsAndTextEditorPlugins

## 실습 및 교재 다운로드
```
git clone https://github.com/zbum/go_for_spring_developer.git
```

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

## 중요한 형식과 코딩 규칙
* Go는 실수를 방지하도록 엄격한 코딩 규칙을 적용합니다.
* Go가 제공하는 표준 도구(gofmt)를 제공해 형식은 자동으로 조정할 수 있습니다.
### 코딩 규칙
* 패키지를 임포트 하면 해당 패키지 기능을 사용해야 합니다.
* 변수를 선언했다면 반드시 사용해야 합니다..
* 중괄호를 사용하는 형식은 1가지 방법 밖에 없습니다.
```go
package main

import(
    "fmt"
)

// 컴파일 오류가 발생합니다.
func main()
{
    fmt.Println("Go has strict rules for curly braces!")
}
```
* 코드가 한문장이거나 아예 존재하지 않더라도 코드 블록은 중괄호로 감싸야 합니다.
* 함수는 여러 개의 값을 반환할 수 있습니다.
* 같은 종류의 데이터일지라도 다른 데이터 타입으로 자동 변환되지 않습니다. (정수를 부동소수점으로 자동변환하지 않습니다.)








## 정리
* 웹 애플리케이션을 구현하기 위한 Go 언어의 기본 기능을 살펴 보았습니다. 물론 Go 언어에 대해 모두 배운것은 아닙니다,
* Go 언어를 이용해서 프로젝트를 생성/실행 하는 과정을 살펴 보았습니다.
* 여러가지 타입을 정의 하고, 타입의 변수를 만드는 작업을 확인 하였습니다.
* 오류처리와 포인터에 대한 기본적인 지식을 알게 되었습니다.
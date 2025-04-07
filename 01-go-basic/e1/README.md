
## 패키지
* 패키지는 package 키워드로 정의합니다.
* 실행가능한 애플리케이션을 만들려면 main 패키지가 있어야 합니다.
* 다른 Go 패키지를 사용할때는 import를 사용합니다.
* 표준 라이브러리는 간단히 패키지 이름으로 가져 올 수 있습니다.
```go
import "fmt"
```
* 외부 패키지는 인터넷 주소를 써서 가져와야 합니다.
```go
import "github.com/dooray-go/dooray"
```
> 중요 : Go의 패키지와 Java의 패키지는 다릅니다.
> Go의 패키지 내에 여러 go 파일이 있는 경우, 각 파일에서 동일한 함수, 변수 선언이 불가능합니다.
> Go의 패키지는 java 파일 하나로 생각하면 조금 비슷합니다.
> 계속 진행해 보면 그다지 어려운 부분은 아니지만 처음에 조금 당황할 수 있습니다.


## 외부 패키지 호출
* 위 프로그램에서 출력하는 내용을 두레이메신저로 발송하는 작업을 해보겠습니다.
1. https://pkg.go.dev 사이트에서 dooray-go를 검색합니다.
2. 검색결과 상단에서 "github.com/dooray-go/dooray" 를 복사합니다.
3. 다음과 같이 hello.go 를 수정합니다. (배우지 않은 문법은 신경쓰지 않습니다.!!)
### (e1)
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

### workshop(e1/w1)
* 외부 패키지는 "github.com/agrison/go-commons-lang/stringUtils" 을 사용합니다.
    * https://pkg.go.dev/github.com/agrison/go-commons-lang/stringUtils 에서 제공하는 함수와 사용법을 확인합니다.
* 패키지를 다운로드 하기 위해 다음 명령어를 실행합니다.
```shell
go get -u github.com/agrison/go-commons-lang/stringUtils
```
* main 함수에서 "Go for Spring Developers!!" 문자열을 뒤집습니다.
* 뒤집은 문자열을 표준 출력에 표시합니다.


## 모듈 배포하기 (라이브러리)
### 개요
* Go 언어는 패키지 단위로 라이브러리를 배포합니다.
* 라이브러리 패키지는 go.mod 파일을 포함해야 합니다.

### 모듈 생성
* 모듈을 생성하기 위해서는 go mod init 명령어를 사용합니다.
```shell
$ mkdir -p zbum/stringutils
$ cd zbum/stringutils
$ go mod init github.com/zbum/stringutils
```
* go.mod 파일이 생성됩니다. 이것으로 모듈 생성은 끝났습니다. 

### 모듈 개발
* 입력받은 문자열을 대문자로 바꾸는 함수를 만들겠습니다.
* 아래 코드를 stringutils.go 파일에 작성합니다.
```go
package stringutils

func ToUpperCase(str string) string {
    return strings.ToUpper(str)
}
```
* 테스트 코드도 stringutils_test.go 파일에 작성합니다. 
```go
package stringutils

import "testing"

func TestToUpperCase(t *testing.T) {
    data := "Hello, world."
    want := "HELLO, WORLD."
    if got := ToUpperCase(data); got != want {
        t.Errorf("ToUpperCase(%q) = %q, want %q", data, got, want)
    }
}
```
* `go test` 명령어로 테스트를 실행합니다.
```shell
$ go test
PASS
ok  	github.com/zbum/stringutils	0.655s
```

### 모듈 게시(Publish)
* 모듈을 배포하기 위해서는 git 저장소를 만들고 초기 커밋을 합니다.
```shell
$ git init
$ git add .
$ git commit -m "stringutils: initial commit"
$
```
* 의미론적 버전 (Semantic Versioning) 규칙에 따라 버전을 태그합니다.
    * 의미론적 버전은 vMAJOR.MINOR.PATCH 로 구성됩니다.
    * API에 이전 버전과 호환되지 않는 변경을 할 때 MAJOR 버전을 높이세요 . 이는 절대적으로 필요한 경우에만 수행해야 합니다.
    * 종속성을 변경하거나 새로운 함수, 메서드, 구조체 필드 또는 유형을 추가하는 등 API에 이전 버전과의 호환성을 유지하는 변경 사항을 적용하면 MINOR 버전을 높입니다 .
    * 버그 수정처럼 모듈의 공개 API나 종속성에 영향을 주지 않는 사소한 변경을 한 후에는 PATCH 버전을 올립니다 .
* v1.0.1-alpha 또는 v2.2.2-beta.2 의 형태로 사전 릴리즈 버전을 지정할 수 있습니다. 

* 초기 버전 배포 (v0.1.0)
```shell
## go mod tidy 로 필요 없는 모듈 종속성 제거
$ go mod tidy

## 모든 테스트 통과 확인
$ go test ./...

## git tag 를 사용하여 새버전을 태그 합니다.
$ git tag v0.1.0

## 새로운 태그를 원본 저장
$ git push origin v0.1.0
```
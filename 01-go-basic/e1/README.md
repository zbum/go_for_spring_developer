
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

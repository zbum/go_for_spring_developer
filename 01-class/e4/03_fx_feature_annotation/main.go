package main

import (
	"github.com/google/go-github/v57/github"
	"go.uber.org/fx"
	"net/http"
)

// 함수와 값을 fx.Provide, fx.Supply, fx.Invoke, fx.Decorate, fx.Replace 에 전달하기 전에
// fx.Annotate 함수를 이용해서 주석을 설정할 수 있습니다.

// 이를 통해 매개변수 또는 결과 개체를 사용하기 위해 함수를 수동으로 래핑하지 않고도 일반 Go 함수를 재사용하여 다음을 수행할 수 있습니다.

// ## 함수에 주석달기 (Annotating)
// 1. fx.Provide, fx.Invoke, fx.Decorate 에 전달하는 함수가 있습니다.

func main() {
	fx.
		New(
			fx.Provide(
				fx.Annotate(
					NewHTTPClient,
					//fx.ResultTags(`name:"Client"`),
					fx.As(new(HTTPClient)),
				),
				NewGitHubClient,
			),
		).
		Run()
}

func NewHTTPClient() (*http.Client, error) {
	return &http.Client{}, nil
}

// 2. NewHTTPClient 함수를 fx.Annotate 로 래핑합니다.

// 3. fx.Annotate 내부에서 주석을 전달합니다.

// ## 스트럭트를 인터페이스로 캐스팅하기

// 1. NewGitHubClient 함수는 NewHTTPClient의 결과를 소비합니다.
func NewGitHubClient(client HTTPClient) *github.Client {
	return nil
}

// 2. 이 함수는 Fx 애플리케이션에 제공(Provide)되어 있습니다.

// ## 작업
// 1. NewHTTPClient 가 반환하는 *http.Client와 일치하는 인터페이스를 선언합니다.
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// 컴파일 타임에 인터페이스 호환을 점검 합니다.
var _ HTTPClient = (*http.Client)(nil)

// 2. 스트럭트 대신 인터페이스를 사용하도록 NewGitHubClient를 수정합니다.

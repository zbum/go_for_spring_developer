package _2_fx_feature_result

import (
	"go.uber.org/fx"
	"net/http"
	"net/url"
)

// 결과 객체는 특정 함수나 메소드의 결과값을 전달하기 위한 객체 입니다.
// 파라미터 객체와 마찬가지로, 이 객체는 일반적으로 특정 함수를 위해 배타적으로 정의 합니다. 그말은 다른 함수와 공유하지 않는다는 말입니다.
// Fx 에서는 파라미터 객체는 배타적으로 내보내기한 필드를 포함하고 있습니다. 그리고 fx.Out 으로 태깅해 두어야 합니다.

// ## 결과 오브젝트 사용하기
// Fx 의 결과 오브젝트를 사용하려면 다음의 순서를 따르세요.

// 1. Client struct 가 제공됩니다.
type Client struct {
	url  url.URL
	http *http.Client
}

type ClientConfig struct {
	URL url.URL
}

// 2. Result 접미사를 스트럭트를 정의합니다.
// 3. fx.Out 을 embed 합니다.
// 4. 생성자가 생산하는 값을 필드로 추가합니다.
type ClientResult struct {
	fx.Out

	Client    *Client
	Inspector *Inspector
}

// 4. NewClient 생성함수를 작성합니다.
// ClientResult 를 반환 타입으로 작성합니다. 반드시 `by value` 로 작업합니다.

func NewClient() (ClientResult, error) {
	client := &Client{}
	return ClientResult{Client: client}, nil
}

// 결과 객체를 함수에 포함하면, FX 의 기능을 사용할 수 있습니다.

// ## 새로운 결과 필드 추가

// 파라미터 오브젝트에 새 필드를 추가하여 생성자용 파리미터를 추가할 수 있습니다.
// 1. Inspector 라는 새로운 스트럭트를 작성합니다.
type Inspector struct {
}

// 2. Inspector 를 ClientResult 에 새로운 필드로 추가합니다.

package _1_fx_feature_parameter

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
	"net/url"
)

// Parameter Object 는 특정 함수나 메소드에 파라미터를 전달하기 위한 객체 입니다.
// 이 객체는 일반적으로 특정 함수를 위해 배타적으로 정의 합니다. 그말은 다른 함수와 공유하지 않는다는 말입니다.
// Fx 에서는 파라미터 객체는 배타적으로 내보내기한 필드를 포함하고 있습니다. 그리고 fx.In 으로 태깅해 두어야 합니다.

// ## 파라미터 오브젝트 사용하기
// Fx 의 파라미터 오브젝트를 사용하려면 다음의 순서를 따르세요.

// 1. Client struct 가 제공됩니다.
type Client struct {
	url  url.URL
	http *http.Client
}

type ClientConfig struct {
	URL url.URL
}

// 2. Params 접미사를 스트럭트를 정의합니다.
// 3. fx.In 을 embed 합니다.
// 4. Client 가 필요로 하는 속성을 추가 합니다.
type ClientParams struct {
	fx.In

	Config     ClientConfig
	HTTPClient *http.Client
	Logger     *zap.Logger `optional:"true"`
}

// 4. NewClient 생성함수를 작성합니다.
// ClientParams 를 매개변수로 작성합니다. 반드시 `by value` 로 작업합니다.

func NewClient(p ClientParams) (*Client, error) {
	log := p.Logger
	if log == nil {
		log = zap.NewNop()
	}
	return &Client{
		url:  p.Config.URL,
		http: p.HTTPClient,
	}, nil
}

// 파라미터 오브젝트를 함수에 포함하면, FX 의 기능을 사용할 수 있습니다.

// ## 새로운 파라미터 추가

// 파라미터 오브젝트에 새 필드를 추가하여 생성자용 파리미터를 추가할 수 있습니다.
// 1. 새로운 필드를 추가합니다. 하위 호환성을 보장하기 위해 Optional 로 생성합니다.

// 2. 생성자에서 이 필드를 소비합니다. 이 필드가 nil 인 경우를 처리해야 합니다.

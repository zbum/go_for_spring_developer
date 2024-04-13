## 설정 외부화
* 스프링 부트는 application.properties, application.yaml 등으로 설정을 외부화 합니다. 
* 실행시 spring.profiles.active 와 같은 옵션으로 프로파일에 하댕하는 설정을 가지고 실행합니다.
* Go 언어에서 자체가 지원하는 기능은 없지만 viper와 같은 외부 라이브러리를 활용할 수 있습니다.
  * https://github.com/spf13/viper?tab=readme-ov-file

## 설치
```shell
go get github.com/spf13/viper
```
## Viper
* 다음의 기능을 제공합니다. 

1. 기본값 설정
2. JSON, TOML, YAML, HCL, envfile 및 Java properties 설정 파일에서 읽기
3. 실시간 변경 감시 및 구성 파일 다시 읽기(선택 사항)
4. 환경 변수에서 읽기
5. 원격 구성 시스템(etcd 또는 Consul)에서 읽기 및 변경 감시
6. 명령행 플래그에서 읽기
7. 버퍼에서 읽기
8. 명시적 값 설정

## 기본값 설정
* 기본값을 제공하는 것은 좋은 설정입니다.
* viper 라이브러리에서 기본값을 설정하는 것은 viper.SetDefault 라는 함수를 사용합니다.
```go
func setDefaults() {
    viper.SetDefault("server.port", "8080")
}
```

## 설정 파일 읽기
* Viper 는 여러 디렉토리에서 설정파일을 검색합니다.
* 설정파일을 viper 객체당 1개의 파일만 읽을 수 있습니다.
* 다음은 여러 디렉토리에서 application.properties 를 읽는 설정의 예입니다.
```go
viper.SetConfigName("application") // 확장자를 제외한 파일이름
viper.SetConfigType("properties") // 설정파일의 형식
viper.AddConfigPath("/etc/manty/")   // 설정파일을 찾을 디렉토리
viper.AddConfigPath("$HOME/.manty")  // 여러번 추가할 수 있습니다. 
viper.AddConfigPath(".")             // 실행 경로
err := viper.ReadInConfig() // 파일을 찾아 읽기
if err != nil { 
	panic(fmt.Errorf("fatal error config file: %w", err))
}
```

## embed fs 에서 설정파일 읽기
* 스프링부트의 설정파일은 jar 내부에 포함됩니다. 
* 이와 비슷하게 실행파일 내에 설정파일을 포함하고 싶다면 go embed 를 사용해야 합니다. 
* 그럴때 viper 에서 설정파일을 읽는 방법입니다. 
* 이 방법을 사용하면 실시간 변경을 감지 할 수 없습니다.
```go
	viper.SetConfigType("yaml") // 설정파일의 형식은 반드시 입력해야 합니다.
	err = viper.ReadConfig(bytes.NewReader(embedFile)) // ReadConfig 함수로 파일을 읽기
```

## 설정파일 쓰기



## 참고자료
* https://github.com/spf13/viper?tab=readme-ov-file
* Dooray golang prototype
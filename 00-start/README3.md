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

## IDE - GoLand
* https://www.jetbrains.com/ko-kr/go/
* 하지만 유료..

## IDE - intellij ultimate
* 다음 플러그인을 설치 합니다.
* 설정 > 플러그인 > go 검색 후 설치
* https://plugins.jetbrains.com/plugin/9568-go

## IDE - Visual Studio Code
* https://marketplace.visualstudio.com/items?itemName=golang.go

## IDE 참고자료
* https://go.dev/wiki/IDEsAndTextEditorPlugins

## 실습 및 교재 다운로드
```
git clone https://github.com/zbum/go_for_spring_developer.git
```
* git clone 후 다음 명령을 수행하여 라이브러리를 다운로드 받습니다. 
```
$ go mod tidy
```
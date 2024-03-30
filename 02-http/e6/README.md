## Other Muxes - gorilla/mux
* https://github.com/gorilla/mux

* gorilla/mux 패키지는Go 웹 서버를 구축하기 위한 강력한 HTTP 라우터 및 URL matcher 입니다.

## 기능
* gorilla/mux 패키지는 요청을 해당 핸들러와 매핑하기 위해 요청 라우터(request router)와 디스패처를 구현합니다.
* mux라는 이름은 "HTTP 요청 멀티플렉서"를 나타냅니다. 표준 라이브러리의 http.ServeMux와 같이,  mux.Router는 요청을 저장한 경로 목록과 매칭시키고 URL 또는 기타 조건과 일치하는 경로에 대한 핸들러를 호출합니다. 
<br />
* 주요 기능은 다음과 같습니다.
1. 표준라이브러리의 http.ServeMux와 호환되도록  http.Handler 인터페이스를 구현합니다. 
2. 요청은 URL 호스트, 경로, 경로 접두사, schemes(http/https), 헤더 및 쿼리 값, HTTP 메소드를 기반으로 합니다.
3. 커스텀 매쳐를 사용하여 일치될 수 있습니다.
4. URL 호스트, 경로 및 쿼리 값에는 선택적 정규 표현식이 포함된 변수가 있을 수 있습니다.
5. 등록된 URL은 리소스에 대한 참조를 유지하는 데 도움이 되도록 작성하거나 "역방향"으로 만들 수 있습니다.
6. 경로는 하위 라우터로 사용될 수 있습니다. 중첩된 경로는 상위 경로가 일치하는 경우에만 테스트됩니다. 이는 호스트, 경로 접두사 또는 기타 반복되는 속성과 같은 공통 조건을 공유하는 경로 그룹을 정의하는 데 유용합니다. 보너스로 이는 요청 일치를 최적화합니다.

## 설치
```
go get -u github.com/gorilla/mux
```

## 예제




## 유용한 함수
mux.Vars(r)



## 실습
* 다음의 RESTful 서버를 개발하세요.
```http request
GET /hello

```
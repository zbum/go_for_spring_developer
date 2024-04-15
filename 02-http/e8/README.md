## Other Muxes - gorilla/mux
* [gorilla-mux](https://github.com/gorilla/mux)

* gorilla/mux 패키지는Go 웹 서버를 구축하기 위한 강력한 HTTP 라우터 및 URL matcher 입니다.
* Go 언어 1.22 버전 부터 http 패키지의 기능이 많이 보강되었지만 좀 더 사용하기 쉽고 많은 기능을 제공합니다.

## 기능
* gorilla/mux 패키지는 요청을 해당 핸들러와 매핑하기 위해 요청 라우터(request router)와 디스패처를 구현합니다.
* 표준 라이브러리의 http.ServeMux와 같이, mux.Router는 요청을 저장한 경로 목록과 매칭시키고 URL 또는 기타 조건과 일치하는 경로에 대한 핸들러를 호출합니다. 
<br />
* 주요 기능은 다음과 같습니다.
1. 표준라이브러리의 http.ServeMux와 호환되도록 http.Handler 인터페이스를 구현(mux.Router)합니다. 
2. 요청은 URL 호스트, 경로, 경로 접두사, schemes(http/https), 헤더 및 쿼리 값, HTTP 메소드를 기반으로 합니다.
3. 서브 라우터를 만들 수 있습니다. 중첩된 라우터는 상위 라우터가 일치하는 경우에만 처리됩니다. 이는 호스트, 경로 접두사 또는 기타 반복되는 속성과 같은 공통 조건을 공유하는 경로 그룹을 정의하는 데 유용합니다.

## 설치
```
go get -u github.com/gorilla/mux
```

## 사용법
* 다음의 코드로 Router 를 생성합니다. 표준 라이브러리가 아니므로 url 을 포함한 경로로 import 를 하여야 합니다.
```go
package main

import (
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
}
```
### 라우터 설정
```go
    // "/" 요청을 handlerFunction 으로 전달합니다.
    r.HandleFunc("/", handlerFunction)

    // "POST" "/url" 요청을 handlerPostFunction 으로 전달합니다.
    r.HandleFunc("/url", handlerPostFunction).Methods(http.MethodPost)

    // NotFoundHandler 라는 속성에 핸들러를 설정할 수 있습니다. 아무 곳에도 속하지 않는 요청은 이 핸들러가 처리합니다.
    r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { 
        http.Error(w, "Not Found by Manty", http.StatusNotFound)
    })

    // MethodNotAllowedHandler 라는 속성에 핸들러를 설정할 수 있습니다. url 은 일치하지만 Http 메소드가 일치하지 않으면  이 핸들러가 처리합니다.
    r.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        http.Error(w, "Method Not Allowed by Manty", http.StatusMethodNotAllowed)
    })
```

### 서브라우터
* 부모 라우터에서 설정한 호스트, 경로 접두사(PathPrefix), HTTP 메소드 를 기반으로 하위 라우터에서 그 속성을 이어 받는 라우터를 구성할때 사용합니다.
* PathPrefix 를 이용한 예
```go
// PathPrefix 를 이용해서 Subrouter 를 생성합니다.
groupsMux := r.PathPrefix("/groups").Subrouter()

// groupsMux에 "/{id}" Path 를 적용하면 "/groups/{id}" 형태의 Path 에 대한 라우칭을 처리할 수 있습니다.
groupsMux.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Requested Group id : %s\n", mux.Vars(r)["id"])
}).Methods(http.MethodGet)

groupsSubRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Requested All Groups : \n")
}).Methods(http.MethodGet)
````

* Method 를 이용한 예
```go
deleteSubRouter := r.Methods(http.MethodDelete).Subrouter()
deleteSubRouter.HandleFunc("/delete/{id}", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "DELETE Requested delete id : %s\n", mux.Vars(r)["id"])
})
```

> SpringMvc 에서 Class 수준의 @RequestMapping 을 선언하는 것과 비슷한 역할을 합니다.

## 유용한 함수
### mux.Vars(r *http.Request)
* HTTP 요청의 변수를 map[string]string 타입으로 반환합니다.
* SpringMvc 의 @PathVariable 와 동일한 기능을 제공합니다. 

> 참고 : URL 에 포함된 QueryString 은 표준 라이브러리의 http.URL.Query() 함수로 얻을 수 있습니다.


## Workshop
* 다음의 RESTful 서버를 Gorilla Mux를 이용하여 개발하세요.
```http request
GET /hello

```
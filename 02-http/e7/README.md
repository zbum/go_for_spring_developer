## 웹페이지
* Go 언어에서의 웹페이지 처리는 http.ResponseWriter에 HTML 문법의 응답과 Content-Type 헤더만 바꿔주면 웹페이지를 표시하는 handler 를 구현할 수 있습니다.

### 기본적인 웹페이지 표시
```go
func (h MemberHandler) IndexPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	query := r.URL.Query()
	fmt.Fprintf(w, "<html><body> <h1>hello!!</h1> <h2> name query : %s</h2></body></html>", query["name"])
}
```
* 위의 코드는 html 문자열을 소스코드에 포함시켜야 한다는 문제점이 있습니다.

## 템플릿 사용
* 다음 코드는 Go 템플릿을 이용한 handler 예제입니다.
```go
func (h MemberHandler) IndexPageWithTemplate(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	query := r.URL.Query()
	// 템플릿 파싱을 요청때마다 실행합니다.
	t, err := template.ParseFiles("index2.html")
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, query)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
}
```
* 위 예제에서 사용자 요청이 올때마다 템플릿 파일을 파싱하는 문제가 있습니다.

## 템플릿 캐시
* 파싱한 템플릿을 캐싱하기 위해 전역 변수로 처리하겠습니다.
```go
var fsTemplates = template.Must(template.ParseFiles("index2.html"))

func (h MemberHandler) IndexPageWithTemplateCache(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	query := r.URL.Query()

	err := fsTemplates.ExecuteTemplate(w, "index2.html", query)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
}
```
* 위의 예제에서 index2.html 템플릿 파일은 실행 경로에 존재해야 합니다. 따라서 배포할때 index2.html 파일을 따로 배포해야합니다.


## go embed
* 빌드한 패키지 내에 파일을 포함시키고 그 파일을 사용하기 위해서는 //go:embed 지시자를 사용합니다. 
* 이 지시자로 받을 수 있는 타입은 string, []byte, FS 입니다. 

### hello.txt 파일의 내용을 stirng 변수에 할당하기
```go
import _ "embed"

//go:embed hello.txt
var s string
print(s)
```
### hello.txt 파일의 내용을 byte 슬라이스 변수에 할당하기
```go
import _ "embed"

//go:embed hello.txt
var s []byte
print(string(s))
```

### 하나 이상의 파일을 파일 시스템(embed.FS)에 할당하기
```go
import "embed"

//go:embed hello.txt
var f embed.FS
data, _ := f.ReadFile("hello.txt")
print(string(data))
```

## 템플릿 임베드
* go 언어의 장점인 1개의 실행파일로 빌드할 수 있는 장점을 위해 template을 go embed를 이용해 처리하겠습니다.
* go:embed 지시자는 상위 디렉토리로 접근할 수 없기 때문에 embed.FS를 선언한 하위에 디렉토리를 생성하여 템플릿을 저장합니다.
### web_template.go
* 아래와 같이 지시자를 사용했다면 템플릿은 web_template.go 파일이 존재하는 디렉토리의 하위 html 디렉토리에 위치해야 합니다.
```go
package templates

import "embed"

//go:embed html
var TemplatesFS embed.FS
```

### 임베드 템플릿 읽기
* 임베드 템플릿은 html/index2.html 경로로 접근해야 합니다.
```go
var embedTemplates = template.Must(template.ParseFS(templates.TemplatesFS, "html/index2.html"))
```


### 
## 출처
* https://go.dev/doc/articles/wiki/#tmp_10
* https://pkg.go.dev/embed
* https://pkg.go.dev/text/template
* https://pkg.go.dev/html/template
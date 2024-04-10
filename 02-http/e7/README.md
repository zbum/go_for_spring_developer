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


## 템플릿 임베드
* go 언어의 장점인 1개의 실행파일로 빌드할 수 있는 장점을 위해 template을 go embed를 이용해 처리하겠습니다.


## 출처
* https://go.dev/doc/articles/wiki/#tmp_10
* https://pkg.go.dev/embed
* https://pkg.go.dev/text/template
* https://pkg.go.dev/html/template
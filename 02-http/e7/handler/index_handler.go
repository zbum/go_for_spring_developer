package handler

import (
	"fmt"
	"go_for_spring_developer/02-http/e7/templates"
	"html/template"
	"net/http"
)

var fsTemplates = template.Must(template.ParseFiles("index2.html"))
var embedTemplates = template.Must(template.ParseFS(templates.TemplatesFS, "html/index2.html"))

type MemberHandler struct {
}

func NewIndexHandler() *MemberHandler {
	return &MemberHandler{}
}

func (h MemberHandler) IndexPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	query := r.URL.Query()
	fmt.Fprintf(w, "<html><body> <h1>hello!!</h1> <h2> name query : %s</h2></body></html>", query["name"])
}

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

func (h MemberHandler) IndexPageWithTemplateCache(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	query := r.URL.Query()

	err := fsTemplates.ExecuteTemplate(w, "index2.html", query)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
}

func (h MemberHandler) IndexPageWithTemplateCacheAndEmbed(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	query := r.URL.Query()

	err := embedTemplates.ExecuteTemplate(w, "index2.html", query)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
}

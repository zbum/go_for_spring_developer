## 템플릿
* Go 언어는 Thymeleaf와 같은 템플릿 엔진을 표준라이브러리에서 제공합니다. 
* 템플릿은 정적인 텍스트와 내장된 커맨드들로 구성된다. 커맨드는 {{ … }} 형식의 구분자로 표기합니다.
* 템플릿 패키지는 text/template 와 html/template 가 제공되며 모두 동일한 인터페이스를 구현하고 있습니다. 
* html/template는 웹보안 문제를 강화한 버전입니다. (html escaping 이 주요기능)

## go template 사용예
### Simple
* Inventory 구조체의 값을 템플릿에서 바인딩하여 표준출력에 표시하는 내용입니다.
* {{ … }} 형식의 구분자에 구조체 필드를 .으로 시작하여 설정합니다.
* 외부 패키지가 읽어야 하기 때문에 구조체 필드는 대문자로 시작해야 합니다.
```go
type Student struct {
    Id         int64
    Name       string
    Department string
}

func simple() {
    template1 := `## Go for Spring Developer (GFSD) - [simple]
Hello {{.Name}}!!
`
    template1Tmpl, err := template.New("template1").Parse(template1)
    if err != nil {
        return
    }
    
    err = template1Tmpl.Execute(os.Stdout, Student{Name: "Manty"})
    if err != nil {
        return
    }
}
```

## loop
* 고 템플릿에서 반복문은 range 키워드를 이용해서 처리할 수 있습니다.
```gotemplate
{{range .Students}}
	Student Name is {{.Name}}
{{end}}
```

## condition
* 고 템플릿에서 조건 분기는 if 키워드를 이용해서 처리할 수 있습니다.
```gotemplate
{{if eq .Name "Manty"}}Good {{else}}Nice {{end}}Student Name is {{.Name}}
```

## 출처
* https://pkg.go.dev/text/template
* https://pkg.go.dev/html/template
package main

import (
	"fmt"
	"os"
	"text/template"
)

type Student struct {
	Id         int64
	Name       string
	Department string
}

func main() {
	simple()
	loop()
	condition()
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

func loop() {
	template1 := `## Go for Spring Developer (GFSD) - [loop]
{{range .Students}}
	Student Name is {{.Name}}
{{end}}
`
	template1Tmpl, err := template.New("template1").Parse(template1)
	if err != nil {
		fmt.Println(err)
		return
	}

	students := []Student{
		{Name: "Manty"},
		{Name: "Zbum"},
		{Name: "Dongmyo"},
		{Name: "Boorng"},
	}

	data := make(map[string]any)
	data["Students"] = students
	err = template1Tmpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func condition() {
	template1 := `## Go for Spring Developer (GFSD) - [condition]
{{range .Students}}
	{{if eq .Name "Manty"}}Good {{else}}Nice {{end}}Student Name is {{.Name}}
{{end}}
`
	template1Tmpl, err := template.New("template1").Parse(template1)
	if err != nil {
		fmt.Println(err)
		return
	}

	students := []Student{
		{Name: "Manty"},
		{Name: "Zbum"},
		{Name: "Dongmyo"},
		{Name: "Boorng"},
	}

	data := make(map[string]any)
	data["Students"] = students
	err = template1Tmpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println(err)
		return
	}
}

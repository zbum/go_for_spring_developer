## 스프링 부트 프로젝트를 Go 프로젝트로 포팅하기
* 제공하는 예제
  * https://github.com/zbum/gfsd_spring_boot_ex1

## Project Main
* 스프링 부트 프로젝트는 @SpringBootApplication 애너테이션을 설정하여 프로젝트 실행합니다. 
```java
package kr.co.manty.gfsd_spring_boot_ex1;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class GfsdSpringBootEx1Application {

    public static void main(String[] args) {
        SpringApplication.run(GfsdSpringBootEx1Application.class, args);
    }

}
```
* Go 에서는 main.go 를 생성하여 프로젝트를 실행합니다. main.go 는 웹서버를 실행하는 역할을 합니다.
```go
package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
```

## 스프링 부트의 Controller의 포팅
* Controller는 웹요청을 해석하고 Service 로 처리를 위임하고 응답을 담당합니다. 
```java
package kr.co.manty.gfsd_spring_boot_ex1.student;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/students")
public class StudentController {

    private final StudentService studentService;

    public StudentController(StudentService studentService) {
        this.studentService = studentService;
    }

    @GetMapping("/{student-id}")
    public Student retrieveStudent(@PathVariable("student-id") Long id) {
        return studentService.getStudent(id);
    }

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public Student registerStudent(@RequestBody Student student) {
        studentService.registerStudent(student);
        return student;
    }
}
```
* StudentController가 처리하는 "GET /students/{id}", "POST /students"를 Go Handler 로 다음과 같이 개발합니다. 
```go
package student

import "net/http"

type Student struct {
    Id    int64
    Name  string
    Score int
}

type StudentHandler struct {
}

func NewStudentHandler() *StudentHandler {
    return &StudentHandler{}
}

func (h *StudentHandler) GetStudent(w http.ResponseWriter, r *http.Request) {
}

func (h *StudentHandler) RegisterStudent(w http.ResponseWriter, r *http.Request) {
}
```
* 이 핸들러가 동작하려면 핸들러 객체를 생성하고 라우터에 등록합니다. 
```go
package main

import (
	"gfsd_go_mux_ex1/student"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func initStudentHandler() *student.StudentHandler {
	return student.NewStudentHandler()
}
func main() {
	studentHandler := initStudentHandler()

	r := mux.NewRouter()
	r.HandleFunc("/students/{id}", studentHandler.GetStudent).Methods("GET")
	r.HandleFunc("/students", studentHandler.RegisterStudent).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}
```
## StudentService 포팅
* Go로 작성한 StudentService는 다음과 같이 GetStudent, RegisterStudent 를 제공합니다. 지금은 특별한 기능을 제공하지 않습니다.
```go
package student

type StudentService struct {
}

func NewStudentService() *StudentService {
	return &StudentService{}
}

func (s *StudentService) GetStudent(id int64) *Student {
	return &Student{id, "Manty", 100}
}

func (s *StudentService) RegisterStudent(student *Student) {
}
```

* StudentHandler가 StudentService를 사용할 수 있도록 의존성을 주입하고 사용하도록 Struct 와 생성자 함수를 수정합니다.
```go
type StudentHandler struct {
	studentService *StudentService
}

func NewStudentHandler(studentService *StudentService) *StudentHandler {
	return &StudentHandler{studentService}
}
```

* main.go 에서 의존성 주입을 처리합니다.
```go
func initStudentHandler() *student.StudentHandler {
	studentService := student.NewStudentService()
	return student.NewStudentHandler(studentService)
}
```

* StudentHandler 의 메소드 중에 GetStudent를 다음과 같이 구현합니다.
```go
func (h *StudentHandler) GetStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	student := h.studentService.GetStudent(id)
	studentBytes, err := json.Marshal(student)
	w.Write(studentBytes)
}
```

* 이제 서버를 실행하고 curl 로 테스트 해보면 다음과 같이 응답합니다. 
```shell
$ curl localhost:8080/students/1
{"Id":1,"Name":"Manty","Score":100}%  
```

* 여기서 json 속성이 대문자로 시작합니다. Go에서는 외부 패키지에서 접근하기 위해서 Struct의 속성을 대문자로 선언해야 하는 제약이 있습니다.
* 이 속성을 소문자로 바꾸기 위해서 다음과 같이 주석을 추가합니다. 
```go
type Student struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Score int    `json:"score"`
}
```
* 응답을 확인해 보면 json 속성이 소문자로 보이는 것을 알 수 있습니다.
```shell
$ curl localhost:8080/students/1
{"id":1,"name":"Manty","score":100}%  
```

## StudentRepository 작성
* Ent 프레임워크를 사용합니다. 
* 터미털에서 다음을 실행하여 Student 스키마를 생성합니다. 
```shell
$ go run -mod=mod entgo.io/ent/cmd/ent new --target student/ent/schema Student
```
* 생성된 student.go 파일은 다음과 같습니다. 
```go
package schema

import "entgo.io/ent"

// Student holds the schema definition for the Student entity.
type Student struct {
	ent.Schema
}

// Fields of the Student.
func (Student) Fields() []ent.Field {
	return nil
}

// Edges of the Student.
func (Student) Edges() []ent.Edge {
	return nil
}
```
* 필요한 필드를 추가합니다. 여기서는 id, name, score 를 추가할 수 있습니다. 
```go
// Fields of the Student.
func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Positive(),
		field.String("name").
			Default("unknown"),
		field.Int("score"),
	}
}
```

* 코드 생성을 위해 다음의 파일을 student/ent 디렉토리에 생성합니다. 파일명은 generate.go 입니다.
```go
package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema
```

* 이제 다음을 실행하여 코드를 생성합니다.
```shell
$ go generate ./...
```

* 이제 StudentRepository를 작성할 준비가 다 되었습니다. 
* FindById 메소드의 내용을 다음과 같이 작성합니다.



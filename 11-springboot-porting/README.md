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

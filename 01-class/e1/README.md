## Class
* Go는 클래스(Class)가 없습니다.
* Struct가 Class의 역할을 수행 할 수 있기는 하지만 메서드도 구조체로부터 분리되는 구성을 가지고 있습니다.
* 단일 상속도 없고 당연히 다중 상속도 없습니다.
* 객체지향스럽지 않은 언어로 보일 수 있겠지만 충분히 객체지향적입니다.

<br />

## Class 와 메소드
* Java 코드로 작성한 간단한 StudentService 클래스는 다음과 같이 작성되어 있습니다.
* 클래스를 만들기 위해 class 키워드를 사용했고, Student 타입을 반환하는 getStudent 메서드와 Student를 등록하기 위해 registerStudent 를 선언하였습니다. 
* getStudent 와 registerStudent 는 클래스 및 패키지 외부에서도 접근 가능하도록 public 접근제어자를 사용했습니다.
```java

@Service
public class StudentService {

    private final StudentRepository studentRepository;

    public StudentService(StudentRepository studentRepository) {
        this.studentRepository = studentRepository;
    }

    public Student getStudent(Long id) {
        return studentRepository.findById(id)
                .orElseThrow(() -> new StudentNotFoundException(id));
    }

    @Transactional
    public void registerStudent(Student student) {
        studentRepository.save(student);
    }
}

```
* 이제 이 클래스를 Go 언어로 작성해 보겠습니다.
```go
package main

import "fmt"

type StudentService struct {
    studentRepository StudentRepository
}

func NewStudentService(studentRepository StudentRepository) *StudentService {
    return &StudentService{studentRepository}
}

func (r StudentService) GetStudent(id int64) *Student {
    return r.studentRepository.FindById(id)
}

func (r StudentService) RegisterStudent(student Student) error {
    if r.studentRepository.FindById(student.id) != nil {
        return fmt.Errorf("user Already Exists: %d", student.id)
    }
    r.studentRepository.Save(student)
    return nil
}
```
* Go언어에서는 Struct로 클래스를 대체합니다. 
* 타입에는 함수를 추가할 수 있는데 이것을 타입 메소드라고 합니다. 
* 메소드는 함수에 리시버를 이용해서 작성합니다.
* NewStudentService 라는 Struct에 GetStudent, RegisterStudent 메소드를 추가합니다.
* GetStudent, RegisterStudent 메소드는 value 리시버를 사용하고 있습니다.
* NewStudentService 메소드는 외부 패키지에서도 접근할 수 있도록 GetStudent, RegisterStudent와 같이 대문자로 시작하는 함수명을 가지고 있습니다. 이것은 Go 에서 Export 하였다고 말합니다.


## Interface
* 인터페이스 타입은 구체적인 동작을 구현할 메서드의 집합입니다.
* 이 인터페이스를 구현하려면 인터페이스에서 정의한 모든 메서드를 구현해야합니다. (Java의 implements 가 필요하지 않습니다.)
* 함수의 매개변수 타입이 인터페이스라면 이 인터페이스를 구현한 변수를 전달할 수 있습니다.

### Interface 의 예
* io.Reader
```go
type Reader interface {
    Read(p []byte) (n int, error)
}
```
* io.Writer
```go
type Writer interface {
    Write(p []byte) (n int, error)
}
```
* io.Reader 를 구현하는 예 (수정 및 기능 확인 필요)
```go
type Reader struct {
    data []byte
    readIndex int64
}

func (r *Reader) Read(p []byte) (n int, err error) {
    if r.readIndex >= int64(len(r.data)) {
        err = io.EOF
        return
    }

    n = copy(p, r.data[r.readIndex:])
    r.readIndex += int64(n)
    return
}
```
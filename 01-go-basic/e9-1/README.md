### 포인터 반환
* 다음의 코드에서 Init() 함수가 반환하는 Student 타입의 데이터는 반환시 main() 함수가 실행되는 스택으로 복사됩니다.
```go
package main

import "fmt"

type Student struct {
}

func Init() Student {
	var student = Student{}
	return student
}

func main() {
	student := Init()
	fmt.Println(student)
}
```
* 하지만 Init() 함수가 Student 타입의 포인터를 반환하면 상황은 달라집니다.
* 포인터를 반환하는 순간 Init() 함수의 스택이 사라지므로 student 변수의 데이터를 유지하기 위해 Go컴파일러는 student 데이터를 힙으로 이동시킵니다.
```go
package main

import "fmt"

type Student struct {
}

func Init() *Student {
	var student = Student{}
	return &student
}

func main() {
	student := Init()
	fmt.Println(student)
}

```

* 스택으로 전달되는 과정을 확인해 보려면 다음의 명령으로 실행해 보면 알 수 있습니다.
```go
$ go build -gcflags '-m'
# go_for_spring_developer/00-1-go-basic/e9-1
./pointer_return.go:8:6: can inline Init
./pointer_return.go:14:17: inlining call to Init
./pointer_return.go:15:13: inlining call to fmt.Println
./pointer_return.go:9:6: moved to heap: student
./pointer_return.go:14:17: moved to heap: student
./pointer_return.go:15:13: ... argument does not escape
```
* 스택의 데이터를 힙으로 이동하는 부분을 위의 내용으로 확인 할 수 있습니다.
* 힙에 전달된 메모리는 가비지 컬렉션 대상이 되고, 성능에 영향을 주기 때문에 필요한 경우가 아니면 포인터를 사용하지 않는 것이 좋습니다.

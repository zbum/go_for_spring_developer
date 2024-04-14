
## type - 구조체
### 구조체 정의
* 여러 값의 집합을 하나의 데이터 타입으로 묶어 하나의 객체처럼 주고 받는 방법입니다.
* 구조체는 새로운 데이터 타입이므로 type 키워드를 사용하고 구조체의 이름, struct 키워드를 사용해야 합니다.
```go
type Student struct {
    Id   int64
    Name string
}
```
### 구조체 타입 변수 초기화
* new() 키워드
    * 적절한 메모리 공간을 할당하고 제로 값으로 만든다.
    * 할당된 메모리의 포인터를 반환한다.

* 구조체는 구조체 이름과 중괄호를 이용해서 초기화 할 수 있습니다.
    * 초기값을 입력하지 않으면 컴파일러가 필드의 타입별 제로 값을 할당합니다.
#### 구조체 타입 변수(E8)
* new 키워드로 구조체 초기화
```go
	// Student 구조체 포인터 생성
	mantyPtr := new(Student)
	mantyPtr.Id = 1
	mantyPtr.Name = "manty"
	fmt.Println(mantyPtr)
```
* zero 값으로 구조체 초기화
```go
	zero := Student{}
	fmt.Println(zero)
```
* 필드값을 가진 구조체 초기화
```go
	comtin := Student{
		Id:   2,
		Name: "comtin",
	}
	fmt.Println(comtin)
```
* 구조체 포인터 생성
```go
	comtinPtr := &Student{
		Id:   2,
		Name: "comtin",
	}
	fmt.Println(comtinPtr)
```
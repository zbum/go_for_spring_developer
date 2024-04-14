## encoding/json 패키지
## Marshal
* Go 데이터를 json 으로 직렬화하기 위해서는 Marshal() 함수를 사용합니다.
* 구조체 혹은 map 데이타를 JSON으로 인코딩하게 되는데, 해당 Go 데이타 값을 json.Marshal()의 파라미터로 전달하면, JSON으로 인코딩된 바이트배열과 에러객체를 리턴합니다.
* 만약 JSON으로 인코딩된 []byte을 다시 문자열로 변경할 필요가 있다면, string([]byte)과 같이 변경할 수 있다.
* 한가지 유의할 점은 JSON의 Key는 문자열이어야 한다. Go 구조체의 경우 자동으로 필드명을 문자열로 사용하게 되지만, map인 경우는 map[string]T 처럼 Key가 string인 map만 지원합니다.

```go
func marshal() {
	response := StudentResponse{Id: 10, Name: "Manty"}
	marshal, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	fmt.Println("[1]marshal:", string(marshal))
}
```
* 출력 결과
```
[1]marshal: {"Id":10,"Name":"Manty"}
```
## Unmarshal
* JSON으로 인코딩된 데이타를 다시 디코딩하기 위해서는 encoding/json 패키지의 Unmarshal() 함수를 사용합니다.
* Unmarshal() 함수의 첫번째 파라미터에는 JSON 데이타를, 두번째 파라미터에는 출력할 구조체(혹은 map)를 포인터로 지정합니다. 
* 리턴값은 에러객체이고, 에러가 없을 경우, 두번째 파라미터에 Json 데이터가 설정 됩니다.
```go
	requestString := `{"Id":10,"Name":"Manty"}`
	var request StudentRequest
	err := json.Unmarshal([]byte(requestString), &request)
	if err != nil {
		panic(err)
	}
	fmt.Println("[2] unmarshal : ", request)
```

## JSON Tag
* 위 예제에서 StudentResponse 구조체에서 Id, Name 은 외부에서 필드에 접근할 수 있도록 대문자로 작성되어 있습니다. 
* 결과적으로 marshal 결과도 대문자로 표시되는 문제가 발생합니다. 
* 이와 같이 JSON 처리에서 구체적인 지시를 하기 위해 Tag 를 구조체에 적용합니다. 
* 다음 예제는 json.Marshal, json.Unmarshal 처리시에 Id 대신 id, Name 대신 name 으로 처리하도록 tag 를 처리한 모습입니다.
```go
type StudentResponse struct {
    Id   *int64 `json:"id"`
    Name string `json:"name"`
}
```

## JSON Tag - omitempty
* StudentResponse 구조체  초기화에서 Name 값만 설정하였습니다.
```go

type StudentResponse struct {
    Id   *int64 `json:"id"`
    Name string `json:"name"`
}

func marshal() {
    response := StudentResponse{Name: "Manty"}
    marshal, err := json.Marshal(response)
    if err != nil {
        panic(err)
    }
    fmt.Println("[1]marshal:", string(marshal))
}
```
* marshal() 함수의 결과는 다음과 같습니다.
```
[1]marshal: {"id":0,"name":"Manty"}
```
* Go 언어에서 포인터 변수가 아닌 경우는 항상 초기 값을 가지고 있습니다. 그래서 id:0 이 표시됩니다. 
* 구조체에 omitempty json tag 를 추가해서 결과에서 제외 할 수 있습니다. 

```go
type StudentResponse struct {
    Id   *int64 `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
}
```
* omitempty 가 동작하기 위해서는 nil 값을 담을 수 있는 포인터 타입의 필드이어야 한다는 점을 주의하시기 바랍니다.






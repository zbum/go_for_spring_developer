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



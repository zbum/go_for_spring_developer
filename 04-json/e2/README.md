## JSON Tag
* 위 예제에서 StudentResponse 구조체에서 Id, Name 은 외부에서 필드에 접근할 수 있도록 대문자로 작성되어 있습니다.
* 결과적으로 marshal 결과도 대문자로 표시되는 문제가 발생합니다.
* 이와 같이 JSON 처리에서 구체적인 지시를 하기 위해 Tag 를 구조체에 적용합니다.
* 다음 예제는 json.Marshal, json.Unmarshal 처리시에 Id 대신 id, Name 대신 name 으로 처리하도록 tag 를 처리한 모습입니다.
```go
type StudentResponse struct {
    Id   int64  `json:"id"`
    Name string `json:"name"`
}
```
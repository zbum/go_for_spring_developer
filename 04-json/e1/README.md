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
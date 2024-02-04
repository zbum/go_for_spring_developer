## Context
* Context 패키지는 데드라인, 취소 시그널이나 요청 범위(request scoped) 의 추가적인 값을 전달하기 위해서 사용합니다.
* 서버로 들어오는 요청은 컨텍스트를 생성해야 하며, 서버로 나가는 호출은 컨텍스트를 수락해야 합니다. 
* 이들 사이의 함수 호출 체인은 컨텍스트를 전파해야 하며 
* 선택적으로 WithCancel, WithDeadline, WithTimeout 또는 WithValue를 사용하여 생성된 파생 컨텍스트로 대체해야 합니다. 
* 컨텍스트가 취소되면 해당 컨텍스트에서 파생된 모든 컨텍스트도 취소됩니다.


## Context 를 이용한 값 전달

* 다음의 Go main 함수는 컨텍스트를 생성하고 값을 추가한 후, 최종적으로 조회하는 예제 입니다.  
```go
func main() {
	ctx := context.Background() //<-- Context 생성
	ctx = context.WithValue(ctx, "one", 1) //<-- 위에서 생성한 컨텍스트를 Parent 로 해서 one = 1 값을 가진 Child 컨텍스트를 생성한다.
	ctx = context.WithValue(ctx, "two", 2) //<-- 위에서 생성한 컨텍스트를 Parent 로 해서 two = 2 값을 가진 Child 컨텍스트를 생성한다.

	fmt.Println(ctx.Value("one"))
	fmt.Println(ctx.Value("two"))
	fmt.Println(ctx.Value("three"))
}
```
* 컨텍스트는 수정할 수 없기 때문에 부모 컨텍스트를 포함한 자식 컨텍스트를 생성하는 방식으로 구현합니다. 



## Context
* Context는 작업의 "맥락"을 나타내는 개념
* 데드라인, 취소 시그널이나 요청 범위(request scoped)의 추가적인 값을 전달하기 위한 인터페이스

## Context의 특징
* 컨텍스는 immutable tree structure 로 구현되어 있어서 고루틴에 안전합니다.
* 컨텍스트가 취소되면 해당 컨텍스트에서 파생된 모든 컨텍스트도 취소됩니다.

## Context 인터페이스 
* Context 인터페이스는 다음과 같은 메서드를 가지고 있습니다.
  * Done() <-chan struct{}  : 취소 시그널을 전달하는 메서드
  * Err() error : 작업이 왜 끝났는지 알려주는 메서드
  * Deadline() (time.Time, bool) : 작업의 마감 시한을 알려주는 메서드
  * Value(key interface{}) interface{} : 컨텍스트에 저장된 값을 조회하는 메서드

## Context 생성
* Context는 context 패키지에서 제공하는 여러 함수로 생성할 수 있습니다.
  * Background() : 최상위 컨텍스트를 생성합니다. 
  * TODO() : 작업을 나중에 구현할 때 사용합니다. 
  * WithCancel(parent Context) : 부모 컨텍스트를 취소할 수 있는 자식 컨텍스트를 생성합니다.
  * WithDeadline(parent Context, d time.Time) : 부모 컨텍스트의 마감 시한을 설정한 자식 컨텍스트를 생성합니다.
  * WithTimeout(parent Context, timeout time.Duration) : 부모 컨텍스트의 마감 시한을 설정한 자식 컨텍스트를 생성합니다.
  * WithValue(parent Context, key interface{}, val interface{}) : 부모 컨텍스트에 값을 추가한 자식 컨텍스트를 생성합니다.


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



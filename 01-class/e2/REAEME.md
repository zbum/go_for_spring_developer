## Interface
* 인터페이스 타입은 구체적인 동작을 구현할 메서드의 집합입니다.
* 이 인터페이스를 구현하려면 인터페이스에서 정의한 모든 타입 메서드를 구현해야합니다.
* Java의 implements 키워드가 필요하지 않고, 인터페이스가 요구하는 타입 메서드를 모두 구현하면 인터페이스를 만족하게 됩니다.
    * 이런 방식을 덕타이핑(Duck Typing)이라고 합니다.
> 만약 어떤 새가 오리처럼 걷고, 헤엄치고, 꽥꽥거리는 소리를 낸다면 나는 그 새를 오리라고 부를 것이다.

* 빈 인터페이스(interface{})는 아무 타입 메소드를 가지고 있지 않기 때문에 모든 타입에서 구현했다고 볼 수 있습니다.

## 인터페이스를 파라미터로 사용하는 함수
* 슬라이스를 정렬하는 예를 보겠습니다.
* Go 언어에서는 func sort.Sort(data sort.Interface) 라는 정렬을 위한 함수를 제공합니다. 
* 매개변수의 타입은 sort.Interface 이며 이 인터페이스는 다음과 같이 정의 되어 있습니다.
```go
type Interface interface {
	// 컬랙션의 원소 개수를 반환해야 합니다.
	Len() int

	// i 번째 인덱스의 원소가 j 번째 인덱스의 원소보다 먼저 위치해야 하면 true를 반환합니다.
	Less(i, j int) bool

	// i 번째 인덱스의 원소와 j 번째 인덱스의 원소를 교환합니다.
	Swap(i, j int)
}
```
* 즉, Len(), Less(i, j int) bool, Swap(i, j int) 를 타입 메소드롤 선언한 타입이라면 해당 인터페이스를 만족합니다.
* 그러면 일반적인 슬라이스를 정렬하도록 슬라이스에 위의 타입 메소드를 구현하겠습니다.

```go
// 타입이 구조체가 아니라도 상관 없습니다.
type SortableSlice []int

func (s SortableSlice) Len() int {
	return len(s)
}

func (s SortableSlice) Less(i, j int) bool {
	return i > j
}

func (s SortableSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
```
## 실습 (workshop01)
* 다음의 타입을 Id로 정렬하도록 Student 구조체에 타입 메소드를 추가해 주세요.


## type assertion
* 인터페이스의 값을 특정 타입처럼 쓸 수 있게 해주는 기능을 말합니다. 
* Java의 타입 캐스팅과 비슷한 기능입니다.

<br />
* interface{} 를 반환하는 다음의 함수가 있습니다. 이 함수는 실제 int 타입의 12를 반환합니다. 

```go
func returnNumber() interface{} {
	return 12
}
```

* 이 함수를 호출하여 받은 결괏값은 interface{} 타입이라 필드, 메소드가 정의 되어 있지 않습니다. 
```go
	anInt := returnNumber()
    anInt++ // 컴파일 에러!!!
	fmt.Println(anInt)
}
```
* 인터페이스로 받은 값을 실제 타입으로 바꾸기 위해 type assertion 을 다음과 같이 수행합니다.
```go
	anInt := returnNumber()
	number := anInt.(int) // type assertion
	number++
	fmt.Println(number)
```

* 위 예제에서 type assertion은 런타임에 해당 변수가 변환 가능한 것인지 아닌지 확인하지 않았으므로 위험을 내포하고 있습니다. 
* 만약 type assertion 이 실패 한다면 Panic 을 일으키기 때문에 성공 실패 여부를 확인할 필요가 있습니다. 
```go
	anInt := returnNumber()
	number, ok := anInt.(int)
	if ok {
		number++
		fmt.Println(number)
	} else {
		fmt.Println("Type assertion Failed")
	}
```

## type switch
* 인터페이스의 타입을 모르거나 경우의 수가 많을 때는 타입 스위치를 사용합니다. 
```go
func switchType(x interface{}) {

	switch T := x.(type) {
	case Unknown1:
		fmt.Println("Unknown type")
	case Unknown2:
		fmt.Println("Entry type")
	default:
		fmt.Printf("Not supported type: %T\n", T)
	}
}
```

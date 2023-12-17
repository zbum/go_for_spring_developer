## Class
* Go는 클래스(Class)가 없다 
* Struct가 Class의 역할을 수행 할 수 있기는 하지만 메서드도 구조체로부터 분리되는 구성을 가지고 있다. 
* 단일 상속도 없고 당연히 다중 상속도 없다. 
* 객체지향스럽지 않은 언어로 보일 수 있겠지만 충분히 객체지향적이다. 
* 그냥 좀 다른 방법으로 객체를 지향하고 있을 따름이다.

<br />

* 간단한 클래스를 만들어 봅시다.
```go
package main

import "fmt"

type Greeter struct {
}

func (g *Greeter) Greet(name string) string {
	return fmt.Sprintf("Hello, %s!!", name)
}
```
* Go언어에서는 Struct로 클래스를 대체합니다. 
* 메소드는 리시버를 이용해서 작성합니다.
* Greeter 이라는 Struct에 Greet 메소드를 추가합니다.


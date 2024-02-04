## Class
* Go는 클래스(Class)가 없습니다.
* Struct가 Class의 역할을 수행 할 수 있기는 하지만 메서드도 구조체로부터 분리되는 구성을 가지고 있습니다.
* 단일 상속도 없고 당연히 다중 상속도 없습니다.
* 객체지향스럽지 않은 언어로 보일 수 있겠지만 충분히 객체지향적입니다.

<br />

## Class 와 메소드
* Java 코드로 작성한 간단한 Greeter 클래스가 다음과 같이 작성되어 있습니다.
* 클래스를 만들기 위해 class 키워드를 사용했고, String 타입을 반환하는 greet 메서드와 message 변수를 변경하는 changeMessage 를 선언하였습니다. 
* greet 와 message 는 클래스 및 패키지 외부에서도 접근 가능하도록 public 접근제어자를 사용했습니다.
```java
public class Greeter {
    private String message;
    
    public String greet(String name) {
        return this.message + "," + name + "!!";
    }
    
    public void changeMessage(String message) {
        this.message = message;
    }
}
```
* 이제 이 클래스를 Go 언어로 작성해 보겠습니다.
```go
package main

import "fmt"

type Greeter struct {
	message string
}

func (g Greeter) Greet(name string) string {
	return fmt.Sprintf("%s, %s!!", g.message, name)
}

func (g *Greeter) ChangeMessage(message string) {
	g.message = message
}
```
* Go언어에서는 Struct로 클래스를 대체합니다. 
* 메소드는 함수에 리시버를 이용해서 작성합니다.
* Greeter 이라는 Struct에 Greet 메소드를 추가합니다.
* Greet 메소드는 value 리시버를 사용하고 있고 ChangeMessage 메소드는 포인터 리시버를 사용하고 있습니다.
* Greeter 클래스의 필드를 조회할 때는 value 리시버를 사용하고, Greeter의 필드의 값을 수정할때는 포인터 리시버를 사용합니다.
* Greeter의 메소드는 외부에서도 접근할 수 있도록 Greet, ChangeMessage 와 같이 대문자로 시작하는 함수명을 가지고 있습니다. 이것은 Go 에서 Export 하였다고 말합니다.




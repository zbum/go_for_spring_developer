## Go 언어의 특징 (What makes Go great?)
* Code is easy to read.
  * It's none magical language
* Nothing is hidden.
  * It's easy to figure out what functions invoke by what function in what order.
* Native binary
  * VM 이나 공유 라이브러리가 설치되었는지를 걱정할 필요가 없음

## Bean
* 저런 특징 때문에 프레임워크가 자동으로 처리하는 부분이 극히 없음
* 따라서 Bean 을 작성하려면 직접 객체를 생성하고 의존성을 부여
* bean scope 속성도 개발자가 직접 작성
<br />

* 간단한 사용자 정보를 제공하는 빈을 다음과 같이 만들어 봅시다.


```go
package bean

type UserRepository struct {
}

func (u *UserRepository) FindAll() []string {
	return []string{
		"Manty",
		"Benjamin",
	}
}
```


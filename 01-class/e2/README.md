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

func NewUserRepository() *UserRepository {
	fmt.Println("init UserRepository")
	return new(UserRepository)
}

```

* 스트럭트를 선언합니다. 
* FindAll() 이라는 메소드를 선언합니다. 
* 보통 생성자 대신 NewUserRepository 형식의 팩토리 메소드를 작성하여 생성합니다.

## Dependency Injection
* UserRepository에 의존하는 UserService 객체를 작성해 보겠습니다. 
```go
package bean

import "fmt"

type UserService struct {
	UserRepository *UserRepository
}

func (s *UserService) GetUsers() []string {
	return s.UserRepository.FindAll()
}

func NewUserService(userRepository *UserRepository) *UserService  {
	fmt.Println("init UserService")
	return &UserService{userRepository}
}
```

* UserService 스트럭트는 `UserRepository *UserRepository` 필드를 가지고 있습니다. 
* 팩토리 메소드는 의존성 주입을 받기 위한 `*UserRepository` 를 인자로 가지고 있고 UserService 초기화 할때, `*UserRepository` 를 사용합니다.



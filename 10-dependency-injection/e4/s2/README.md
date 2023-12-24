# [실습] Fx Simple Dependency Injection

## UserService에 의존하는 UserHandler
* UserRepository, UserService 라는 스트럭트가 제공됩니다.
* UserService 는 UserRepository 에 종속성을 가지고 있습니다. 즉, UserRepository 가 적절히 제공될 때만 정상적으로 기능을 제공할 수 있습니다.

```go
type UserRepository struct{}

func (UserRepository) GetUsers() []User {
    return []User{
        {Id: 1, Name: "Manty", Age: 40},
        {Id: 2, Name: "Manty1", Age: 41},
        {Id: 3, Name: "Manty2", Age: 42},
    }
}

type UserService struct {
    userRepository *UserRepository
}

func (s UserService) GetUsers() []User {
    return s.userRepository.GetUsers()
}
```

* Go 언어에서 관례적으로 생성자 함수를 사용합니다. 보통은 스트럭트 이름 앞에 New 접두사를 사용합니다. 
* NewUserService 함수에서는 *UserRepository 를 인자로 받아서 UserService 객체를 생성합니다.
```go
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func NewUserService(repository *UserRepository) *UserService {
	return &UserService{
        userRepository: repository,
    }
}
```

* UserService의 GetUsers함수를 실행하는 Start 함수는 다음과 같습니다.
```go
func Start(s *UserService) {
	fmt.Println(s.GetUsers())
}
```

* main 함수에서 UserRepository 와 UserService 객체를 생성하고 Start 함수를 다음과 같이 실행합니다.
```go
func main() {
	r := NewUserRepository()
	s := NewUserService(r)

	Start(s)
}
```

## Fx 사용
* main 함수를 다음과 같이 수정해 보겠습니다. 
```go
func main() {
	fx.
		New(
			fx.Provide(NewUserService),
			fx.Provide(NewUserRepository),
			fx.Invoke(Start),
		).
		Run()
}
```
* `fx.New(Options...).Run()` 으로 Fx 애플리케이션을 실행합니다. 
* fx.Provide 로 생성 및 의존성 주입을 처리할 컴포넌트 생성함수를 제공합니다. 생성함수의 인자에 의해서 자동으로 의존성이 주입됩니다. 
* 모든 컴포넌트의 생성과 의존성 주입이 끝나면 fx.Invoke 로 제공한 함수를 실행합니다. 이때도 의존성 주입이 도움을 줍니다.

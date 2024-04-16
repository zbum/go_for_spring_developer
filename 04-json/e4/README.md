## 입력값 검증(JSR303)
* 스프링 부트를 사용한다면 spring-boot-starter-validation 라이브러리 의존성을 적용하여 @Validated, @Valid 에너테이션으로 입력갑 검증을 할 수 있습니다. 
* Go 에서는 표준 라이브러리에서 제공하는 입력값 검증은 없으나 검증용 메소드를 만들거나 3rd 라이브러리를 사용해 볼 수 있습니다.

## 라이브러리 없이 구현하기
* 다음과 같이 IsValid() 메소드로 구현할 수 있습니다. 
* 하지만 필드가 많아지거나 내용이 복잡해 지면 코드 가독성이 매우 떨어지는 단점이 있습니다.
```go
type SimpleRequest struct {
	Age  uint
	Name string
}

func (sr SimpleRequest) IsValid() bool {
	if sr.Age < 18 {
		return false
	}

	if len(sr.Name) == 0 {
		return false
	}
	return true
}
```

## 3rd Party 라이브러리
* 다음의 라이브러라가 가장 활발한 활동을 하는 라이브러리입니다.
* https://github.com/go-playground/validator
* 입력값 검증을 위한 다양한 구조테그를 제공합니다. 대표적인 것은 다음과 같습니다.
### Field
| Tag     | Description                         | Usage                   |
|---------|-------------------------------------|-------------------------|
| eqfield | 다른 필드와 동일 값을 가짐                     | eqfield=ConfirmPassword |
| cidr    | Classless Inter-Domain Routing CIDR | cidr                    |
| gt      | 더 커야 함                              | gt=10                 |






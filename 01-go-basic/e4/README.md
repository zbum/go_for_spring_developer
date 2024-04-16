
2. switch 문에는 표현식이 없고 case 문에 평가 표현식이 있는 경우
#### (E4)
```go

    value, err := strconv.Atoi(argument)
    if err != nil {
    fmt.Println("Can not convert to int:", argument)
    return
    }

	switch {
	case value == 0:
		fmt.Println("영")
	case value > 0:
		fmt.Println("양의 정수")
	case value < 0:
		fmt.Println("음의 정수")
	default:
		fmt.Println("이 조건에 올 수 없습니다.", value)
	}
```
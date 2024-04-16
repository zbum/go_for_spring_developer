package main

import (
	"errors"
	"fmt"
)

// MemberNotFoundError custom error (implements Error)
type MemberNotFoundError struct {
	MemberId   int32
	MemberName string
}

func (e *MemberNotFoundError) Error() string {
	return fmt.Sprintf("Member[%d] is not Exist.", e.MemberId)
}

func main() {
	errors.New("aaa")
	_, err := FindMemberById(10)
	processError(err)

	_, err = FindMemberByName("Zbum")
	processError(err)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered!!", r)
		}
	}()
	panic("I called a panic()!!")
	fmt.Printf("After panic()")

}

func processError(err error) {
	var t *MemberNotFoundError
	switch {
	case errors.As(err, &t):
		fmt.Println("It must be 404 status code.", t.MemberId, t.MemberName)
	default:
		fmt.Println("It must be 500 status code.", t)
	}
}

// FindMemberById 직접 에러를 발생시킵니다.
func FindMemberById(id int32) (string, error) {
	return "", &MemberNotFoundError{MemberId: id}
}

// FindMemberByName wrapping error 예제
func FindMemberByName(name string) (string, error) {
	return "", fmt.Errorf("error %w", &MemberNotFoundError{MemberName: name})
}

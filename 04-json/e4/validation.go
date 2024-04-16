package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

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

type ValidatorRequest struct {
	Age             uint   `validate:"gt=18"`
	Name            string `validate:"required"`
	Password        string `validate:"eqfield=ConfirmPassword"`
	ConfirmPassword string
}

func main() {

	// (1) Simple
	request := SimpleRequest{17, "Kid"}
	fmt.Printf("(1) validation result : %t \n", request.IsValid())

	// (2) validator 사용
	validate := validator.New(validator.WithRequiredStructEnabled())

	validatorRequest := ValidatorRequest{Age: 11, Password: "1234", ConfirmPassword: "1235"}
	fmt.Printf("(2) validation result : %v \n", validate.Struct(validatorRequest))
}

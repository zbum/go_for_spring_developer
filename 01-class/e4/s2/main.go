package main

import (
	"fmt"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Invoke(PrintHello),
		fx.Provide(NewUserService),
	).
		Run()
}

func PrintHello() {
	fmt.Println("Hello")
}

type UserService struct {
}

func NewUserService() *UserService {
	fmt.Println("creating NewUserService.")
	return &UserService{}
}

package main

import (
	"fmt"
	"go.uber.org/fx"
)

func main() {
	fx.
		New(
			fx.Invoke(StartHandler),
			fx.Provide(NewUserService),
			fx.Provide(NewUserHandler),
		).
		Run()
}

func StartHandler(handler *UserHandler) {
	handler.GetUser()
}

type UserService struct {
}

func NewUserService() *UserService {
	fmt.Println("creating NewUserService.")
	return &UserService{}
}

func (*UserService) GetUser() {
	fmt.Println("GetUser of UserService called.")
}

type UserHandler struct {
	UserService *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) GetUser() {
	fmt.Println("GetUser of UserHandler called.")
	h.UserService.GetUser()
}

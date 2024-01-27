package main

import (
	"go_for_spring_developer/01-class/e2/bean"
	"log"
)

func main() {
	userRepository := bean.NewUserRepository()
	userService := bean.NewUserService(userRepository)
	log.Printf("%v \n", userService.GetUsers())
}

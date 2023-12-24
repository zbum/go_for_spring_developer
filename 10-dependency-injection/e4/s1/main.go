package main

import (
	"fmt"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}
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

func NewUserService(repository *UserRepository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

func (s UserService) GetUsers() []User {
	return s.userRepository.GetUsers()
}

func Start(s *UserService) {
	fmt.Println(s.GetUsers())
}

func main() {
	r := NewUserRepository()
	s := NewUserService(r)

	Start(s)
}

//func main() {
//	fx.
//		New(
//			fx.Provide(NewUserService),
//			fx.Provide(NewUserRepository),
//			fx.Invoke(Start),
//		).
//		Run()
//}

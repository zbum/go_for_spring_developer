package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/fx"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}
func (r UserRepository) GetUsers() []User {
	return []User{
		{Id: 1, Name: "Manty1", Age: 10},
		{Id: 2, Name: "Manty1", Age: 20},
		{Id: 3, Name: "Manty1", Age: 30},
		{Id: 4, Name: "Manty1", Age: 40},
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
	fx.
		New(
			fx.Provide(NewUserService),
			fx.Provide(NewUserRepository),
			fx.Invoke(Start),
		).
		Run()
}

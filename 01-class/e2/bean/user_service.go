package bean

import "fmt"

type UserService struct {
	UserRepository *UserRepository
}

func (s *UserService) GetUsers() []string {
	return s.UserRepository.FindAll()
}

func NewUserService(userRepository *UserRepository) *UserService  {
	fmt.Println("init UserService")
	return &UserService{userRepository}
}

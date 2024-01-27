package bean

import "fmt"

type UserRepository struct {
}

func (u *UserRepository) FindAll() []string {
	return []string{
		"Manty",
		"Benjamin",
	}
}

func NewUserRepository() *UserRepository {
	fmt.Println("init UserRepository")
	return new(UserRepository)
}

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

var UserRepositoryBean *UserRepository

func init() {
	UserRepositoryBean = &UserRepository{}
	fmt.Println(UserServiceBean)
}

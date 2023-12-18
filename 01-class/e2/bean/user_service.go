package bean

type UserService struct {
	UserRepository *UserRepository
}

func (s *UserService) GetUsers() []string {
	return s.UserRepository.FindAll()
}

var UserServiceBean *UserService

func init() {
	UserServiceBean = &UserService{UserRepositoryBean}
}

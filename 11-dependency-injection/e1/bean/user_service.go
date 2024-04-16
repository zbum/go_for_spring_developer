package bean

type UserService struct {
	UserRepository *UserRepository
}

func (s *UserService) GetUsers() []string {
	return s.UserRepository.FindAll()
}

func NewUserService(repository *UserRepository) *UserService {
	return &UserService{repository}
}

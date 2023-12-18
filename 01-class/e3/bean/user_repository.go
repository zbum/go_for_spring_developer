package bean

type UserRepository struct {
}

func (u *UserRepository) FindAll() []string {
	return []string{
		"Manty",
		"Benjamin",
	}
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

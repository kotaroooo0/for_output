package mock_callcount

type UserService interface {
	FindUser(string) (User, error)
}

type UserServiceImpl struct {
	UserRepository UserRepository
}

func (s UserServiceImpl) FindUserById(id string) (User, error) {
	u, err := s.UserRepository.FindUser(id)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

type UserRepository interface {
	FindUser(string) (User, error)
}

type User struct {
	Id   string
	Name string
}

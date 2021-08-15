package usecase

type UsersRepository interface {
	Get(userID int) (domain.User, error)
}

package mock_callcount

import "github.com/go-redis/redis/v7"

type UserRepositoryImpl struct {
	Client *redis.Client
}

func (h UserRepositoryImpl) FindUser(id string) (User, error) {
	// なんか
	return User{Id: id}, nil
}

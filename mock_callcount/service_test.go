package mock_callcount

import (
	"testing"

	"github.com/go-redis/redis/v7"
)

type UserRepositoryMock struct {
	Client            *redis.Client
	FindUserCallCount int
}

func (h UserRepositoryMock) FindUser(id string) (User, error) {
	h.FindUserCallCount++
	// なんか
	return User{Id: id}, nil
}

func TestFindUserById(t *testing.T) {

	userRepositoryMock := UserRepositoryMock{Client: testClient}
	userService := UserService{
		UserRepository: userRepositoryMock,
	}

	userService.FindUserById()
	userService.FindUserById()
	// 呼び出し回数を取得できる
	callCount := userService.UserRepository.FindUserCallCount
}

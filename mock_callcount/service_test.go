package mock_callcount

import (
	"testing"

	"github.com/go-redis/redis/v7"
)

type UserRepositoryMock struct {
	Client *redis.Client
}

func (h UserRepositoryMock) FindUser(id string) (User, error) {
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
	// ここでUserRepositoryのFindUserメソッドは何度呼ばれたかをテストしたい
	// ここではServiceのFindByIdとRepositoryのFindUserは1:1なので明らかに2回
	userService.FindUserById()
	userService.FindUserById()
}

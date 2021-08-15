package firebase

import (
	"firebase.google.com/go/messaging"
)

type PushRepositoryImpl struct {
	client *messaging.Client
}

func NewPushRepository(client *messaging.Client) *PushRepositoryImpl {
	return &PushRepositoryImpl{
		client: client,
	}
}

func (pr *PushRepositoryImpl) MessageLike(postID, userID int) error {
	// 略
	return nil
}

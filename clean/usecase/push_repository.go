package usecase

type PushRepository interface {
	MessageLike(postID, userID int) error
}

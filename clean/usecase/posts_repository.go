package usecase

type PostsRepository interface {
	LikePost(postID, userID int) error
}

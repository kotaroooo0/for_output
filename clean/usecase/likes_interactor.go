package usecase

type LikesInteractor interface {
	LikePost(postID, postedBy, likeUserID int) error
}

type LikesInteractorImpl struct {
	usersRepository UsersRepository
	postsRepository PostsRepository
	pushRepository  PushRepository
}

func NewLikesInteractor(ur UsersRepository, pr PostsRepository) *LikesInteractorImpl {
	return &LikesInteractorImpl{
		usersRepository: ur,
		postsRepository: pr,
	}
}

func (li *LikesInteractorImpl) LikePost(postID, likeUserID int) error {
	// いいね
	if err := li.postsRepository.LikePost(postID, likeUserID); err != nil {
		return err
	}

	// 通知を送信など
	user, err := li.usersRepository.Get(likeUserID)
	if err != nil {
		return err
	}
	// 略
	li.pushRepository.MessageLike(postID, user.Id)
	// 略

	return nil
}

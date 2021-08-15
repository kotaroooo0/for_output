package controller

type LikesController struct {
	likesInteractor usecase.LikesInteractor
}

func NewLikesController(li usecase.LikesInteractor) *LikesController {
	return &LikesController{
		likesInteractor: li,
	}
}

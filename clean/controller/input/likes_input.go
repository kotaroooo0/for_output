package controller

import (
	"errors"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

type LikesCreateInput struct {
	PostID int
	UserID int
}

func ParseLikesCreateInput(req events.APIGatewayProxyRequest) (*LikesCreateInput, error) {
	postID, ok := req.PathParameters["postId"]
	if !ok {
		return nil, errors.New("paramater postId is not found")
	}
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		return nil, err
	}

	userID, ok := req.QueryStringParameters["userId"]
	if !ok {
		return nil, errors.New("paramater userId is not found")
	}
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return nil, err
	}

	return &LikesCreateInput{
		PostID: postIDInt,
		UserID: userIDInt,
	}, nil
}

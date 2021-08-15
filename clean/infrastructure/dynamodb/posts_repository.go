package dynamodb

import (
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/kotaroooo0/for_output/clean/usecase/errors"
)

type PostsRepositoryImpl struct {
	db dynamodb.DynamoDB
}

func NewPostsRepository(db dynamodb.DynamoDB) *PostsRepositoryImpl {
	return &PostsRepositoryImpl{
		db: db,
	}
}

const postsTableName = "Posts"

var postAttributeNames = map[string]*string{
	"#ID": aws.String("Id"),
	"#L":  aws.String("LikedByIds"),
}

func (pr *PostsRepositoryImpl) LikePost(postID, userID int) error {
	input := &dynamodb.UpdateItemInput{
		TableName:                aws.String(postsTableName),
		ExpressionAttributeNames: postAttributeNames,
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":pid": {N: aws.String(strconv.Itoa(postID))},
			":uid": {N: aws.String(strconv.Itoa(userID))},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {N: aws.String(strconv.Itoa(postID))},
		},
		UpdateExpression: aws.String("ADD LikedByIds :uid"),
	}
	if _, err := pr.db.UpdateItem(input); err != nil {
		return errors.Wrap(err)
	}
	return nil
}

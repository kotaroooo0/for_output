package dynamodb

import (
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	domain "github.com/kotaroooo0/for_output/clean/domain/entity"
	"github.com/kotaroooo0/for_output/clean/usecase/errors"
)

type UsersRepositoryImpl struct {
	db dynamodb.DynamoDB
}

func NewUserRepository(db dynamodb.DynamoDB) *UsersRepositoryImpl {
	return &UsersRepositoryImpl{
		db: db,
	}
}

const usersTableName = "Users"

var usersTableFields = map[string]*string{
	"#ID": aws.String("Id"),
}

func (ur *UsersRepositoryImpl) Get(userID int) (*domain.User, error) {
	input := &dynamodb.GetItemInput{
		TableName:                aws.String(usersTableName),
		ExpressionAttributeNames: usersTableFields,
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {N: aws.String(strconv.Itoa(userID))},
		},
		ProjectionExpression: aws.String("Id"),
	}

	user := domain.User{}
	result, err := ur.db.GetItem(input)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	if len(result.Item) == 0 {
		return nil, &errors.AppError{
			Code:    errors.UserNotFoundCode,
			Message: "user item not found",
		}
	}
	if err := dynamodbattribute.UnmarshalMap(result.Item, &user); err != nil {
		return nil, errors.Wrap(err)
	}
	return &user, nil
}

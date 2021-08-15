package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kotaroooo0/for_output/clean/controller"
	"github.com/kotaroooo0/for_output/clean/infrastructure/dynamodb"
	"github.com/kotaroooo0/for_output/clean/usecase"
)

func main() {
	dynamodb := dynamodb.NewClient()
	ur := database.NewUserRepository(dynamoDBCli)
	pr := database.NewPostsRepository(dynamoDBCli)
	cr := database.NewCommunityRepository(dynamoDBCli)
	cur := database.NewCommunityUserRepository(dynamoDBCli)
	nr := database.NewNotificationRepository(dynamoDBCli)
	pur := messageing.NewPushRepository(firebaseCli, dynamoDBCli)
	tr := database.NewTableItemCounterRepository(dynamoDBCli)
	interactor := usecase.NewLikesInteractor(ur, pr, cr, cur, nr, pur, tr)
	controller := controller.NewLikesController(interactor)
	lambda.Start(controller.CreateLike)
}

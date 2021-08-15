package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func NewClient() *dynamodb.DynamoDB {
	config := aws.NewConfig().WithRegion("dummy").WithEndpoint("dummy")
	return dynamodb.New(session.Must(session.NewSession(config)))
}

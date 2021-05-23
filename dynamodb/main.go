package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	ctx := context.Background()
	config := aws.NewConfig().
		WithRegion("ap-northeast-1").
		WithEndpoint("http://127.0.0.1:8000").
		WithCredentials(credentials.NewStaticCredentials("dummy", "dummy", "dummy"))

	client := client{
		dynamodb.New(session.Must(session.NewSession(config))),
	}
	if err := client.update(ctx); err != nil {
		log.Fatalln(err)
	}
}

type client struct {
	dynamodb *dynamodb.DynamoDB
}

func (c *client) update(ctx context.Context) error {
	err := c.updateWithItem(ctx)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			if aerr.Code() == dynamodb.ErrCodeConditionalCheckFailedException {
				err = c.updateNoItem(ctx)
				if err != nil {
					if aerr, ok := err.(awserr.Error); ok {
						if aerr.Code() == dynamodb.ErrCodeConditionalCheckFailedException {
							err = c.updateWithItem(ctx)
						}
					}
				}
			}
		}
	}
	return err
}

func (c *client) updateWithItem(ctx context.Context) error {
	updateItemInput := &dynamodb.UpdateItemInput{
		TableName: aws.String("Users"),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {N: aws.String("0")},
		},
		ExpressionAttributeNames: map[string]*string{
			"#WEAPONS": aws.String("Weapons"),
			"#WEAPON":  aws.String("bow"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":st": {S: aws.String("normal")},
		},
		UpdateExpression:    aws.String("set #WEAPONS.#WEAPON = :st"),
		ConditionExpression: aws.String("attribute_exists(Id) and attribute_exists(Weapons)"),
	}
	_, err := c.dynamodb.UpdateItemWithContext(ctx, updateItemInput)
	return err
}

func (c *client) updateNoItem(ctx context.Context) error {
	updateItemInput := &dynamodb.UpdateItemInput{
		TableName: aws.String("Users"),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {N: aws.String("0")},
		},
		ExpressionAttributeNames: map[string]*string{
			"#WEAPONS": aws.String("Weapons"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":map": {M: map[string]*dynamodb.AttributeValue{
				"bow": {S: aws.String("normal")},
			}},
		},
		UpdateExpression:    aws.String("set #WEAPONS = :map"),
		ConditionExpression: aws.String("attribute_exists(Id) and attribute_not_exists(Weapons)"),
	}
	_, err := c.dynamodb.UpdateItemWithContext(ctx, updateItemInput)
	return err
}

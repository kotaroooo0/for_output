package firebase

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func NewMessageingClient() (*messaging.Client, error) {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("dummy"))
	if err != nil {
		return nil, err
	}
	return app.Messaging(ctx)
}

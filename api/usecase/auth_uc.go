package usecase

import (
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
)


func (u *Usecase) Initializing() *firebase.App {

	ctx := context.Background()
	opt := option.WithCredentialsFile("/Users/hirezakinamito/Downloads/term6-namito-hirezaki-firebase-adminsdk-n58or-2522cb339e.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return app
}


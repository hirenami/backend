package usecase

import (
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
	"os"
)


func (u *Usecase) Initializing() *firebase.App {

	credsFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
    if credsFile == "" {
        log.Fatalf("GOOGLE_APPLICATION_CREDENTIALS is not set")
    }

    ctx := context.Background()
    app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile(credsFile))
    if err != nil {
        log.Fatalf("error initializing app: %v\n", err)
    }

    return app
}


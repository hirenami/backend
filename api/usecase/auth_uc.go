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
    ctx := context.Background()

    var app *firebase.App
    var err error

    if credsFile == "" {
        // デフォルトの認証情報を使用 (GCP環境向け)
        app, err = firebase.NewApp(ctx, nil)
    } else {
        // ローカル開発環境向け
        app, err = firebase.NewApp(ctx, nil, option.WithCredentialsFile(credsFile))
    }

    if err != nil {
        log.Fatalf("Error initializing Firebase App: %v\n", err)
    }

    return app
}


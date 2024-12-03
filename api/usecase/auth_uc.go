package usecase

import (
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
	"os"
	"strings"
	"encoding/json"
)


func (u *Usecase) Initializing() *firebase.App {

	typeEnv := os.Getenv("TYPE")
	projectID := os.Getenv("PROJECT_ID")
	privateKeyID := os.Getenv("PRIVATE_KEY_ID")
	privateKey := os.Getenv("PRIVATE_KEY")
	clientEmail := os.Getenv("CLIENT_EMAIL")
	clientID := os.Getenv("CLIENT_ID")
	authURI := os.Getenv("AUTH_URI")
	tokenURI := os.Getenv("TOKEN_URI")
	authProviderCertURL := os.Getenv("AUTH_PROVIDER_X509_CERT_URL")
	clientCertURL := os.Getenv("CLIENT_X509_CERT_URL")
	universe_domain := os.Getenv("UNIVERSE_DOMAIN")

	// 必要な情報がすべて環境変数から取得できたかを確認
	if typeEnv == "" || projectID == "" || privateKeyID == "" || privateKey == "" || clientEmail == ""  || clientID == "" || authURI == "" || tokenURI == "" || authProviderCertURL == "" || clientCertURL == "" || universe_domain == "" {
		log.Println("Required environment variables are missing.")
	}

	// 秘密鍵を実際の改行文字に変換
	privateKey = strings.ReplaceAll(privateKey, "\\n", "\n")

	// JSON構造体を作成
	serviceAccount := map[string]interface{}{
		"type":                      typeEnv,
		"project_id":                projectID,
		"private_key_id":            privateKeyID,
		"private_key":               privateKey,
		"client_email":              clientEmail,
		"client_id":                 clientID,
		"auth_uri":                  authURI,
		"token_uri":                 tokenURI,
		"auth_provider_x509_cert_url": authProviderCertURL,
		"client_x509_cert_url":      clientCertURL,
		"universe_domain":           universe_domain,
	}

	// JSONにエンコード
	serviceAccountJSON, err := json.Marshal(serviceAccount)
	if err != nil {
		log.Printf("Failed to marshal service account JSON: %v", err)
		log.Println("error")
	}

	ctx := context.Background()
	opt := option.WithCredentialsJSON(serviceAccountJSON)
	
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Printf("error initializing app: %v\n", err)
		log.Println("error2")
	}
	return app
}


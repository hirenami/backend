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

	typeEnv := os.Getenv("_TYPE")
	projectID := os.Getenv("_PROJECT_ID")
	privateKeyID := os.Getenv("_PRIVATE_KEY_ID")
	privateKey := os.Getenv("_PRIVATE_KEY")
	clientEmail := os.Getenv("_CLIENT_EMAIL")
	clientID := os.Getenv("_CLIENT_ID")
	authURI := os.Getenv("_AUTH_URI")
	tokenURI := os.Getenv("_TOKEN_URI")
	authProviderCertURL := os.Getenv("_AUTH_PROVIDER_X509_CERT_URL")
	clientCertURL := os.Getenv("_CLIENT_X509_CERT_URL")
	universe_domain := os.Getenv("_UNIVERSE_DOMAIN")

	log.Println("typeEnv: ", typeEnv)
	log.Println("projectID: ", projectID)
	log.Println("privateKeyID: ", privateKeyID)
	log.Println("privateKey: ", privateKey)
	log.Println("clientEmail: ", clientEmail)
	log.Println("clientID: ", clientID)
	log.Println("authURI: ", authURI)
	log.Println("tokenURI: ", tokenURI)
	log.Println("authProviderCertURL: ", authProviderCertURL)
	log.Println("clientCertURL: ", clientCertURL)
	log.Println("universe_domain: ", universe_domain)


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



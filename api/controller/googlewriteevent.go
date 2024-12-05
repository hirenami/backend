package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2/google"
)

// UserEvent の構造体
type UserEvent struct {
	VisitorID     string            `json:"visitorId"`
	EventType     string            `json:"eventType"`
	ProductDetails []ProductDetail   `json:"productDetails"`
	EventTime     string            `json:"eventTime"`
}

type ProductDetail struct {
	ID string `json:"id"`
}

func (c *Controller) writedata(w http.ResponseWriter, r *http.Request) {
	// Google Cloud 認証情報の取得
	ctx := context.Background()
	credentials, err := google.FindDefaultCredentials(ctx, "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		log.Println("Error finding credentials:", err)
		http.Error(w, "failed to get credentials", http.StatusInternalServerError)
		return
	}

	tokenSource := credentials.TokenSource
	token, err := tokenSource.Token()
	if err != nil {
		log.Println("Error getting token:", err)
		http.Error(w, "failed to get token", http.StatusInternalServerError)
		return
	}

	listingId := mux.Vars(r)["listingId"]
	uid := r.Context().Value(uidKey).(string)
	visitId,err := c.Usecase.GetIdByUID(ctx,uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 検索リクエストの作成
	apiURL := "https://retail.googleapis.com/v2/projects/71857953091/locations/global/catalogs/default_catalog/userEvents:write"
	userEvent := UserEvent{
		VisitorID: visitId,
		EventType: "view-item",
		ProductDetails: []ProductDetail{
			{ID: listingId},
		},
		EventTime: time.Now().UTC().Format(time.RFC3339Nano), // 現在時刻を ISO 8601 形式で設定
	}

	body, err := json.Marshal(userEvent)
	if err != nil {
		log.Println("Error marshaling request body:", err)
		http.Error(w, "failed to marshal request body", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))
	if err != nil {
		log.Println("Error creating request:", err)
		http.Error(w, "failed to create request", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	req.Header.Set("x-goog-user-project" , "term6-namito-hirezaki")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	// リクエスト送信
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		http.Error(w, "failed to send request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// ステータスコードに応じてレスポンスを返す
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Error in search response", resp.StatusCode)
		return
	}

	// 正常に処理された場合は200 OKでレスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
}

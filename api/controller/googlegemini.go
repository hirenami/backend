package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"golang.org/x/oauth2/google"
)

func (c *Controller) Gemini(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Google Cloud 認証情報を取得
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

	query := r.URL.Query().Get("query")

	// API URL
	apiURL := "https://discoveryengine.googleapis.com/v1alpha/projects/71857953091/locations/global/collections/default_collection/engines/geminilast_1733597306757/servingConfigs/default_search:answer"

	// リクエストボディの準備
	requestBody := map[string]interface{}{
		"query": map[string]interface{}{
			"text":    query, // 質問内容
		},
		"relatedQuestionsSpec": map[string]interface{}{
			"enable": true, // 関連する質問を有効化
		},
		"answerGenerationSpec": map[string]interface{}{
			"ignoreAdversarialQuery":    true,
			"ignoreNonAnswerSeekingQuery": true,
			"ignoreLowRelevantContent":   true,
			"includeCitations":           true,
			"modelSpec": map[string]interface{}{
				"modelVersion": "preview",
			},
		},
		"pageSize": 1, // 結果を1件だけ取得
	}

	// JSON に変換
	body, err := json.Marshal(requestBody)
	if err != nil {
		log.Println("Error marshaling request body:", err)
		http.Error(w, "failed to marshal request body", http.StatusInternalServerError)
		return
	}

	// HTTP リクエストの作成
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))
	if err != nil {
		log.Println("Error creating request:", err)
		http.Error(w, "failed to create request", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	req.Header.Set("x-goog-user-project", "term6-namito-hirezaki")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	// HTTP クライアントでリクエスト送信
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		http.Error(w, "failed to send request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// レスポンスの読み取り
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response:", err)
		http.Error(w, "failed to read response", http.StatusInternalServerError)
		return
	}

	// レスポンスステータスの確認
	if resp.StatusCode != http.StatusOK {
		log.Printf("API Error: %s", string(respBody))
		http.Error(w, "Error in API response", resp.StatusCode)
		return
	}

	// レスポンスを出力
	fmt.Println("API Response:", string(respBody))

	// レスポンスをクライアントに返す
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}
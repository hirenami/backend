package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2/google"
)

func (c *Controller) searchProducts(w http.ResponseWriter, r *http.Request) {
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

	query := mux.Vars(r)["query"]
	visitId , ok := r.Context().Value(uidKey).(string)
	if !ok {
		http.Error(w, "Userid not found in context", http.StatusUnauthorized)
		return
	}

	// 検索リクエストの作成
	apiURL := "https://retail.googleapis.com/v2/projects/71857953091/locations/global/catalogs/default_catalog/placements/default_search:search"
	requestBody := map[string]interface{}{
		"query":    query, // 検索クエリ（例: "*" すべての商品）
		"visitorId": visitId, // ユーザーの訪問ID
		"pageSize": 10, // ページサイズ（1回のリクエストで返すアイテム数）
	}

	body, err := json.Marshal(requestBody)
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

	// レスポンスを読み取り
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response:", err)
		http.Error(w, "failed to read response", http.StatusInternalServerError)
		return
	}

	// 結果を表示
	fmt.Println("Search Response:")
	fmt.Println(string(respBody))

	// ステータスコードに応じてレスポンスを返す
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Error in search response", resp.StatusCode)
		return
	}

	// レスポンスがJSON形式なら、エンコードして返す
	var jsonResponse interface{}
	err = json.Unmarshal(respBody, &jsonResponse)
	if err != nil {
		log.Println("Error unmarshalling response:", err)
		http.Error(w, "failed to parse response", http.StatusInternalServerError)
		return
	}

	// 正常に処理された場合は200 OKでレスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(jsonResponse)
	if err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

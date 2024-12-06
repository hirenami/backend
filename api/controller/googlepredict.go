package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/oauth2/google"
)

func (c *Controller) GetPredicts(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Google Cloud 認証情報の取得
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

	uid := r.Context().Value(uidKey).(string)
	visitId, err := c.Usecase.GetIdByUID(ctx, uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 予測 API のリクエストボディ
	apiURL := "https://retail.googleapis.com/v2/projects/71857953091/locations/global/catalogs/default_catalog/servingConfigs/recently_viewed_default:predict"
	requestBody := map[string]interface{}{
		"userEvent": map[string]interface{}{
			"visitorId": visitId,
			"eventType": "detail-page-view",
			"eventTime": time.Now().UTC().Format(time.RFC3339Nano),
		},
		"pageSize": 5,
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
	req.Header.Set("x-goog-user-project", "term6-namito-hirezaki")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	// 予測 API 呼び出し
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		http.Error(w, "failed to send request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// 予測 API のレスポンスを処理
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response:", err)
		http.Error(w, "failed to read response", http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Error in predict response", resp.StatusCode)
		return
	}

	var predictResponse struct {
		Results []struct {
			Product struct {
				ID string `json:"id"`
			} `json:"product"`
		} `json:"results"`
	}
	err = json.Unmarshal(respBody, &predictResponse)
	if err != nil {
		log.Println("Error unmarshalling response:", err)
		http.Error(w, "failed to parse response", http.StatusInternalServerError)
		return
	}

	// API から取得した ID
	apiIds := []string{}
	for _, result := range predictResponse.Results {
		apiIds = append(apiIds, result.Product.ID)
	}

	// データベースからランダムに ID を取得
	dbIds, err := c.Usecase.GetRandomListingsUsecase(ctx) // データベースから最大10件取得
	if err != nil {
		http.Error(w, "failed to get product IDs from DB", http.StatusInternalServerError)
		return
	}

	// ID をマージしつつ重複を除外
	uniqueIds := map[string]struct{}{}
	mergedIds := []string{}

	for _, id := range apiIds {
		if _, exists := uniqueIds[id]; !exists {
			uniqueIds[id] = struct{}{}
			mergedIds = append(mergedIds, id)
		}
	}

	for _, id := range dbIds {
		Id := strconv.FormatInt(id, 10)
		if _, exists := uniqueIds[Id]; !exists {
			uniqueIds[Id] = struct{}{}
			mergedIds = append(mergedIds, Id)
		}
	}

	// 上位 5 件を返却
	if len(mergedIds) > 5 {
		mergedIds = mergedIds[:5]
	}

	// レスポンスとして返す
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"productIds": mergedIds,
	})
	if err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"golang.org/x/oauth2/google"
)

// Product構造体: 各商品データ
type Product struct {
	ID          string   `json:"id"`
	Categories  []string `json:"categories"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	PriceInfo   PriceInfo `json:"priceInfo"`
	Availability string   `json:"availability"`
	Uri         string   `json:"uri"`
	Images      []Image  `json:"images"`
}

// Image構造体: 商品の画像情報
type Image struct {
	Uri string `json:"uri"`
	Height int `json:"height"`
	Width int `json:"width"`
}

// PriceInfo構造体: 商品の価格情報
type PriceInfo struct {
	Price        float64 `json:"price"`
	CurrencyCode string  `json:"currencyCode"`
}

// InputConfig構造体: リクエストのルート構造
type InputConfig struct {
	ProductInlineSource ProductInlineSource `json:"productInlineSource"`
}

type Top struct {
	InputConfig InputConfig `json:"inputConfig"`
}


// ProductInlineSource構造体: 商品データの配列
type ProductInlineSource struct {
	Products []Product `json:"products"`
}

func (c *Controller)handleImportProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// リクエストボディの解析
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// APIに渡すデータを構築
	inputConfig := InputConfig{
		ProductInlineSource: ProductInlineSource{
			Products: []Product{product},
		},
	}
	top := Top{
		InputConfig: inputConfig,
	}

	body, err := json.Marshal(top)
	if err != nil {
		http.Error(w, "Failed to marshal payload", http.StatusInternalServerError)
		return
	}
	log.Printf("Request body: %s", body)

	// Google Cloud認証
	credentials, err := google.FindDefaultCredentials(r.Context(), "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		http.Error(w, "Failed to authenticate with Google Cloud", http.StatusInternalServerError)
		return
	}

	// アクセストークン取得
	tokenSource := credentials.TokenSource
	token, err := tokenSource.Token()
	if err != nil {
		http.Error(w, "Failed to retrieve access token", http.StatusInternalServerError)
		return
	}

	// Google Cloud Retail APIにリクエストを送信
	apiURL := "https://retail.googleapis.com/v2/projects/71857953091/locations/global/catalogs/default_catalog/branches/0/products:import"
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	req.Header.Set("x-goog-user-project" , "term6-namito-hirezaki")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to call Google API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Google Cloud APIからのレスポンスを返す
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(respBody)
}
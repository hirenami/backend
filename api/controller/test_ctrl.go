package controller

import (
	"api/sqlc"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func (c *Controller) Test(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		// プリフライトリクエストのヘッダーを設定
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	req := sqlc.Tweet{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	firebaseUid, ok := r.Context().Value(uidKey).(string)
	if !ok {
		http.Error(w, "Userid not found in context", http.StatusUnauthorized)
		return
	}
	ctx := context.Background()
	userId, err := c.Usecase.GetIdByUID(ctx, firebaseUid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	content := req.Content
	media_url := req.MediaUrl.String
	err = c.Usecase.CreateTweetUsecase(ctx, userId, content, media_url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Default().Println("User Userid: ", userId)
}

func (c *Controller) Test2(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		setCORSHeaders(w)
		w.WriteHeader(http.StatusOK)
		return
	}

	setCORSHeaders(w)

	ctx := context.Background()
	vars := mux.Vars(r)
	log.Printf("Vars: %v\n", vars) // デバッグ用に出
	userId := vars["userId"]
	tweet, err := c.Usecase.GetUsersTweetUsecase(ctx, userId)
	if err != nil {
		log.Printf("userId=%s", userId)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(tweet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
	log.Default().Println("User Userid: ", userId)
}

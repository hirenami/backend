package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

// POST /tweet
func (c *Controller) CreateTweetTweetCtrl(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		setCORSHeaders(w)
		w.WriteHeader(http.StatusOK)
		return
	}

	setCORSHeaders(w)

	req := Tweet{}
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
	media_url := req.MediaUrl
	err = c.Usecase.CreateTweetUsecase(ctx, userId, content, media_url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Default().Println("User Userid: ", userId)
}

// PUT /tweet/{tweetId}
func (c *Controller) UpdateTweetCtrl(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		setCORSHeaders(w)
		w.WriteHeader(http.StatusOK)
		return
	}

	setCORSHeaders(w)

	req := Tweet{}
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

	vars := mux.Vars(r)
	log.Printf("Vars: %v\n", vars) // デバッグ用に出
	tweetId := vars["tweetId"]
	TweetId, err := strconv.Atoi(tweetId) // strconv.Atoi は int を返す
	if err != nil {
		// 変換エラーの処理
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	content := req.Content
	media_url := req.MediaUrl
	err = c.Usecase.EditTweetUsecase(ctx, userId, int32(TweetId), content, media_url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Default().Println("User Userid: ", userId)
}

// DELETE /tweet/{tweetId}
func (c *Controller) DeleteTweetCtrl(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		setCORSHeaders(w)
		w.WriteHeader(http.StatusOK)
		return
	}

	setCORSHeaders(w)

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

	vars := mux.Vars(r)
	log.Printf("Vars: %v\n", vars) // デバッグ用に出
	tweetId := vars["tweetId"]
	TweetId, err := strconv.Atoi(tweetId) // strconv.Atoi は int を返す
	if err != nil {
		// 変換エラーの処理
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	err = c.Usecase.EraseTweetUsecase(ctx, userId, int32(TweetId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Default().Println("User Userid: ", userId)
}

// GET /tweet/{userId}
func (c *Controller) GetUsersTweetCtrl(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		setCORSHeaders(w)
		w.WriteHeader(http.StatusOK)
		return
	}

	setCORSHeaders(w)

	vars := mux.Vars(r)
	log.Printf("Vars: %v\n", vars) // デバッグ用に出
	userId := vars["userId"]

	firebaseUid, ok := r.Context().Value(uidKey).(string)
	if !ok {
		http.Error(w, "Userid not found in context", http.StatusUnauthorized)
		return
	}
	ctx := context.Background()
	Id, err := c.Usecase.GetIdByUID(ctx, firebaseUid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if bool,err :=c.Usecase.IsBlockedckUsecase(ctx, userId,Id)  ; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}else if bool {
		http.Error(w, "This user is blocked", http.StatusUnauthorized)
		return
	}
	
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

// GET /tweet
func (c *Controller) GetTweetCtrl(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		setCORSHeaders(w)
		w.WriteHeader(http.StatusOK)
		return
	}

	setCORSHeaders(w)

	vars := mux.Vars(r)
	log.Printf("Vars: %v\n", vars) // デバッグ用に出
	tweetId := vars["tweetId"]
	TweetId, err := strconv.Atoi(tweetId) // strconv.Atoi は int を返す
	if err != nil {
		// 変換エラーの処理
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	tweet, err := c.Usecase.GetTweetUsecase(ctx, int32(TweetId))
	if err != nil {
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
}
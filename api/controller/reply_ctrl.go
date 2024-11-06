package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"api/sqlc"
)

// POST /reply/{tweetId}

func (c *Controller) CreateReplyCtrl(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		setCORSHeaders(w)
		w.WriteHeader(http.StatusOK)
		return
	}

	setCORSHeaders(w)
	
	uid := r.Context().Value(uidKey).(string)
	ctx := context.Background()
	userId,err := c.Usecase.GetIdByUID(ctx,uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tweetId := mux.Vars(r)["tweetId"]
	TweetId, err := strconv.Atoi(tweetId) // strconv.Atoi は int を返す
	if err != nil {
		// 変換エラーの処理
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	var req Tweet
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	content := req.Content
	media_url := req.MediaUrl

	err = c.Usecase.CreateReplyUsecase(ctx, userId, content, media_url,int32(TweetId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GET /reply/{tweetId}

func (c *Controller) GetReplyCtrl(w http.ResponseWriter, r *http.Request) {
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
	Id, err := c.Usecase.GetIdByUID(ctx, firebaseUid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	tweetId := mux.Vars(r)["tweetId"]
	TweetId, err := strconv.Atoi(tweetId) // strconv.Atoi は int を返す
	if err != nil {
		// 変換エラーの処理
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	reply,user, islike, isretweet, err := c.Usecase.GetReplyUsecase(ctx,int32(TweetId),Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Tweet     []sqlc.Tweet
		User      []sqlc.User
		IsLike    []bool
		IsRetweet []bool
	}{
		Tweet:     reply,
		User:      user,
		IsLike:    islike,
		IsRetweet: isretweet,
	}

	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

// GET /reply/{tweetId}/replied

func (c *Controller) GetTweetRepliedToCtrl(w http.ResponseWriter, r *http.Request) {
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
	Id, err := c.Usecase.GetIdByUID(ctx, firebaseUid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	tweetId := mux.Vars(r)["tweetId"]
	TweetId, err := strconv.Atoi(tweetId) // strconv.Atoi は int を返す
	if err != nil {
		// 変換エラーの処理
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	reply,user, islike, isretweet, err := c.Usecase.GetTweetRepliedToUsecase(ctx,Id,int32(TweetId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Tweet     []sqlc.Tweet
		User      []sqlc.User
		IsLike    []bool
		IsRetweet []bool
	}{
		Tweet:     reply,
		User:      user,
		IsLike:    islike,
		IsRetweet: isretweet,
	}

	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
package controller

import (
	"net/http"
	"github.com/gorilla/mux"
	"context"
	"strconv"
	"encoding/json"
)

// POST /like/{tweetId}

func (c *Controller) CreateLikeCtrl(w http.ResponseWriter, r *http.Request) {
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
	err = c.Usecase.CreateLikeUsecase(ctx, userId, int32(TweetId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DELETE /like/{tweetId}

func (c *Controller) DeleteLikeCtrl(w http.ResponseWriter, r *http.Request) {
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
	err = c.Usecase.DeleteLikeUsecase(ctx, userId, int32(TweetId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *Controller) GetUserslikeCtrl (w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		setCORSHeaders(w)
		w.WriteHeader(http.StatusOK)
		return
	}

	setCORSHeaders(w)
	
	uid := r.Context().Value(uidKey).(string)
	ctx := context.Background()
	Id,err := c.Usecase.GetIdByUID(ctx,uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userId := mux.Vars(r)["userId"]
	likes, err := c.Usecase.GetUserslikeUsecase(ctx, Id, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(likes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"api/sqlc"
)

// GET /timeline
func (c *Controller) GetTimelineCtrl(w http.ResponseWriter, r *http.Request) {
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

	tweet, user, islike, isretweet, err  := c.Usecase.GetTimelineUsecase(ctx, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(struct {
		Tweet     []sqlc.Tweet
		User      []sqlc.User
		IsLike    []bool
		IsRetweet []bool
	}{
		Tweet:     tweet,
		User:      user,
		IsLike:    islike,
		IsRetweet: isretweet,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
package controller

import (
	"net/http"
	"context"
	"encoding/json"
	"api/sqlc"
)

// GET /notifications

func (c *Controller) GetNotificationsCtrl(w http.ResponseWriter, r *http.Request) {
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
	notifications,users,tweets,err := c.Usecase.GetNotificationsUsecase(ctx, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	response := struct {
		Notification []sqlc.Notification
		User []sqlc.User 
		Tweet []sqlc.Tweet
	}{
		Notification: notifications,
		User: users,
		Tweet: tweets,
	}
	
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
	

}
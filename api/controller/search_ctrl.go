package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

// GET /serach/{keyword}

func (c *Controller) SearchByKeywordCtrl(w http.ResponseWriter, r *http.Request) {
	keyword := mux.Vars(r)["keyword"]

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

	tweets, err := c.Usecase.SearchByKeywordUsecase(ctx,Id, keyword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tweets)

}

// GET /serach/{keyword}/user
func (c *Controller) SearchByUserCtrl(w http.ResponseWriter, r *http.Request) {
	keyword := mux.Vars(r)["keyword"]

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

	users, err := c.Usecase.SearchByUserUsecase(ctx,Id, keyword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

}

// GET /serach/{keyword}/tag
func (c *Controller) SearchByHashtagCtrl(w http.ResponseWriter, r *http.Request) {
	keyword := mux.Vars(r)["keyword"]

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

	tweets, err := c.Usecase.SearchByHashtagUsecase(ctx,Id, keyword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tweets)

}
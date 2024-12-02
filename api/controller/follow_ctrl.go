package controller

import (
	"context"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

// POST /follow/{userId}
func (c *Controller) CreateFollowCtrl(w http.ResponseWriter, r *http.Request) {
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

	followId := mux.Vars(r)["userId"]

	if bool,err :=c.Usecase.IsBlockedckUsecase(ctx, userId,followId)  ; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}else if bool {
		http.Error(w, "This user is blocked", http.StatusUnauthorized)
		return
	}

	err = c.Usecase.CreateFollowUsecase(ctx, userId, followId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Default().Println("User Userid: ", userId)
}

// DELETE /follow/{userId}
func (c *Controller) DeleteFollowCtrl(w http.ResponseWriter, r *http.Request) {
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

	followId := mux.Vars(r)["userId"]
	err = c.Usecase.DelateFollowUsecase(ctx, userId, followId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Default().Println("User Userid: ", userId)
}

// GET /follow/{userId}/following
func (c *Controller) GetFollowingCtrl(w http.ResponseWriter, r *http.Request) {
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

	userId := mux.Vars(r)["userId"]


	following, err := c.Usecase.GetFollowingUsecase(ctx,Id, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(following)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
	log.Default().Println("User Userid: ", userId)
}

// GET /follow/{userId}/follower
func (c *Controller) GetFollowerCtrl(w http.ResponseWriter, r *http.Request) {
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

	userId := mux.Vars(r)["userId"]
	
	follower, err := c.Usecase.GetFollowerUsecase(ctx, Id,userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(follower)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
	log.Default().Println("User Userid: ", userId)
}
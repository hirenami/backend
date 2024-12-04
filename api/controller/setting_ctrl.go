package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"api/model"
)

// POST /user/create
func (c *Controller) CreateAccount(w http.ResponseWriter, r *http.Request) {
	req := model.User{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	username := req.Username
	userId := req.Userid
	ctx := context.Background()
	firebaseUid, ok := r.Context().Value(uidKey).(string)
	if !ok {
		http.Error(w, "User Userid not found in context", http.StatusUnauthorized)
		log.Println("User Userid not found in context")
		return
	}
	err := c.Usecase.CreateAccount(ctx, firebaseUid, username, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	w.WriteHeader(http.StatusOK)
	log.Default().Println("User Userid: ", firebaseUid)
}

// PUT /user/edit
func (c *Controller) UpdateProfileCtrl(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value(uidKey).(string)
	ctx := context.Background()
	userId,err := c.Usecase.GetIdByUID(ctx,uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("User Userid: ", userId)

	req := model.User{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	err = c.Usecase.CreateProfileUsecase(ctx, userId, req.Username, req.Biography, req.HeaderImage, req.IconImage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

//PUT /user/isprivate

func (c *Controller) ChangePrivacyCtrl(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value(uidKey).(string)
	ctx := context.Background()
	userId,err := c.Usecase.GetIdByUID(ctx,uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var isPrivate bool
	err = json.NewDecoder(r.Body).Decode(&isPrivate)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	err = c.Usecase.UpdatePrivateUsecase(ctx, userId, isPrivate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
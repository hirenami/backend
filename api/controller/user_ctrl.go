package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

// GET /user/{userId}
func (c *Controller) GetProfileCtrl(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		setCORSHeaders(w)
		w.WriteHeader(http.StatusOK)
		return
	}

	setCORSHeaders(w)
	userId := mux.Vars(r)["userId"]

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

	profile, err := c.Usecase.GetProfileUsecase(ctx,Id, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)

}

func (c *Controller) GetMyProfileCtrl(w http.ResponseWriter, r *http.Request) {
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


	profile,err := c.Usecase.GetProfileUsecase(ctx,Id,Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)

	
}
package controller

import (
	"context"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"encoding/json"
)

func (c *Controller) UpdatePremiumCtrl(w http.ResponseWriter, r *http.Request) {
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

	err = c.Usecase.UpdatePremiumUsecase(ctx, Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *Controller) GetPurchaseCtrl(w http.ResponseWriter, r *http.Request) {
	purchaseId := mux.Vars(r)["purchaseId"]
	purchaseid, err := strconv.Atoi(purchaseId)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	ctx := context.Background()

	purchase, err := c.Usecase.GetPurchaseUsecase(ctx, int32(purchaseid))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(purchase)
}

func (c *Controller) GetMyPurchaseCtrl(w http.ResponseWriter, r *http.Request) {
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

	purchases, err := c.Usecase.GetPurchasesByUserUsecase(ctx, Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(purchases)
}

func (c *Controller) CreatePurchaseCtrl(w http.ResponseWriter, r *http.Request) {
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

	var listingId int64
	err = json.NewDecoder(r.Body).Decode(&listingId)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = c.Usecase.CreatePurchaseUsecase(ctx, userId, listingId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

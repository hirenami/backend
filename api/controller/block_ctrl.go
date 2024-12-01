package controller

import (
	"context"
	"errors"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

func (c *Controller) CreateBlockCtrl(w http.ResponseWriter, r *http.Request) {
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

	blockId := mux.Vars(r)["blockId"]

	// ユースケースのメソッドを呼び出し
	err = c.Usecase.CreateBlockUsecase(ctx, Id, blockId)
	if err != nil {
		if errors.Is(err, errors.New("user does not exist")) {
			http.Error(w, "user does not exist", http.StatusBadRequest)
			return
		} else if errors.Is(err, errors.New("block user does not exist")) {
			http.Error(w, "block user does not exist", http.StatusBadRequest)
			return
		}
		log.Printf("error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *Controller) DeleteBlockCtrl(w http.ResponseWriter, r *http.Request) {
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

	blockId := mux.Vars(r)["blockId"]

	// ユースケースのメソッドを呼び出し
	err = c.Usecase.DeleteBlockUsecase(ctx, Id, blockId)
	if err != nil {
		if errors.Is(err, errors.New("user does not exist")) {
			http.Error(w, "user does not exist", http.StatusBadRequest)
			return
		} else if errors.Is(err, errors.New("block user does not exist")) {
			http.Error(w, "block user does not exist", http.StatusBadRequest)
			return
		}
		log.Printf("error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *Controller) GetBlocksCtrl(w http.ResponseWriter, r *http.Request) {
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

	blocks, err := c.Usecase.GetBlocksUsecase(ctx, Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(blocks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}

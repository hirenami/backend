package controller

import (
	"context"
	"errors"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func (c *Controller) CreateFollowRequestCtrl (w http.ResponseWriter, r *http.Request) {
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

	followId := mux.Vars(r)["followId"]

	// ユースケースのメソッドを呼び出し
	err = c.Usecase.CreateKeyFollowUsecase(ctx, Id, followId)
	if err != nil {
		if errors.Is(err, errors.New("user does not exist")) {
			http.Error(w, "user does not exist", http.StatusBadRequest)
			return
		} else if errors.Is(err, errors.New("follow user does not exist")) {
			http.Error(w, "follow user does not exist", http.StatusBadRequest)
			return
		} else if errors.Is(err, errors.New("can't follow yourself")) {
			http.Error(w, "can't follow yourself", http.StatusBadRequest)
			return
		}
		log.Printf("error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *Controller) DeleteFollowRequestCtrl (w http.ResponseWriter, r *http.Request) {
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

	followId := mux.Vars(r)["followId"]

	// ユースケースのメソッドを呼び出し
	err = c.Usecase.DeleteKeyFollowUsecase(ctx, Id, followId)
	if err != nil {
		if errors.Is(err, errors.New("user does not exist")) {
			http.Error(w, "user does not exist", http.StatusBadRequest)
			return
		} else if errors.Is(err, errors.New("follow user does not exist")) {
			http.Error(w, "follow user does not exist", http.StatusBadRequest)
			return
		}
		log.Printf("error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *Controller) ApproveRequestCtrl (w http.ResponseWriter, r *http.Request) {
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

	followId := mux.Vars(r)["followId"]

	// ユースケースのメソッドを呼び出し
	err = c.Usecase.ApproveRequest(ctx, Id, followId)
	if err != nil {
		if errors.Is(err, errors.New("user does not exist")) {
			http.Error(w, "user does not exist", http.StatusBadRequest)
			return
		} else if errors.Is(err, errors.New("follow user does not exist")) {
			http.Error(w, "follow user does not exist", http.StatusBadRequest)
			return
		} else if errors.Is(err, errors.New("can't follow yourself")) {
			http.Error(w, "can't follow yourself", http.StatusBadRequest)
			return
		}
		log.Printf("error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *Controller) RejectRequestCtrl (w http.ResponseWriter, r *http.Request) {
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

	followId := mux.Vars(r)["followId"]

	// ユースケースのメソッドを呼び出し
	err = c.Usecase.RejectRequest(ctx, Id, followId)
	if err != nil {
		if errors.Is(err, errors.New("user does not exist")) {
			http.Error(w, "user does not exist", http.StatusBadRequest)
			return
		} else if errors.Is(err, errors.New("follow user does not exist")) {
			http.Error(w, "follow user does not exist", http.StatusBadRequest)
			return
		} else if errors.Is(err, errors.New("can't follow yourself")) {
			http.Error(w, "can't follow yourself", http.StatusBadRequest)
			return
		}
		log.Printf("error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

// GET /serach/{keyword}

func (c *Controller) SearchByKeywordCtrl(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		setCORSHeaders(w)
		w.WriteHeader(http.StatusOK)
		return
	}

	setCORSHeaders(w)
	keyword := mux.Vars(r)["keyword"]

	ctx := context.Background()
	tweets, err := c.Usecase.SearchByKeywordUsecase(ctx, keyword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tweets)

}
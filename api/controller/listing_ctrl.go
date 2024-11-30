package controller

import (
	"net/http"
	"github.com/gorilla/mux"
	"context"
	"strconv"
	"encoding/json"
	"api/model"
)

func (c *Controller) GetListing(w http.ResponseWriter, r *http.Request) {
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

	listingId := mux.Vars(r)["listingId"]
	ListingId, err := strconv.Atoi(listingId) // strconv.Atoi は int を返す
	if err != nil {
		// 変換エラーの処理
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	listing, err := c.Usecase.GetListingUsecase(ctx, userId,int64(ListingId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(listing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func (c *Controller) GetListingByTweet(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		setCORSHeaders(w)
		w.WriteHeader(http.StatusOK)
		return
	}

	setCORSHeaders(w)

	tweetId := mux.Vars(r)["tweetId"]
	TweetId, err := strconv.Atoi(tweetId) // strconv.Atoi は int を返す
	if err != nil {
		// 変換エラーの処理
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	listing, err := c.Usecase.GetListingByTweetUsecase(ctx, int32(TweetId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(listing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func (c *Controller) CreateListing(w http.ResponseWriter, r *http.Request) {
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

	type CreateListing struct {
		Listing model.Listing `json:"listing"`
		Content string `json:"content"`
		Media_url string `json:"media_url"`
	}
	var listing CreateListing
	if err := json.NewDecoder(r.Body).Decode(&listing); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.Usecase.CreateListingUsecase(ctx, int64(listing.Listing.Listingid), userId,listing.Content, listing.Media_url, listing.Listing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *Controller) GetUserListings(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		setCORSHeaders(w)
		w.WriteHeader(http.StatusOK)
		return
	}

	setCORSHeaders(w)

	userId := mux.Vars(r)["userId"]
	ctx := context.Background()
	listings, err := c.Usecase.GetUserListingsUsecase(ctx, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(listings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func (c *Controller) GetMyListingCtrl(w http.ResponseWriter, r *http.Request) {
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
	listings, err := c.Usecase.GetUserListingsUsecase(ctx, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(listings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
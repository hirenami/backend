package controller

import (
	"net/http"
	"github.com/gorilla/mux"
)

func SetupRoutes(controller *Controller) *mux.Router {
	r := mux.NewRouter()

	r.Handle("/timeline", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetTimelineCtrl))).Methods("GET","OPTIONS")

	r.Handle("/tweet", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.CreateTweetTweetCtrl))).Methods("POST","OPTIONS")
	r.Handle("/tweet/{tweetId}/tweetid", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetTweetCtrl))).Methods("GET","OPTIONS")
	r.Handle("/tweet/{tweetId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.UpdateTweetCtrl))).Methods("PUT","OPTIONS")
	r.Handle("/tweet/{tweetId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.DeleteTweetCtrl))).Methods("PATCH","OPTIONS")
	r.Handle("/tweet/{userId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetUsersTweetCtrl))).Methods("GET","OPTIONS")

	r.Handle("/retweet/{tweetId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.IsRetweetCtrl))).Methods("GET","OPTIONS")
	r.Handle("/retweet/{tweetId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.CreateRetweetCtrl))).Methods("POST","OPTIONS")
	r.Handle("/retweet/{tweetId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.DeleteRetweetCtrl))).Methods("DELETE","OPTIONS")
	r.Handle("/retweet/{tweetId}/quote", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.CreateQuoteCtrl))).Methods("POST","OPTIONS")

	r.Handle("/follow/{userId}/following", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetFollowingCtrl))).Methods("GET","OPTIONS")
	r.Handle("/follow/{userId}/follower", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetFollowerCtrl))).Methods("GET","OPTIONS")
	r.Handle("/follow/{userId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.IsFollowingCtrl))).Methods("GET","OPTIONS")
	r.Handle("/follow/{userId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.CreateFollowCtrl))).Methods("POST","OPTIONS")
	r.Handle("/follow/{userId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.DeleteFollowCtrl))).Methods("DELETE","OPTIONS")

	r.Handle("/user/{userId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetProfileCtrl))).Methods("GET","OPTIONS")
	r.Handle("/user", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetMyProfileCtrl))).Methods("GET","OPTIONS")
	r.Handle("/user/create", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.CreateAccount))).Methods("POST","OPTIONS")
	r.Handle("/user/edit", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.UpdateProfileCtrl))).Methods("PUT","OPTIONS")
	r.Handle("/user/delete", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.DeleteAccountCtrl))).Methods("PATCH","OPTIONS")

	r.Handle("/like/{tweetId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.CreateLikeCtrl))).Methods("POST","OPTIONS")
	r.Handle("/like/{tweetId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.DeleteLikeCtrl))).Methods("DELETE","OPTIONS")
	r.Handle("/like/{tweetId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.IsLikedCtrl))).Methods("GET","OPTIONS")



	r.Handle("/test", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.Test)))
	r.Handle("/test2/{userId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.Test2)))
	r.HandleFunc("/login", controller.Login)
	return r
}

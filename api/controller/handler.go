package controller

import (
	"net/http"
	"github.com/gorilla/mux"
)

func SetupRoutes(controller *Controller) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/import-products", controller.handleImportProducts).Methods("POST","OPTIONS")
	r.Handle("/api/search-products/{query}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.searchProducts))).Methods("GET","OPTIONS")

	r.Handle("/timeline", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetTimelineCtrl))).Methods("GET","OPTIONS")

	r.Handle("/tweet", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.CreateTweetTweetCtrl))).Methods("POST","OPTIONS")
	r.Handle("/tweet/{tweetId}/tweetid", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetTweetCtrl))).Methods("GET","OPTIONS")
	r.Handle("/tweet/{tweetId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.UpdateTweetCtrl))).Methods("PUT","OPTIONS")
	r.Handle("/tweet/{tweetId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.DeleteTweetCtrl))).Methods("DELETE","OPTIONS")
	r.Handle("/tweet/{userId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetUsersTweetCtrl))).Methods("GET","OPTIONS")

	r.Handle("/retweet/{tweetId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.IsRetweetCtrl))).Methods("GET","OPTIONS")
	r.Handle("/retweet/{tweetId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.CreateRetweetCtrl))).Methods("POST","OPTIONS")
	r.Handle("/retweet/{tweetId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.DeleteRetweetCtrl))).Methods("DELETE","OPTIONS")
	r.Handle("/retweet/{tweetId}/quote", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.CreateQuoteCtrl))).Methods("POST","OPTIONS")

	r.Handle("/follow/{userId}/following", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetFollowingCtrl))).Methods("GET","OPTIONS")
	r.Handle("/follow/{userId}/follower", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetFollowerCtrl))).Methods("GET","OPTIONS")
	r.Handle("/follow/{userId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.CreateFollowCtrl))).Methods("POST","OPTIONS")
	r.Handle("/follow/{userId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.DeleteFollowCtrl))).Methods("DELETE","OPTIONS")

	r.Handle("/user/{userId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetProfileCtrl))).Methods("GET","OPTIONS")
	r.Handle("/user", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetMyProfileCtrl))).Methods("GET","OPTIONS")
	r.HandleFunc("/user/create", controller.CreateAccount).Methods("POST","OPTIONS")
	r.Handle("/user/edit", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.UpdateProfileCtrl))).Methods("PUT","OPTIONS")
	r.Handle("/user/delete", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.DeleteAccountCtrl))).Methods("PATCH","OPTIONS")
	r.Handle("/user/private", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.ChangePrivacyCtrl))).Methods("PUT","OPTIONS")
	r.Handle("/premium", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.UpdatePremiumCtrl))).Methods("PATCH","OPTIONS")

	r.Handle("/reply/{tweetId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.CreateReplyCtrl))).Methods("POST","OPTIONS")
	r.Handle("/reply/{tweetId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetReplyCtrl))).Methods("GET","OPTIONS")
	r.Handle("/reply/{tweetId}/replied", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetTweetRepliedToCtrl))).Methods("GET","OPTIONS")
	r.Handle("/reply/{userId}/user", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetUsersReplyCtrl))).Methods("GET","OPTIONS")

	r.Handle("/like/{tweetId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.CreateLikeCtrl))).Methods("POST","OPTIONS")
	r.Handle("/like/{tweetId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.DeleteLikeCtrl))).Methods("DELETE","OPTIONS")
	r.Handle("/like/{userId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetUserslikeCtrl))).Methods("GET","OPTIONS")

	r.Handle("/notifications", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetNotificationsCtrl))).Methods("GET","OPTIONS")
	r.Handle("/notifications/{notificationId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.UpdateNotificationStatusCtrl))).Methods("PUT","OPTIONS")

	r.Handle("/search/{keyword}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.SearchByKeywordCtrl))).Methods("GET","OPTIONS")
	r.Handle("/search/{keyword}/user", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.SearchByUserCtrl))).Methods("GET","OPTIONS")
	r.Handle("/search/{keyword}/hashtag", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.SearchByHashtagCtrl))).Methods("GET","OPTIONS")

	r.Handle("/listing/{listingId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetListing))).Methods("GET","OPTIONS")
	r.Handle("/listing/{tweetId}/tweetid", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetListingByTweet))).Methods("GET","OPTIONS")
	r.Handle("/listing/{userId}/userid", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetUserListings))).Methods("GET","OPTIONS")
	r.Handle("/listing", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.CreateListing))).Methods("POST","OPTIONS")
	r.Handle("/listing", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetMyListingCtrl))).Methods("GET","OPTIONS")

	r.Handle("/purchase/{purchaseId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetPurchaseCtrl))).Methods("GET","OPTIONS")
	r.Handle("/purchase", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.CreatePurchaseCtrl))).Methods("POST","OPTIONS")
	r.Handle("/purchase", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetMyPurchaseCtrl))).Methods("GET","OPTIONS")

	r.HandleFunc("/dm/{userId}", controller.handleConnection).Methods("GET","OPTIONS")
	r.Handle("/dm", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetAllDmsCtrl))).Methods("GET","OPTIONS")
	r.Handle("/dm", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.CreateDm))).Methods("POST","OPTIONS")
	r.Handle("/dm/{userId}/handle", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetDmsCtrl))).Methods("GET","OPTIONS")

	r.Handle("/block/{blockId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.CreateBlockCtrl))).Methods("POST","OPTIONS")
	r.Handle("/block/{blockId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.DeleteBlockCtrl))).Methods("DELETE","OPTIONS")
	r.Handle("/block", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetBlocksCtrl))).Methods("GET","OPTIONS")

	r.Handle("/keyfollow/{followId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.CreateFollowRequestCtrl))).Methods("POST","OPTIONS")
	r.Handle("/keyfollow/{followId}", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.DeleteFollowRequestCtrl))).Methods("DELETE","OPTIONS")
	r.Handle("/keyfollow", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.GetFollowRequestsCtrl))).Methods("GET","OPTIONS")
	r.Handle("/keyfollow/{followId}/approve", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.ApproveRequestCtrl))).Methods("POST","OPTIONS")
	r.Handle("/keyfollow/{followId}/reject", controller.FirebaseAuthMiddleware()(http.HandlerFunc(controller.RejectRequestCtrl))).Methods("DELETE","OPTIONS")

	r.HandleFunc("/login", controller.Login)
	return r
}

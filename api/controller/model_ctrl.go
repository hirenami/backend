package controller

import (
	"api/usecase"
	"net/http"
)

type Controller struct {
	Usecase *usecase.Usecase
}

func NewController(usecase *usecase.Usecase) *Controller {
	return &Controller{
		Usecase: usecase,
	}
}

type ContextKey string

const (
	uidKey ContextKey = "firebaseUid" // ユーザーID用
	idKey  ContextKey = "userId"      // その他のID用
)

func setCORSHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

type UpdateProfile struct {
	Username    string `json:"username"`
	HeaderImage string `json:"header_image"`
	IconImage   string `json:"icon_image"`
	Biography   string `json:"biography"`
}

type User struct {
	Firebaseuid string `json:"firebaseuid"`
	Userid      string `json:"userid"`
	Username    string `json:"username"`
	HeaderImage string `json:"header_image"`
	IconImage   string `json:"icon_image"`
	Biography   string `json:"biography"`
	Isprivate   bool   `json:"isprivate"`
	Isfrozen    bool   `json:"isfrozen"`
	Isdeleted   bool   `json:"isdeleted"`
	Isadmin     bool   `json:"isadmin"`
}

type Tweet struct {
	Tweetid     int32  `json:"tweetid"`
	Userid      string `json:"userid"`
	Retweetid   int32  `json:"retweetid"`
	Isquote     bool   `json:"isquote"`
	Isreply     bool   `json:"isreply"`
	Content     string `json:"content"`
	MediaUrl    string `json:"media_url"`
	Likes       int32  `json:"likes"`
	Retweets    int32  `json:"retweets"`
	Replies     int32  `json:"replies"`
	Impressions int32  `json:"impressions"`
	Isdeleted   bool   `json:"isdeleted"`
}

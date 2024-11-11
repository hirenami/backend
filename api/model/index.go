package model

import (
	"api/sqlc"
	"time"
)

type TweetParams struct {
	Tweet    sqlc.Tweet `json:"tweet"`
	User     sqlc.User  `json:"user"`
	Likes    bool       `json:"likes"`
	Retweets bool       `json:"retweets"`
}

type NotificationParams struct {
	Notification sqlc.Notification `json:"notification"`
	User         sqlc.User         `json:"user"`
	Tweet        sqlc.Tweet        `json:"tweet"`
}

type Tweet struct {
	Tweetid     int32     `json:"tweetid"`
	Userid      string    `json:"userid"`
	Retweetid   int32     `json:"retweetid"`
	Isquote     bool      `json:"isquote"`
	Isreply     bool      `json:"isreply"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Content     string    `json:"content"`
	MediaUrl    string    `json:"media_url"`
	Likes       int32     `json:"likes"`
	Retweets    int32     `json:"retweets"`
	Replies     int32     `json:"replies"`
	Impressions int32     `json:"impressions"`
	Isdeleted   bool      `json:"isdeleted"`
}

type User struct {
	Firebaseuid string    `json:"firebaseuid"`
	Userid      string    `json:"userid"`
	Username    string    `json:"username"`
	CreatedAt   time.Time `json:"created_at"`
	HeaderImage string    `json:"header_image"`
	IconImage   string    `json:"icon_image"`
	Biography   string    `json:"biography"`
	Isprivate   bool      `json:"isprivate"`
	Isfrozen    bool      `json:"isfrozen"`
	Isdeleted   bool      `json:"isdeleted"`
	Isadmin     bool      `json:"isadmin"`
}

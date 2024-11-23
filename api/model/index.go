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
	Isblocked   bool      `json:"isblocked"`
	Isprivate   bool      `json:"isprivate"`
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
	Review      int32     `json:"review"`
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
	Ispremium   bool      `json:"ispremium"`
	Isdeleted   bool      `json:"isdeleted"`
	Isadmin     bool      `json:"isadmin"`
}

type UpdateProfile struct {
	Username    string `json:"username"`
	HeaderImage string `json:"header_image"`
	IconImage   string `json:"icon_image"`
	Biography   string `json:"biography"`
}

type Profile struct {
	User        sqlc.User `json:"user"`
	Follows     int32     `json:"follows"`
	Followers   int32     `json:"followers"`
	Isfollows   bool      `json:"isfollows"`
	Isfollowers bool      `json:"isfollowers"`
	Isblocked   bool      `json:"isblocked"`
	Isprivate   bool      `json:"isprivate"`
}

type Conversation struct {
	User sqlc.User    // ユーザー情報
	Dms  []sqlc.Dm    // DM一覧
}

type ListingParams struct {
	Listing  sqlc.Listing `json:"listing"`
	User     sqlc.User    `json:"user"`
	Tweet    sqlc.Tweet   `json:"tweet"`
}

type PurchaseParams struct {
	Purchase sqlc.Purchase `json:"purchase"`
	Listing  sqlc.Listing  `json:"listing"`
	User     sqlc.User     `json:"user"`
	Tweet	sqlc.Tweet    `json:"tweet"`
}

type Listing struct {
	Listingid        int32     `json:"listingid"`
	Userid           string    `json:"userid"`
	Tweetid          int32     `json:"tweetid"`
	CreatedAt        time.Time `json:"created_at"`
	Listingname      string    `json:"listingname"`
	Listingdescription string    `json:"listingdescription"`
	Listingprice     int32     `json:"listingprice"`
	Type             string    `json:"type"`
	Stock            int32     `json:"stock"`
	Condition        string    `json:"condition"`
}

type Purchase struct {
	Purchaseid int32     `json:"purchaseid"`
	Userid     string    `json:"userid"`
	Listingid  int32     `json:"listingid"`
	CreatedAt  time.Time `json:"created_at"`
	Status     string    `json:"status"`
}
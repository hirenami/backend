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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

type UpdateProfile struct {
	Username    string `json:"username"`
	HeaderImage string `json:"header_image"`
	IconImage   string `json:"icon_image"`
	Biography   string `json:"biography"`
}
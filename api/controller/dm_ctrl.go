package controller

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var clients = make(map[string]*websocket.Conn) // ユーザーID -> WebSocket接続
var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

// WebSocket接続時にユーザーIDをクエリパラメータで受け取る
func (c *Controller) handleConnection(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータからユーザーIDを取得
	userID := mux.Vars(r)["userId"]

	// WebSocket接続のアップグレード
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// ユーザーIDをキーにしてWebSocket接続をマップに追加
	clients[userID] = conn
	log.Printf("User %s connected", userID)

	type Message struct {
		SenderID   string `json:"senderId"`
		ReceiverID string `json:"receiverId"`
		Content    string `json:"content"`
	}
	log.Println("conn" , clients)
	// メッセージ受信処理
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			delete(clients, userID) // 接続が切れた場合、マップから削除
			return
		}
		

		// メッセージをJSONとしてパース
		var msg Message
		err = json.Unmarshal(p, &msg)
		if err != nil {
			log.Println("Error unmarshalling message:", err)
			continue
		}
		log.Println("msg", msg)
		

		// メッセージの送信先ユーザーID（receiverId）
		receiverConn, ok := clients[msg.ReceiverID]
		if ok {
			err = receiverConn.WriteMessage(messageType, p)
			if err != nil {
				log.Println(err)
			}
			log.Println("receiverConn", msg.ReceiverID)
		} else {
			log.Printf("Receiver %s not connected", msg.ReceiverID)
		}
		senderConn, ok := clients[msg.SenderID]
		if ok {
			err = senderConn.WriteMessage(messageType, p)
			if err != nil {
				log.Println(err)
			}
			log.Println("senderConn", msg.SenderID)
		} else {
			log.Printf("Sender %s not connected", msg.SenderID)
		}
	}
}

func (c *Controller) GetAllDmsCtrl(w http.ResponseWriter, r *http.Request) {
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

	dms, err := c.Usecase.GetAllDms(ctx, Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dms)
}

func (c *Controller) CreateDm(w http.ResponseWriter, r *http.Request) {
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

	type Request struct {
		ReceiverId string `json:"receiverId"`
		Content    string `json:"content"`
		MediaUrl   string `json:"media_url"`
	}
	var req Request
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.Usecase.CreateDm(ctx, Id, req.ReceiverId, req.Content, req.MediaUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK) 

}

func (c *Controller) GetDmsCtrl(w http.ResponseWriter, r *http.Request) {
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

	vars := mux.Vars(r)
	userId := vars["userId"]

	dms, err := c.Usecase.GetDms(ctx, Id, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dms)
}

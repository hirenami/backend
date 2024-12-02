package controller

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var clients = sync.Map{}
var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

// WebSocket接続時にユーザーIDをクエリパラメータで受け取る
func (c *Controller) handleConnection(w http.ResponseWriter, r *http.Request) {
    userID := mux.Vars(r)["userId"]
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    defer func() {
        clients.Delete(userID)
        conn.Close()
    }()

    clients.Store(userID, conn)
    log.Printf("User %s connected", userID)

    type Message struct {
        SenderID   string `json:"senderId"`
        ReceiverID string `json:"receiverId"`
        Content    string `json:"content"`
    }

    for {
        _, p, err := conn.ReadMessage()
        if err != nil {
            log.Println("ReadMessage error:", err)
            return
        }

        var msg Message
        if err := json.Unmarshal(p, &msg); err != nil {
            log.Println("Error unmarshalling message:", err)
            continue
        }

        if receiverConn, ok := clients.Load(msg.ReceiverID); ok {
            if err := receiverConn.(*websocket.Conn).WriteMessage(websocket.TextMessage, p); err != nil {
                log.Println("Error sending message to receiver:", err)
            }
        } else {
            log.Printf("Receiver %s not connected", msg.ReceiverID)
        }

        if senderConn, ok := clients.Load(msg.SenderID); ok {
            if err := senderConn.(*websocket.Conn).WriteMessage(websocket.TextMessage, p); err != nil {
                log.Println("Error sending message to sender:", err)
            }
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

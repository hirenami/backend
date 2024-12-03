package controller

import (
	"context"
	"log"
	"net/http"
	"encoding/json"
)

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		setCORSHeaders(w)
		w.WriteHeader(http.StatusOK)
		return
	}

	setCORSHeaders(w)

	app := c.Usecase.Initializing()
	ctx := context.Background()
	client, err := app.Auth(ctx)

	if err != nil {
		http.Error(w, "Firebase auth initialization error", http.StatusInternalServerError)
		return
	}

	// フロントエンドから送られたIDトークンを取得
	idToken := r.Header.Get("Authorization")

	// `Bearer ` プレフィックスを除去する処理が必要かもしれない
	if len(idToken) > 7 && idToken[:7] == "Bearer " {
		idToken = idToken[7:]
	}

	if idToken == "" {
		http.Error(w, "Authorization token not found", http.StatusUnauthorized)
		return
	}

	// トークンの検証
	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Printf("idtoken = %s", idToken)
		http.Error(w, "Invalid tokenss", http.StatusUnauthorized)
		return
	}

	// トークンが正しい場合、ユーザー情報を取得
	log.Printf("Verified Userid token: %v\n", idToken)
	log.Printf("User Userid: %v\n", token.UID)

	Id, err := c.Usecase.GetIdByUID(ctx, token.UID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ここでDB処理や他のアクションを行う
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Id)

}

// Firebase認証を行うミドルウェア
func (c *Controller) FirebaseAuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodOptions {
				setCORSHeaders(w)
				w.WriteHeader(http.StatusOK)
				return
			}

			setCORSHeaders(w)

			// Firebase 初期化
			app := c.Usecase.Initializing()
			ctx := r.Context()
			client, err := app.Auth(ctx)

			if err != nil {
				http.Error(w, "Firebase auth initialization error", http.StatusInternalServerError)
				return
			}

			// フロントエンドから送られたIDトークンを取得
			idToken := r.Header.Get("Authorization")
			if len(idToken) > 7 && idToken[:7] == "Bearer " {
				idToken = idToken[7:]
			}

			if idToken == "" {
				http.Error(w, "Authorization token not found", http.StatusUnauthorized)
				return
			}

			// トークンの検証
			// token, err := client.VerifyIDToken(ctx, idToken)
			// if err != nil {
			// 	log.Printf("Invalid idToken: %s, Error: %v", idToken, err)
			// 	http.Error(w, "Invalid token", http.StatusUnauthorized)
			// 	return
			// }

			// // UIDをコンテキストに追加
			// ctx = context.WithValue(ctx, uidKey, token.UID)

			// 次のハンドラーを呼び出す
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

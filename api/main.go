package main

import (
	"api/controller"
	"api/dao"
	"api/sqlc"
	"api/usecase"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"context"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"cloud.google.com/go/cloudsqlconn"
)

func envload() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {
	envload()
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlUserPwd := os.Getenv("MYSQL_PASSWORD")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
    mysqlHost := os.Getenv("MYSQL_HOST")

	// Cloud SQL Auth Proxyの設定
	dialer, err := cloudsqlconn.NewDialer(context.Background())
	if err != nil {
		log.Fatalf("Could not create new dialer: %v", err)
	}
	defer dialer.Close()

	// MySQL用のDSNを作成
	dsn := fmt.Sprintf("%s:%s@cloudsql(%s)/%s?parseTime=true",
		mysqlUser, mysqlUserPwd, mysqlHost, mysqlDatabase)

	// データベース接続
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("fail: db.Ping, %v\n", err)
	}
	defer db.Close()

	_db := sqlc.New(db)
	Dao := dao.NewDao(db, _db)
	Usecase := usecase.NewUsecase(Dao)
	Controller := controller.NewController(Usecase)

	r:= controller.SetupRoutes(Controller)

	log.Println("Listening...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}




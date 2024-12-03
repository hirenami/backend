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
	"os"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//envload()
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlUserPwd := os.Getenv("MYSQL_PASSWORD")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	mysqlHost := os.Getenv("MYSQL_HOST")
	
	if mysqlUser == "" || mysqlUserPwd == "" || mysqlDatabase == "" {
		log.Fatal("fail :Getenv")
	}

	// MySQL用のDSNを作成
	dsn := fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s?parseTime=true",
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
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
	"github.com/joho/godotenv"
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
    mysqlPort := os.Getenv("MYSQL_PORT")
	if mysqlUser == "" || mysqlUserPwd == "" || mysqlDatabase == "" || mysqlHost == "" || mysqlPort == "" {
		log.Fatal("fail :Getenv")
	}
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", 
        mysqlUser, mysqlUserPwd, mysqlHost, mysqlPort, mysqlDatabase))
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
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




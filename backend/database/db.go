package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"todo-with-gin/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".envファイルのローディングに失敗")
	}

	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	portStr := os.Getenv("PORT")
	sslmode := os.Getenv("SSLMODE")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal("データベース接続用のパスワード取得に失敗")
	}

	my_dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		host, user, password, dbname, port, sslmode)

	db, err := gorm.Open(postgres.Open(my_dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("データベース接続に失敗", err)
	}

	DB = db

	err = DB.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatal("マイグレーションに失敗", err)
	}

	log.Println("マイグレーションに成功")
}

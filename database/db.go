package database

import (
	"log"
	"todo-with-gin/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	my_dsn := "host=localhost user=KokiAoyagi password=kouki0802 dbname=todo port=5432 sslmode=disable"
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

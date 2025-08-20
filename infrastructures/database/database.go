package database

import (
	"fmt"
	"log"
	"os"
	"goTodoApp/infrastructures/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres" 
	"gorm.io/gorm"

)
//DB接続を開始する関数
func InitDB() (*gorm.DB, string) {
	// .env の読み込み（本番環境ではなくてもOKなように）
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, continuing with environment variables")
	}

	// 環境変数の取得
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	secretKey := os.Getenv("SECRET_KEY")

	// PostgreSQL 用 DSN
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	// DB接続
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// マイグレーション
	if err := db.AutoMigrate(&model.Todo{}, &model.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	return db, secretKey
}
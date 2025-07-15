package database

import (
	"fmt"
	"log"
	"os"
	"goTodoApp/infrastructures/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

)
//DB接続を開始する関数
func InitDB() (*gorm.DB, string) {
	// .env の読み込み（必要ならここで）
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// 環境変数の取得
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	secretKey := os.Getenv("SECRET_KEY")

	// DSN作成
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// DB接続
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// マイグレーション
	if err := db.AutoMigrate(&model.Todo{}, &model.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	return db, secretKey
}
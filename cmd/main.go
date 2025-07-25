package main

import (
    "time"
    "goTodoApp/di"
    "goTodoApp/infrastructures/database"
    "goTodoApp/routes"
    "goTodoApp/interface-adapter/middleware"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    // DBとsecretKeyの取得
    db, secretKey := database.InitDB()

    todoController, userController, tokenService:= di.InitControllers(db, secretKey)

    //Ginルータのセットアップ
    router := gin.Default()
    
    //corsの設定
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))

    //ミドルウェアの設定（トークン認証）
    authMiddleware := middleware.TokenAuthMiddleware(tokenService)

    // Todo のルートを登録
	routes.TodoRoutes(router, todoController, authMiddleware)
    routes.UserRoutes(router, userController)

    //サーバーの起動
    router.Run(":8080")
}
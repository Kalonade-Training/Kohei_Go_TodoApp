package main

import (
    "os"
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
    allowOrigins := []string{
        "http://localhost:5174",
        "http://localhost:5173",
        "https://react-todo-app-front.onrender.com"}

    router.Use(cors.New(cors.Config{
        AllowOrigins:     allowOrigins,
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    //ミドルウェアの設定（トークン認証）
    authMiddleware := middleware.TokenAuthMiddleware(tokenService)

    // Todo のルートを登録
	routes.TodoRoutes(router, todoController, authMiddleware)
    routes.UserRoutes(router, userController)

    // // React の index.html を使えるようにする
    // router.LoadHTMLFiles("./dist/index.html")

    // router.StaticFile("/vite.svg", "./dist/vite.svg")

    // router.NoRoute(func(c *gin.Context) {
    // c.HTML(200, "index.html", nil)
    // })

    // サーバーの起動ポートを環境変数から取得（Render用）
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    router.Run(":" + port)
}
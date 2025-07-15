package main

import (
    "goTodoApp/di"
    "goTodoApp/infrastructures/database"
    "goTodoApp/routes"
    "goTodoApp/interface-adapter/middleware"
    "github.com/gin-gonic/gin"
)

func main() {
    // DBとsecretKeyの取得
    db, secretKey := database.InitDB()

    todoController, userController, tokenService:= di.InitControllers(db, secretKey)

    //Ginルータのセットアップ
    router := gin.Default()

    //ミドルウェアの設定（トークン認証）
    authMiddleware := middleware.TokenAuthMiddleware(tokenService)

    // Todo のルートを登録
	routes.TodoRoutes(router, todoController, authMiddleware)
    routes.UserRoutes(router, userController)

    //サーバーの起動
    router.Run(":8080")
}
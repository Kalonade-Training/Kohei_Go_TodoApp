package routes

import (
	"github.com/gin-gonic/gin"
	"goTodoApp/controllers"
)

func TodoRoutes(r *gin.Engine, todoController *controllers.TodoController, authMiddleware gin.HandlerFunc) {
	authorized := r.Group("/", authMiddleware) // 認証ミドルウェア付きグループ

	authorized.POST("/todos", todoController.Create)
	authorized.GET("/todos", todoController.FindAll)
	authorized.GET("/todos/:id", todoController.FindByID)
	authorized.PUT("/todos/:id", todoController.Update)
	authorized.DELETE("/todos/:id", todoController.Delete)
	authorized.POST("/todos/:id/duplicate", todoController.Duplicate)
}
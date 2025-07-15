package routes

import (
	"github.com/gin-gonic/gin"
	"goTodoApp/interface-adapter/handlers"
)

func TodoRoutes(r *gin.Engine, todoController *handlers.TodoController, authMiddleware gin.HandlerFunc) {
	// 認証ミドルウェア付きグループ
	authorized := r.Group("/", authMiddleware) 

	authorized.POST("/todos", todoController.Create)
	authorized.GET("/todos", todoController.FindByUserIDWithFilters)
	authorized.GET("/todos/:id", todoController.FindTodoByID)
	authorized.PUT("/todos/:id", todoController.Update)
	authorized.DELETE("/todos/:id", todoController.Delete)
	authorized.POST("/todos/:id/duplicate", todoController.Duplicate)
}
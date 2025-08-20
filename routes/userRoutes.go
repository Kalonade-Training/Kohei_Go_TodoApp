package routes

import(
	"goTodoApp/interface-adapter/handlers"
	"github.com/gin-gonic/gin"
)
func UserRoutes(r *gin.Engine, userController *handlers.UserController){
	r.POST("/login", userController.Login)
	r.POST("/register", userController.Register)
}
package routes

import(
	"goTodoApp/controllers"
	"github.com/gin-gonic/gin"
)
func UserRoutes(r *gin.Engine, userController *controllers.UserController){
	r.POST("/login", userController.Login)
	r.POST("/register", userController.Register)
}
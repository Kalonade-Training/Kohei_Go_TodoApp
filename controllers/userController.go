package controllers
import(
	"net/http"

	"goTodoApp/usecases/user"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	LoginUser *user.LoginUser
	RegisterUser *user.RegisterUser
}
func NewUserController(
	loginUser *user.LoginUser,
	registerUser *user.RegisterUser,
) *UserController{
	return &UserController{
		LoginUser: loginUser,
		RegisterUser: registerUser,
	}
}

func (ctrl *UserController) Register(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	//リクエストをバインド
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	//ユーザーの登録処理を実行
	_, err := ctrl.RegisterUser.Execute(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user registered successfully"})
}

func (ctrl *UserController) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	//リクエストをバインド
	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	//ログイン処理を実行してトークンを生成
	token, err := ctrl.LoginUser.Execute(req.Username, req.Password)
	if err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token": token,
	})
}
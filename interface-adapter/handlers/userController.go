package handlers
import(
	"net/http"
	"fmt"
	"goTodoApp/application/user/user"
	"github.com/gin-gonic/gin"
	value_object "goTodoApp/domain/value-object"
	req "goTodoApp/interface-adapter/dto/request"
)

type UserController struct {
	LoginUserUC *user.LoginUserUseCase
	RegisterUserUC *user.RegisterUserUseCase
}
func NewUserController(
	loginUser *user.LoginUserUseCase,
	registerUser *user.RegisterUserUseCase,
) *UserController{
	return &UserController{
		LoginUserUC: loginUser,
		RegisterUserUC: registerUser,
	}
}

func (ctrl *UserController) Register(c *gin.Context) {
	var req req.User
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("[DEBUG] JSON Bind Error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	fmt.Printf("[DEBUG] Received request: %+v\n", req)

	username, err := value_object.FromStringUsername(req.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password, err := value_object.NewRawPassword(req.RawPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ctrl.RegisterUserUC.Execute(username, password)
	fmt.Printf("[DEBUG] UseCase Execute returned: token=%s, err=%v\n", token, err)
	if err != nil {
		//重複エラー
		if err.Error() == "username already exists"{
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "register successful",
		"token": token,
		"username": username.Value(),
	})
}

//ユーザー認証を行いTokenをreturn
func (ctrl *UserController) Login(c *gin.Context) {
	var req req.User
	//リクエストをバインド
	if err := c.ShouldBindJSON(&req); err != nil {
		//入力エラー
		c.JSON(http.StatusBadRequest, gin.H{"error": "username & password are required"})
		return
	}

	//usernameのバリデーション
	username, err := value_object.FromStringUsername(req.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})	
		return
	}

	//passwordのバリデーション
	password, err := value_object.NewRawPassword(req.RawPassword)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//ログイン処理を実行してトークンを生成
	token, err := ctrl.LoginUserUC.Execute(username, password)
	if err != nil{
		switch err.Error(){
		case "user not found":
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not registered"})
		case "invalid password":
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})	
		default:
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})		
		}
		return
	}
	
	//ログイン成功
	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token": token,
		"username": username.Value(),
	})
}
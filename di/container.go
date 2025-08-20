package di
import (
	"goTodoApp/infrastructures"
	"goTodoApp/interface-adapter/handlers"
	"goTodoApp/application/user/todo"
	"goTodoApp/application/user/user"
  "goTodoApp/domain/services"
  "gorm.io/gorm"
)

func InitControllers(db *gorm.DB, secretKey string) (
  todoController *handlers.TodoController,
	userController *handlers.UserController,
	tokenService services.ITokenService,
){
	
	//リポジトリの初期化
    todoRepo := infrastructures.NewGormTodoRepository(db)
    userRepo := infrastructures.NewGormUserRepository(db)
    hashService := infrastructures.NewBcryptService()
    tokenService = infrastructures.NewTokenService(secretKey)

		//ユースケースの初期化
    createTodoUC := todo.NewCreateTodoUseCase(todoRepo)
    findTodoByIDUC := todo.NewFindTodoByIDUseCase(todoRepo)
    FindByUserIDWithFiltersUC := todo.NewFindByUserIDWithFiltersUseCase(todoRepo)
    updateTodoUC := todo.NewUpdateTodoUseCase(todoRepo)
    deleteTodoUC := todo.NewDeleteTodoUseCase(todoRepo)
    duplicateTodoUC := todo.NewDuplicateTodoUseCase(todoRepo)

		//ユーザー関連のユースケース
    registerUserUC :=user.NewRegisterUserUseCase(userRepo,hashService,tokenService)
    loginUserUC := user.NewLoginUserUseCase(userRepo, hashService, tokenService)

		 //コントローラの初期化
    todoController = handlers.NewTodoController(
        createTodoUC,
        findTodoByIDUC,
        FindByUserIDWithFiltersUC,
        updateTodoUC,
        deleteTodoUC,
        duplicateTodoUC,
    )

		userController = handlers.NewUserController(
        loginUserUC,
        registerUserUC,
    )

		return todoController, userController, tokenService
}
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "goTodoApp/controllers"
    "goTodoApp/infrastructures"
    "goTodoApp/routes"
    "goTodoApp/usecases"
    "goTodoApp/usecases/user"
    "goTodoApp/middlewares"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "github.com/gin-gonic/gin"
    "goTodoApp/domain/entities"
)

func main() {
    //.envファイルの読み込み
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
    //環境変数の取得
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
    secretKey := os.Getenv("SECRET_KEY")
    //Mysqlの接続情報
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
    //GORMでMysqlに接続
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    //データベースのマイグレート
    err = db.AutoMigrate(&entities.Todo{}, &entities.User{})
    if err != nil {
    log.Fatal("Failed to migrate database:", err)
}
    //リポジトリの初期化
    todoRepo := infrastructures.NewGormTodoRepository(db)
    userRepo := infrastructures.NewGormUserRepository(db)
    hashService := infrastructures.NewBcryptService()
    tokenService := infrastructures.NewTokenService(secretKey)
    
    //ユースケースの初期化
    createTodoUC := usecases.NewCreateTodo(todoRepo)
    findTodoByIDUC := usecases.NewFindTodoByID(todoRepo)
    findAllTodosUC := usecases.NewFindAllTodos(todoRepo)
    updateTodoUC := usecases.NewUpdateTodo(todoRepo)
    deleteTodoUC := usecases.NewDeleteTodo(todoRepo)
    duplicateTodoUC := usecases.NewDuplicateTodo(todoRepo)
    //ユーザー関連のユースケース
    registerUserUC :=user.NewRegisterUser(userRepo,hashService,tokenService)
    loginUserUC := user.NewLoginUser(userRepo, hashService, tokenService)
    //コントローラの初期化
    todoController := controllers.NewTodoController(
        createTodoUC,
        findTodoByIDUC,
        findAllTodosUC,
        updateTodoUC,
        deleteTodoUC,
        duplicateTodoUC,
    )
    userController := controllers.NewUserController(
        loginUserUC,
        registerUserUC,
    )
    //Ginルータのセットアップ
    router := gin.Default()

    //ミドルウェアの設定（トークン認証）
    authMiddleware := middlewares.TokenAuthMiddleware(tokenService)

    // Todo のルートを登録
	routes.TodoRoutes(router, todoController, authMiddleware)
    routes.UserRoutes(router, userController)

    //サーバーの起動
    router.Run(":8080")
}
# GO Todo APP API
##　プロジェクトのディレクトリ
<pre>
  ├── application                                 #ルート
  │   └── user
  │       ├── todo                                #todo ユースケース
  │       │   ├── CreateTodoUsecase.go            #todo作成処理
  │       │   ├── DeleteTodoUsecase.go            #todo削除処理
  │       │   ├── DuplicateTodo.go                #todo複製処理
  │       │   ├── FindByUserIDWithFilters.go      #todo検索処理　（複数）
  │       │   ├── FindTodoByIDUsecase.go          #todo検索処理（単数）
  │       │   └── UpdateTodoUsecase.go            #todo更新処理
  │       └── user                                #user ユースケース
  │           ├── LoginUsecase.go                 #Login処理
  │           └── RegisterUsecase.go              #登録処理
  ├── cmd
  │   └── main.go                                 #エントリポイント
  ├── common                                      #共通の関数
  │   └── context.go                              #認証情報の取得
  ├── di                                          
  │   └── container.go                            #依存性の注入
  ├── docker-compose.yml                          #dockerの設定
  ├── Dockerfile                                  #dockerの設定
  ├── domain                                      #ドメイン層
  │   ├── entities
  │   │   ├── Todo.go                             #Todoエンティティ
  │   │   └── User.go                             #Userエンティティ
  │   ├── repositories
  │   │   ├── ITodoRepository.go                  #todoレポジトリインターフェース
  │   │   ├── IUserRepository.go                  #userレポジトリインターフェース
  │   │   └── TodoFilters.go                      #update用フィルター
  │   ├── services
  │   │   ├── IHashService.go                     #Hash化インターフェース
  │   │   └── ITokenService.go                    #token生成/認証のインターフェース
  │   └── value-object
  │       ├── completedAtVo.go                    ＃完了
  │       ├── descriptionVo.go                    ＃内容
  │       ├── dueDateVo.go                        #締め切り
  │       ├── hashedPasswordVo.go                 #パスワード(hashed)
  │       ├── rawPasswordVo.go                    #パスワード(raw)  
  │       ├── titleVo.go                          #見出し
  │       ├── todoIdVo.go                         #todo ID
  │       ├── userIdVo.go                         #user ID
  │       └── usernameVo.go                       #user name
  ├── go.mod
  ├── go.sum
  ├── infrastructures                             #インフラ層
  │   ├── database
  │   │   └── database.go                         #DB接続
  │   ├── GormTodoRepository.go                   #Todoレポジトリの実装
  │   ├── GormUserRepository.go                   #userレポジトリの実装
  │   ├── HashService.go                          #Hash関連レポジトリの実装
  │   ├── mapper
  │   │   ├── todoMapper.go                       #todo型変換
  │   │   └── userMapper.go                       #user型変換
  │   ├── model
  │   │   ├── todoModel.go                        #todo DBテーブルのモデル
  │   │   └── userModel.go                        #user DBテーブルのモデル
  │   └── TokenService.go                         #token関連レポジトリの実装
  ├── interface-adapter
  │   ├── dto                                     #データ型
  │   │   ├── request                             #リクエスト用
  │   │   │   ├── createTodo.go                   #todo作成
  │   │   │   ├── createUser.go                   #user作成
  │   │   │   ├── todoFilters.go                  #todo検索
  │   │   │   └── updateTodo.go                   #todo更新
  │   │   └── response                            #レスポンス用
  │   │       └── todo.go                         #todo戻り値
  │   ├── handlers                                #HTTPハンドラー
  │   │   ├── todoController.go                   #todoハンドラ
  │   │   └── userController.go                   #userハンドラ
  │   └── middleware                              #ミドルウェア
  │       └── auth_middleware.go                  #認証
  ├── README.md
  ├── routes                                      #ルート
  │   ├── todoRoute.go                            # todo ルーティング
  │   └── userRoutes.go                           # user ルーティング
  └── wait-for-it.sh                              # DB接続を待機
</pre>

package user

import (
	"errors"

	"goTodoApp/domain/entities"
	"goTodoApp/domain/repositories"
	"goTodoApp/domain/services"

	"github.com/google/uuid"
)

type RegisterUser struct {
	userRepo repositories.IUserRepository
	hashService services.IHashService
	tokenService services.ITokenService
}

//CreateUserインスタンスを生成
func NewRegisterUser(
	userRepo repositories.IUserRepository,
	hashService services.IHashService,
	tokenService services.ITokenService,
	) *RegisterUser{
		return &RegisterUser{
			userRepo: userRepo,
			hashService: hashService,
			tokenService: tokenService,
		}
	}
	//新しいユーザーの登録
	func (uc *RegisterUser) Execute(username, password string) (string, error){
		
		//ユーザーの重複チェック
		existingUser, err := uc.userRepo.FindByUsername(username)
		if err != nil && err.Error() != "record not found" {
    return "", err // 予期しないエラーの場合
		}
		if existingUser != nil {
    return "", errors.New("username already exists")
		}
		//パスワードのハッシュ化
		hashedPassword, err := uc.hashService.HashPassword(password)
		if err != nil {
			return "", err
		}
		// UUID生成
		id := uuid.New().String()

		//ユーザー作成
		newUser := &entities.User{
			ID: id,
			Username: username,
			Password: hashedPassword,
		}
		if err := uc.userRepo.Save(newUser); err != nil {
			return "", err
		}

		// JWTトークンの発行
		token, err := uc.tokenService.GenerateJWT(newUser.ID)
		if err != nil {
			return "", err
		}
		return token, nil
	}

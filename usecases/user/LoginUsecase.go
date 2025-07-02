package user

import (
	"errors"

	"goTodoApp/domain/repositories"
	"goTodoApp/domain/services"
)

type LoginUser struct {
	userRepo repositories.IUserRepository
	hashService services.IHashService
	tokenService services.ITokenService
}

//CreateUserインスタンスを生成
func NewLoginUser(
	userRepo repositories.IUserRepository,
	hashService services.IHashService,
	tokenService services.ITokenService,
	) *LoginUser{
		return &LoginUser{
			userRepo: userRepo,
			hashService: hashService,
			tokenService: tokenService,
		}
	}

	//ユーザーログインロジック
	func (uc *LoginUser) Execute(username, password string) (string, error){
		//ユーザーの検索
		user, err := uc.userRepo.FindByUsername(username)
		if err != nil {
			return "", errors.New("invalid username or password")
		}
		//パスワードの検証(true or falseが返ってくる)
		if !uc.hashService.CheckPasswordHash(password, user.Password){
			return "", errors.New("invalid username or password")
		}
		//JWTトークンの発行
		token, err := uc.tokenService.GenerateJWT(user.ID)
		if err != nil {
			return "", err
		}
		return token, nil
	}
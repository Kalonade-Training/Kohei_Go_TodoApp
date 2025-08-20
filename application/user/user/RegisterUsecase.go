package user

import (
	"errors"
	"time"
	"goTodoApp/domain/entities"
	"goTodoApp/domain/repositories"
	"goTodoApp/domain/services"
	value_object "goTodoApp/domain/value-object"
)

type RegisterUserUseCase struct {
	userRepo repositories.IUserRepository
	hashService services.IHashService
	tokenService services.ITokenService
}

//CreateUserインスタンスを生成
func NewRegisterUserUseCase(
	userRepo repositories.IUserRepository,
	hashService services.IHashService,
	tokenService services.ITokenService,
	) *RegisterUserUseCase{
		return &RegisterUserUseCase{
			userRepo: userRepo,
			hashService: hashService,
			tokenService: tokenService,
		}
	}
	//新しいユーザーの登録
	func (uc *RegisterUserUseCase) Execute(username value_object.Username, rawPassword value_object.RawPassword) (string, error){

		//ユーザーの重複チェック
		existingUser, err := uc.userRepo.FindByUsername(username.Value())
		if err != nil && err.Error() != "record not found" {
    return "", err // 予期しないエラーの場合
		}
		if existingUser != nil {
    return "", errors.New("username already exists")
		}

		//パスワードのハッシュ化
		hashedPassword, err := uc.hashService.HashPassword(rawPassword.Value())
		if err != nil {
			return "", err
		}
		hashedPasswordVo:= value_object.NewHashedPassword(hashedPassword)

		//ユーザー作成
		newUser := entities.NewUser(
			value_object.NewUserID(),
			username,
			hashedPasswordVo,
			time.Now(),
			time.Now(),
		)
	
		if _, err := uc.userRepo.Save(newUser); err != nil {
			return "", err
		}

		// JWTトークンの発行
		token, err := uc.tokenService.GenerateJWT(newUser.ID().Value())
		if err != nil {
			return "", err
		}
		return token, nil
	}

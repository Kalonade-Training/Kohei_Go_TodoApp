// usecases/user/login.go
package user

import (
	"errors"
	"goTodoApp/domain/repositories"
	"goTodoApp/domain/services"
	value_object "goTodoApp/domain/value-object"
)

type LoginUserUseCase struct {
	userRepo     repositories.IUserRepository
	hashService  services.IHashService
	tokenService services.ITokenService
}

func NewLoginUserUseCase(
	userRepo repositories.IUserRepository,
	hashService services.IHashService,
	tokenService services.ITokenService,
) *LoginUserUseCase {
	return &LoginUserUseCase{
		userRepo:     userRepo,
		hashService:  hashService,
		tokenService: tokenService,
	}
}

func (uc *LoginUserUseCase) Execute(username value_object.Username, rawPassword value_object.RawPassword) (string, error) {
	// ユーザーの検索
	user, err := uc.userRepo.FindByUsername(username.Value())
	if err != nil {
		return "", errors.New("User not found")
	}

	// パスワードチェック（ハッシュと比較）
	if !uc.hashService.CheckPasswordHash(rawPassword.Value(), user.HashedPassword().Value()) {
		return "", errors.New("invalid password")
	}

	// JWTトークン発行
	token, err := uc.tokenService.GenerateJWT(user.ID().Value())
	if err != nil {
		return "", err
	}

	return token, nil
}
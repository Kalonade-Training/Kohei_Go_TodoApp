package infrastructures

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"goTodoApp/domain/services"
)

type TokenService struct {
	SecretKey []byte
}

func NewTokenService(secretKey string) services.ITokenService {
	return &TokenService{SecretKey: []byte(secretKey)}
}
//JWTトークン発行
func (s *TokenService) GenerateJWT(userID string) (string, error){
	token :=jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})
	//トークン生成
	signedToken, err := token.SignedString(s.SecretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
//トークンの認証
func(s *TokenService) ValidateJWT(tokenString string) (string, error){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.SecretKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("invalid user ID in token")
	}
	return userID, nil
}
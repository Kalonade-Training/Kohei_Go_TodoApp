package infrastructures

import (
	"golang.org/x/crypto/bcrypt"
	"goTodoApp/domain/services"
)

type BcryptHashService struct{}

func NewBcryptService() services.IHashService{
	return &BcryptHashService{}
}

// HashPassword はパスワードをbcryptでハッシュ化する
func (s *BcryptHashService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash はパスワードとハッシュを比較する
func (s *BcryptHashService) CheckPasswordHash(rawPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))
	return err == nil
}

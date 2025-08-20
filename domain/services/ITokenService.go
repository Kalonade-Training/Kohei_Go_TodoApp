package services

type ITokenService interface {
	GenerateJWT(userID string) (string, error)
	ValidateJWT(token string) (string, error)
}
package services

type IHashService interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(rawPassword string, hashedPassword string) bool
}
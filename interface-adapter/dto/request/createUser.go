package dto

type User struct {
	Username    string `json:"username" binding:"required"`
	RawPassword string `json:"password" binding:"required"`
}
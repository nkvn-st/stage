package userservice

import (
	"stage/internal/messageservice"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Messages []messageservice.Message
}

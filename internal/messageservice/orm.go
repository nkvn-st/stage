package messageservice

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
	UserID uint
}

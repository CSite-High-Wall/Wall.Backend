package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserId          uuid.UUID `json:"user_id" gorm:"type:char(36);primary_key"`
	UserName        string    `json:"user_name"`
	Password        string    `json:"password"`
	TokenIdentifier uuid.UUID `json:"token_identifier" gorm:"type:char(36)"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.UserId = uuid.New()
	user.TokenIdentifier = uuid.Nil
	return
}

package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserId          uuid.UUID `json:"user_id" gorm:"type:char(36);primary_key"`
	UserName        string    `json:"user_name"`
	Password        string    `json:"password"`
	AvatarUrl       string    `json:"avatar_url"`
	TokenIdentifier uuid.UUID `json:"token_identifier" gorm:"type:char(36)"`
	CreatedAt       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	LastLoginTime   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"last_login_time"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.UserId = uuid.New()
	user.TokenIdentifier = uuid.Nil
	user.CreatedAt = time.Now()
	return
}

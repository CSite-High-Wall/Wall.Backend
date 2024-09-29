package model

import (
	"time"

	"github.com/google/uuid"
)

type AuthTokenResponseJsonObject struct {
	UserId      uuid.UUID `json:"user_id"`
	AccessToken string    `json:"access_token"`
	TokenType   string    `json:"token_type"`
	ExpiresIn   int       `json:"expires_in"`
}

type UserInfoResponseJsonObject struct {
	UserId        uuid.UUID `json:"user_id"`
	UserName      string    `json:"user_name"`
	AvatarUrl     string    `json:"avatar_url"`
	CreatedAt     time.Time `json:"created_at"`
	LastLoginTime time.Time `json:"last_login_time"`
}

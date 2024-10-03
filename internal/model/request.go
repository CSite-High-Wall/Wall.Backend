package model

import "github.com/google/uuid"

type RegisterRequestJsonObject struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LoginRequestJsonObject struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type AuthTokenRequestJsonObject struct {
	UserId      uuid.UUID `json:"user_id"`
	AccessToken string    `json:"access_token"`
}

type ExpressionCreateRequestJsonObject struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	Anonymity bool   `json:"anonymity"`
}

type ExpressionUpdateRequestJsonObject struct {
	ExpressionId uint64 `json:"expression_id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
}

type ExpressionAllGetRequestJsonObject struct {
	ExpressionId uint64 `json:"expression_id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
}

type ReviewCreateRequestJsonObject struct {
	ExpressionId uint64 `json:"expression_id"`
	Content      string `json:"content"`
}

type ReviewUpdateRequestJsonObject struct {
	ReviewId uint64 `json:"review_id"`
	Content  string `json:"content"`
}

type PasswordUpdateRequestJsonObject struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

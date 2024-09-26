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

type ReviewCreateRequestJsonObject struct {
	UserId       uuid.UUID `json:"user_id" `
	ExpressionId uint      `json:"expression_id"`
	Content      string    `json:"content"`
}

type ReviewUpdateRequestJsonObject struct {
	UserId  uuid.UUID `json:"user_id" `
	ID      uint      `json:"review_id"`
	Content string    `json:"content"`
}
type ReviewDeleteRequestJsonObject struct {
	UserId uuid.UUID `json:"user_id" `
	ID     uint      `json:"review_id"`
}

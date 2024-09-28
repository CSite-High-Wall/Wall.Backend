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
	ExpressionId uint   `json:"expression_id"`
	Content      string `json:"content"`
}

type ReviewDeleteRequestJsonObject struct {
	ID uint `json:"review_id"`
}

type ReviewUpdateRequestJsonObject struct {
	ID      uint   `json:"review_id"`
	Content string `json:"content"`
}

type ExpressionCreateRequestJsonObject struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	Anonymity bool   `json:"anonymity"`
}

type ExpressionUpdateRequestJsonObject struct {
	ExpressionId uint   `json:"expression_id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
}

type ExpressionDeleteRequestJsonObject struct {
	ExpressionId uint `json:"expression_id"`
}

package model

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserClaims struct {
	UserId          uuid.UUID `json:"user_id"`
	TokenIdentifier uuid.UUID `json:"token_identifier"`
	jwt.RegisteredClaims
}

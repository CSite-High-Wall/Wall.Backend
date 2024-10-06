package model

import (
	"github.com/google/uuid"
)

type BlacklistUserItem struct {
	OwnerUserId   uuid.UUID `json:"owner_user_id" gorm:"type:char(36)"`
	BlockedUserId uuid.UUID `json:"blocked_user_id" gorm:"type:char(36)"`
}

type BlacklistExpressionItem struct {
	OwnerUserId  uuid.UUID `json:"owner_user_id" gorm:"type:char(36)"`
	ExpressionId uint64    `json:"expression_id"`
}

package model

import (
	"github.com/google/uuid"
)

type BlacklistItem struct {
	OwnerUserId   uuid.UUID `json:"owner_user_id" gorm:"type:char(36)"`
	BlockedUserId uuid.UUID `json:"blocked_user_id" gorm:"type:char(36)"`
}

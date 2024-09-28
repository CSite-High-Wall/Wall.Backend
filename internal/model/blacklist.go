package model

import (
	"time"

	"github.com/google/uuid"
)

type Blacklist struct {
	ID      uint      `json:"blacklist_id"`
	BeingId uuid.UUID `json:"being_id" gorm:"type:char(36)"`
	PullId  uuid.UUID `json:"pull_id" gorm:"type:char(36)"`
	Time    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"time"`
}

package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	ReviewId     uint64         `json:"review_id"`
	UserId       uuid.UUID      `json:"user_id" gorm:"type:char(36)"`
	ExpressionId uint64         `json:"expression_id"`
	Content      string         `json:"content"`
	CreatedAt    time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (review *Review) BeforeCreate(tx *gorm.DB) (err error) {
	review.CreatedAt = time.Now()
	return
}

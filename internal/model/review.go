package model

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID           uint      `json:"review_id"`
	UserId       uuid.UUID `json:"user_id" gorm:"type:char(36)"`
	ExpressionId uint64    `json:"expression_id"`
	Content      string    `json:"content"`
	CreatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
	Time         time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"time"`
}

// func (review *Review) BeforeCreate(tx *gorm.DB) (err error) {
// 	review.UserId = uuid.New()
// 	return
// }

package model

import (
	"github.com/google/uuid"
	// "gorm.io/gorm"
)

type Review struct {
	ID      uint      `json:"review_id"`
	UserId  uuid.UUID `json:"user_id" gorm:"type:char(36)"`
	PostId  uint      `json:"post_id"`
	Content string    `json:"content"`
}

// func (review *Review) BeforeCreate(tx *gorm.DB) (err error) {
// 	review.UserId = uuid.New()
// 	return
// }

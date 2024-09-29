package model

//表白在数据库的结构
import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Expression struct {
	UserId       uuid.UUID      `gorm:"type:char(36);not null" json:"user_id"`
	ExpressionId uint64         `gorm:"primaryKey;uniqueIndex;autoIncrement" json:"expression_id"`
	Title        string         `gorm:"default:无标题" json:"title"`
	Content      string         `gorm:"not null" json:"content"`
	Anonymity    bool           `gorm:"not null" json:"anonymity"`
	CreatedAt    time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

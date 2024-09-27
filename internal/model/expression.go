package model

//表白在数据库的结构
import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

var Anonymity = 1   //匿名
var noAnonymity = 2 //不匿名

type Expression struct {
	ExpressionID uint           `gorm:"primaryKey" json:"expression_id"`
	Content      string         `gorm:"not null" json:"content"`
	UserID       uuid.UUID      `gorm:"type:bigint;not null;index" json:"user_id"`
	UserName     string         `json:"user_name"`
	CreatedAt    time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Anonymity    int            `gorm:"not null" json:"anonymity"`
}

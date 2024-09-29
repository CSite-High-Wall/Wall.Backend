package dao

import (
	"gorm.io/gorm"
	"wall-backend/internal/model"
)

type CommunityDao struct {
	db *gorm.DB
}

func NewCommunityDao(db *gorm.DB) CommunityDao {
	return CommunityDao{
		db: db,
	}
}

// 获取所有表白
func (dao CommunityDao) AllExpression() ([]model.Expression, error) {
	var expressions []model.Expression
	result := dao.db.Find(&expressions)
	return expressions, result.Error
}

// 获取特定表白
func (dao CommunityDao) GetExpressionById(expressionId uint) (model.Expression, error) {
	var expression model.Expression
	result := dao.db.First(&expression, expressionId)
	return expression, result.Error
}

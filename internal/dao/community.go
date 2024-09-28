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
func (dao CommunityDao) AllExpression() error {
	var expression model.Expression
	result := dao.db.Find(&expression)
	return result.Error
}

////获取特定表白
//func (dao CommunityDao) PartExpression() (error) {
//	var expression model.Expression
//	result := dao.db.Find(&expression)
//	return  result.Error

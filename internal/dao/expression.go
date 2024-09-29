package dao

//和表白的相关数据库操作

import (
	"time"
	"wall-backend/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ExpressionDao struct {
	db *gorm.DB
}

func NewExpressionDao(db *gorm.DB) ExpressionDao {
	return ExpressionDao{
		db: db,
	}
}

// 创建表白，anonymity 参数指示是否匿名
func (dao ExpressionDao) CreateExpression(userId uuid.UUID, title string, content string, anonymity bool) error {
	return dao.db.Create(&model.Expression{
		UserId:    userId,
		Title:     title,
		Content:   content,
		Anonymity: anonymity,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}).Error
}

// 更新表白内容
func (dao ExpressionDao) UpdateExpression(userId uuid.UUID, expressionId uint64, title string, content string) error {
	return dao.db.Model(&model.Expression{}).Where("user_id=? AND expression_id=?", userId, expressionId).Updates(map[string]interface{}{
		"title":      title,
		"content":    content,
		"updated_at": time.Now(),
	}).Error
}

// 删除表白
func (dao ExpressionDao) DeleteExpression(userId uuid.UUID, expressionId uint64) error {
	return dao.db.Where("user_id=? AND expression_id=?", userId, expressionId).Delete(&model.Expression{}).Error
}

// 根据 ExpressionId 查找表白
func (dao ExpressionDao) FindExpressionByExpressionId(expressionId uint64) (model.Expression, error) {
	var expression model.Expression
	result := dao.db.First(&expression, expressionId)
	return expression, result.Error
}

// 根据 UserId 查找表白
func (dao ExpressionDao) FindExpressionByUserId(userId uuid.UUID) ([]model.Expression, error) {
	var expressions []model.Expression
	result := dao.db.Find(&expressions, "user_id = ?", userId)
	return expressions, result.Error
}

// 获取所有已发布的表白
func (dao ExpressionDao) FetchAllExpression() ([]model.Expression, error) {
	var expressions []model.Expression
	result := dao.db.Find(&expressions)
	return expressions, result.Error
}

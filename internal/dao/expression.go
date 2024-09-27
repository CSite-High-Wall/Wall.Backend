package dao

//和表白的相关数据库操作

import (
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"wall-backend/internal/model"
)

type ExpressionDao struct {
	db *gorm.DB
}

func NewExpressionDao(db *gorm.DB) ExpressionDao {
	return ExpressionDao{
		db: db,
	}
}

//创建非匿名表白

func (dao ExpressionDao) CreateExpression1(ctx context.Context, UserId uuid.UUID, Content string, Anonymity int, UserName string) error {
	err := dao.db.WithContext(ctx).Create(&model.Expression{
		UserID:    UserId,
		Content:   Content,
		Anonymity: Anonymity,
		UserName:  UserName,
	}).Error
	return err
}

//创建匿名表白

func (dao ExpressionDao) CreateExpression2(ctx context.Context, UserId uuid.UUID, Content string, Anonymity int) error {
	err := dao.db.WithContext(ctx).Create(&model.Expression{
		UserID:    UserId,
		Content:   Content,
		UserName:  "匿名",
		Anonymity: Anonymity,
	}).Error
	return err
}

//更新表白内容

func (dao ExpressionDao) UpdateExpression(ctx context.Context, UserId uuid.UUID, ExpressionId uint, Content string) error {
	err := dao.db.WithContext(ctx).Model(&model.Expression{}).Where("user_id=? AND expression_id=?", UserId, ExpressionId).Updates(map[string]interface{}{
		"content": Content,
	}).Error
	return err
}

//删除表白

func (dao ExpressionDao) DeleteExpression(ctx context.Context, UserId uuid.UUID, ExpressionId uint) error {
	err := dao.db.WithContext(ctx).Where("user_id=? AND expression_id=?", UserId, ExpressionId).Delete(&model.Expression{}).Error
	return err
}

//查找用户是否存在

func (dao ExpressionDao) FindUserByUserId(ctx context.Context, userid uuid.UUID) (model.User, error) {
	var User model.User
	result := dao.db.WithContext(ctx).First(&User, userid)
	return User, result.Error
}

func (dao ExpressionDao) FindExpressionByExpressionId(ctx context.Context, expressionid uint) (model.Expression, error) {
	var Expression model.Expression
	result := dao.db.WithContext(ctx).First(&Expression, expressionid)
	return Expression, result.Error
}

package service

import (
	"wall-backend/internal/dao"
	"wall-backend/internal/model"

	"github.com/google/uuid"
)

type ExpressionService struct {
	expressionDao dao.ExpressionDao
}

func NewExpressionService(expressionDao dao.ExpressionDao) ExpressionService {
	return ExpressionService{
		expressionDao: expressionDao,
	}
}

var db dao.ExpressionDao

// anonymity 参数指示是否匿名
func (service ExpressionService) Publish(userId uuid.UUID, requestBody model.ExpressionCreateRequestJsonObject) error {
	return db.CreateExpression(ctx, userId, requestBody.Title, requestBody.Content, requestBody.Anonymity)
}

func (service ExpressionService) Edit(userId uuid.UUID, requestBody model.ExpressionUpdateRequestJsonObject) error {
	return db.UpdateExpression(ctx, userId, requestBody.ExpressionId, requestBody.Content, requestBody.Title)
}

func (service ExpressionService) Delete(userId uuid.UUID, requestBody model.ExpressionDeleteRequestJsonObject) error {
	return db.DeleteExpression(ctx, userId, requestBody.ExpressionId)
}

func (service ExpressionService) FindExpressionByExpressionId(expressionid uint) (model.Expression, error) {
	return db.FindExpressionByExpressionId(ctx, expressionid)
}

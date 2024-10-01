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

// anonymity 参数指示是否匿名
func (service ExpressionService) Publish(userId uuid.UUID, requestBody model.ExpressionCreateRequestJsonObject) error {
	return service.expressionDao.CreateExpression(userId, requestBody.Title, requestBody.Content, requestBody.Anonymity)
}

func (service ExpressionService) Edit(userId uuid.UUID, requestBody model.ExpressionUpdateRequestJsonObject) error {
	return service.expressionDao.UpdateExpression(userId, requestBody.ExpressionId, requestBody.Title, requestBody.Content)
}

func (service ExpressionService) Delete(userId uuid.UUID, expressionId uint64) error {
	return service.expressionDao.DeleteExpression(userId, expressionId)
}

func (service ExpressionService) FindExpressionByExpressionId(expressionid uint64) (model.Expression, error) {
	return service.expressionDao.FindExpressionByExpressionId(expressionid)
}

func (service ExpressionService) FetchAllExpression() ([]model.Expression, error) {
	return service.expressionDao.FetchAllExpression()
}

func (service ExpressionService) FetchUserExpression(userId uuid.UUID) ([]model.Expression, error) {
	return service.expressionDao.FindExpressionByUserId(userId)
}

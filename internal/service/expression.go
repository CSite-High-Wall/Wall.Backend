package service

import (
	"github.com/google/uuid"
	"wall-backend/internal/dao"
	"wall-backend/internal/model"
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

//非匿名

func (service ExpressionService) Publish1(requestBody model.ExpressionCreateRequestJsonObject) error {
	return db.CreateExpression1(ctx, requestBody.UserId, requestBody.Content, requestBody.Anonymity, requestBody.UserName)
}

//匿名

func (service ExpressionService) Publish2(requestBody model.ExpressionCreateRequestJsonObject) error {
	return db.CreateExpression2(ctx, requestBody.UserId, requestBody.Content, requestBody.Anonymity)
}

func (service ExpressionService) Delete(requestBody model.ExpressionDeleteRequestJsonObject) error {
	return db.DeleteExpression(ctx, requestBody.UserId, requestBody.ExpressionID)
}

func (service ExpressionService) Edit(requestBody model.ExpressionUpdateRequestJsonObject) error {
	return db.UpdateExpression(ctx, requestBody.UserId, requestBody.ExpressionID, requestBody.Content)
}

func (service ExpressionService) FindUserByUserId(userid uuid.UUID) (model.User, error) {
	return db.FindUserByUserId(ctx, userid)
}

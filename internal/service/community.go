package service

import (
	"wall-backend/internal/dao"
	"wall-backend/internal/model"
)

type CommunityService struct {
	communityDao dao.CommunityDao
}

func NewCommunityService(communityDao dao.CommunityDao) CommunityService {
	return CommunityService{
		communityDao: communityDao,
	}
}

func (service CommunityService) AllExpression() ([]model.Expression, error) {
	expressions, err := service.communityDao.AllExpression()
	if err != nil {
		return nil, err
	}
	return expressions, nil
}

func (service CommunityService) GetExpressionById(expressionId uint) (model.Expression, error) {
	expression, err := service.communityDao.GetExpressionById(expressionId)
	if err != nil {
		return model.Expression{}, err
	}
	return expression, nil
}

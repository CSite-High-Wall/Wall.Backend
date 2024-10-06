package service

import (
	"wall-backend/internal/dao"
	"wall-backend/internal/model"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

type BlacklistService struct {
	blacklistDao dao.BlacklistDao
}

func NewBlacklistService(dao dao.BlacklistDao) BlacklistService {
	return BlacklistService{
		blacklistDao: dao,
	}
}

func (service BlacklistService) AddBlacklistUserItem(ownerUserId uuid.UUID, blockedUserId uuid.UUID) error {
	return service.blacklistDao.CreateBlacklistUserItem(ownerUserId, blockedUserId)
}

func (service BlacklistService) RemoveBlacklistUserItem(ownerUserId uuid.UUID, blockedUserId uuid.UUID) error {
	return service.blacklistDao.DeleteBlacklistUserItem(ownerUserId, blockedUserId)
}

func (service BlacklistService) GetUserBlacklistByUserId(userId uuid.UUID) ([]model.BlacklistUserItem, error) {
	return service.blacklistDao.FindUserBlacklistByUserId(userId)
}

func (service BlacklistService) FindBlacklistUserItem(owner uuid.UUID, blocked uuid.UUID) (model.BlacklistUserItem, error) {
	return service.blacklistDao.FindBlacklistUserItem(owner, blocked)
}

func (service BlacklistService) FilterUserInBlacklist(userId uuid.UUID, array []gin.H) ([]gin.H, error) {
	blacklist, error := service.GetUserBlacklistByUserId(userId)
	var filteredList []gin.H

	if error != nil {
		return nil, error
	}

	for _, item := range array {
		var exist bool = false

		for _, blacklistItem := range blacklist {
			if blacklistItem.BlockedUserId == item["user_id"] {
				exist = true
				break
			}
		}

		if exist == true {
			continue
		}
		filteredList = append(filteredList, item)
	}

	return filteredList, nil
}

func (service BlacklistService) AddBlacklistExpressionItem(ownerUserId uuid.UUID, expressionId uint64) error {
	return service.blacklistDao.CreateBlacklistExpressionItem(ownerUserId, expressionId)
}

func (service BlacklistService) RemoveBlacklistExpressionItem(ownerUserId uuid.UUID, expressionId uint64) error {
	return service.blacklistDao.DeleteBlacklistExpressionItem(ownerUserId, expressionId)
}

func (service BlacklistService) GetExpressionBlacklistByUserId(ownerUserId uuid.UUID) ([]model.BlacklistExpressionItem, error) {
	return service.blacklistDao.FindExpressionBlacklistByUserId(ownerUserId)
}

func (service BlacklistService) FindBlacklistExpressionItem(ownerUserId uuid.UUID, expressionId uint64) (model.BlacklistExpressionItem, error) {
	return service.blacklistDao.FindBlacklistExpressionItem(ownerUserId, expressionId)
}

func (service BlacklistService) FilterExpressionInBlacklist(userId uuid.UUID, array []gin.H) ([]gin.H, error) {
	blacklist, error := service.GetExpressionBlacklistByUserId(userId)
	var filteredList []gin.H

	if error != nil {
		return nil, error
	}

	for _, item := range array {
		var exist bool = false

		for _, blacklistItem := range blacklist {
			if blacklistItem.ExpressionId == item["expression_id"] {
				exist = true
				break
			}
		}

		if exist == true {
			continue
		}
		filteredList = append(filteredList, item)
	}

	return filteredList, nil
}

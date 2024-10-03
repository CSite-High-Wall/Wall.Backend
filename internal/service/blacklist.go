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

func (service BlacklistService) Add(ownerUserId uuid.UUID, blockedUserId uuid.UUID) error {
	return service.blacklistDao.CreateBlacklistItem(ownerUserId, blockedUserId)
}

func (service BlacklistService) Remove(ownerUserId uuid.UUID, blockedUserId uuid.UUID) error {
	return service.blacklistDao.DeleteBlacklistItem(ownerUserId, blockedUserId)
}

func (service BlacklistService) FindBlacklistItemsByUserId(userId uuid.UUID) ([]model.BlacklistItem, error) {
	return service.blacklistDao.FindBlacklistItemsByUserId(userId)
}

func (service BlacklistService) FindBlacklistItem(owner uuid.UUID, blocked uuid.UUID) (model.BlacklistItem, error) {
	return service.blacklistDao.FindBlacklistItem(owner, blocked)
}

func (service BlacklistService) FilterUserInBlacklist(userId uuid.UUID, array []gin.H) ([]gin.H, error) {
	blacklist, error := service.FindBlacklistItemsByUserId(userId)
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

func (service BlacklistService) AddByExpression(ownerUserId uuid.UUID, expressionId uint64) error {
	return service.blacklistDao.CreateBlacklistExpression(ownerUserId, expressionId)
}

func (service BlacklistService) RemoveByExpression(ownerUserId uuid.UUID, expressionId uint64) error {
	return service.blacklistDao.DeleteBlacklistExpression(ownerUserId, expressionId)
}

func (service BlacklistService) FindBlacklistExpressionByUserId(ownerUserId uuid.UUID) ([]model.BlacklistExpression, error) {
	return service.blacklistDao.FindBlacklistExpressionByUserId(ownerUserId)
}

func (service BlacklistService) FindBlacklistExpression(ownerUserId uuid.UUID, expressionId uint64) (model.BlacklistExpression, error) {
	return service.blacklistDao.FindBlacklistExpression(ownerUserId, expressionId)
}

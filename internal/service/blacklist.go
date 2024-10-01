package service

import (
	"github.com/gin-gonic/gin"
	"wall-backend/internal/dao"
	"wall-backend/internal/model"

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

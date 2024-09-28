package service

import (
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

func (service BlacklistService) Add(userId uuid.UUID, requestBody model.BlacklistCreateRequestJsonObject) error {
	return service.blacklistDao.CreateBlacklist(userId, requestBody.BeingId)
}

func (service BlacklistService) Remove(userId uuid.UUID, requestBody model.BlacklistDeleteRequestJsonObject) error {
	return service.blacklistDao.DeleteBlacklist(userId, requestBody.ID)
}

func (service BlacklistService) AllBlacklist() ([]model.Blacklist, error) {
	blacklists, err := service.blacklistDao.AllBlacklist()
	if err != nil {
		return nil, err
	}
	return blacklists, nil
}

func (service BlacklistService) FindBlacklistById(id uint) (model.Blacklist, error) {
	return service.blacklistDao.FindBlacklistById(id)
}

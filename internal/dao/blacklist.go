package dao

import (
	"time"
	"wall-backend/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlacklistDao struct {
	db *gorm.DB
}

func NewBlacklistDao(db *gorm.DB) BlacklistDao {
	return BlacklistDao{
		db: db,
	}
}

func (a BlacklistDao) CreateBlacklist(UserId uuid.UUID, BeingId uuid.UUID) error {
	err := a.db.Create(&model.Blacklist{
		BeingId: BeingId,
		PullId:  UserId,
		Time:    time.Now(),
	}).Error
	return err

}

func (a BlacklistDao) DeleteBlacklist(UserId uuid.UUID, ID uint) error {
	err := a.db.Where("blacklist_id=? AND pull_id=?", ID, UserId).Delete(&model.Blacklist{}).Error
	return err
}

func (a BlacklistDao) AllBlacklist() ([]model.Blacklist, error) {
	var blacklists []model.Blacklist
	err := a.db.Find(&blacklists).Error
	return blacklists, err
}

func (a BlacklistDao) FindBlacklistById(id uint) (model.Blacklist, error) {
	var blacklist model.Blacklist
	result := a.db.Where("blacklist_id=?", id).First(&blacklist)

	return blacklist, result.Error
}

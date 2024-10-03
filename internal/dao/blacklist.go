package dao

import (
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

// 根据屏蔽关系创建一个黑名单项
func (dao BlacklistDao) CreateBlacklistItem(ownerUserId uuid.UUID, blockedUserId uuid.UUID) error {
	return dao.db.Create(&model.BlacklistItem{
		OwnerUserId:   ownerUserId,
		BlockedUserId: blockedUserId,
	}).Error
}

// 根据屏蔽关系删除一个黑名单项
func (dao BlacklistDao) DeleteBlacklistItem(ownerUserId uuid.UUID, blockedUserId uuid.UUID) error {
	err := dao.db.Model(&model.BlacklistItem{}).Where("owner_user_id = ? AND blocked_user_id = ?", ownerUserId, blockedUserId).Delete(&model.BlacklistItem{}).Error
	return err
}

// 根据屏蔽关系查找一个黑名单项
func (dao BlacklistDao) FindBlacklistItem(ownerUserId uuid.UUID, blockedUserId uuid.UUID) (model.BlacklistItem, error) {
	var blacklistItem model.BlacklistItem
	result := dao.db.First(&blacklistItem, "owner_user_id = ? AND blocked_user_id = ?", ownerUserId, blockedUserId)

	return blacklistItem, result.Error
}

// 更具用户 Id 查找其黑名单
func (dao BlacklistDao) FindBlacklistItemsByUserId(ownerUserId uuid.UUID) ([]model.BlacklistItem, error) {
	var blacklists []model.BlacklistItem
	err := dao.db.Find(&blacklists, "owner_user_id = ?", ownerUserId).Error
	return blacklists, err
}

func (dao BlacklistDao) CreateBlacklistExpression(ownerUserId uuid.UUID, expressionId uint64) error {
	return dao.db.Create(&model.BlacklistExpression{
		OwnerUserId:  ownerUserId,
		ExpressionId: expressionId,
	}).Error
}

func (dao BlacklistDao) DeleteBlacklistExpression(ownerUserId uuid.UUID, expressionId uint64) error {
	err := dao.db.Model(&model.BlacklistExpression{}).Where("owner_user_id = ? AND expression_id = ?", ownerUserId, expressionId).Delete(&model.BlacklistExpression{}).Error
	return err
}

func (dao BlacklistDao) FindBlacklistExpression(ownerUserId uuid.UUID, expressionId uint64) (model.BlacklistExpression, error) {
	var blacklistExpression model.BlacklistExpression
	result := dao.db.First(&blacklistExpression, "owner_user_id = ? AND expression_id = ?", ownerUserId, expressionId)

	return blacklistExpression, result.Error
}

// 更具用户 Id 查找其黑名单
func (dao BlacklistDao) FindBlacklistExpressionByUserId(ownerUserId uuid.UUID) ([]model.BlacklistExpression, error) {
	var blacklists []model.BlacklistExpression
	err := dao.db.Find(&blacklists, "owner_user_id = ?", ownerUserId).Error
	return blacklists, err
}

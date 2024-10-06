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
func (dao BlacklistDao) CreateBlacklistUserItem(ownerUserId uuid.UUID, blockedUserId uuid.UUID) error {
	return dao.db.Create(&model.BlacklistUserItem{
		OwnerUserId:   ownerUserId,
		BlockedUserId: blockedUserId,
	}).Error
}

// 根据屏蔽关系删除一个黑名单项
func (dao BlacklistDao) DeleteBlacklistUserItem(ownerUserId uuid.UUID, blockedUserId uuid.UUID) error {
	err := dao.db.Model(&model.BlacklistUserItem{}).Where("owner_user_id = ? AND blocked_user_id = ?", ownerUserId, blockedUserId).Delete(&model.BlacklistUserItem{}).Error
	return err
}

// 根据屏蔽关系查找一个黑名单项
func (dao BlacklistDao) FindBlacklistUserItem(ownerUserId uuid.UUID, blockedUserId uuid.UUID) (model.BlacklistUserItem, error) {
	var blacklistItem model.BlacklistUserItem
	result := dao.db.First(&blacklistItem, "owner_user_id = ? AND blocked_user_id = ?", ownerUserId, blockedUserId)

	return blacklistItem, result.Error
}

// 更具用户 Id 查找其黑名单
func (dao BlacklistDao) FindUserBlacklistByUserId(ownerUserId uuid.UUID) ([]model.BlacklistUserItem, error) {
	var blacklists []model.BlacklistUserItem
	err := dao.db.Find(&blacklists, "owner_user_id = ?", ownerUserId).Error
	return blacklists, err
}

func (dao BlacklistDao) CreateBlacklistExpressionItem(ownerUserId uuid.UUID, expressionId uint64) error {
	return dao.db.Create(&model.BlacklistExpressionItem{
		OwnerUserId:  ownerUserId,
		ExpressionId: expressionId,
	}).Error
}

func (dao BlacklistDao) DeleteBlacklistExpressionItem(ownerUserId uuid.UUID, expressionId uint64) error {
	err := dao.db.Model(&model.BlacklistExpressionItem{}).Where("owner_user_id = ? AND expression_id = ?", ownerUserId, expressionId).Delete(&model.BlacklistExpressionItem{}).Error
	return err
}

func (dao BlacklistDao) FindBlacklistExpressionItem(ownerUserId uuid.UUID, expressionId uint64) (model.BlacklistExpressionItem, error) {
	var blacklistExpression model.BlacklistExpressionItem
	result := dao.db.First(&blacklistExpression, "owner_user_id = ? AND expression_id = ?", ownerUserId, expressionId)

	return blacklistExpression, result.Error
}

// 更具用户 Id 查找其黑名单
func (dao BlacklistDao) FindExpressionBlacklistByUserId(ownerUserId uuid.UUID) ([]model.BlacklistExpressionItem, error) {
	var blacklists []model.BlacklistExpressionItem
	err := dao.db.Find(&blacklists, "owner_user_id = ?", ownerUserId).Error
	return blacklists, err
}

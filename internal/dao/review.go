package dao

import (
	"context"
	"wall-backend/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReviewDao struct {
	db *gorm.DB
}

func NewReviewDao(db *gorm.DB) ReviewDao {
	return ReviewDao{
		db: db,
	}
}

func (dao ReviewDao) CreateReview(ctx context.Context, UserId uuid.UUID, PostId uint, Content string) error {
	err := dao.db.WithContext(ctx).Create(&model.Review{
		UserId:  UserId,
		PostId:  PostId,
		Content: Content,
	}).Error
	return err
}

func (dao ReviewDao) UpdateReview(ctx context.Context, UserId uuid.UUID, ReviewId uint, Content string) error {
	err := dao.db.WithContext(ctx).Model(&model.Review{}).Where("user_id=? AND review_id=?", UserId, ReviewId).Updates(map[string]interface{}{
		"content": Content,
	}).Error
	return err
}

func (dao ReviewDao) DeleteReview(ctx context.Context, UserId uuid.UUID, ReviewId uint) error {
	err := dao.db.WithContext(ctx).Where("user_id=? AND review_id=?", UserId, ReviewId).Delete(&model.Review{}).Error
	return err
}

// func (dao ReviewDao)GetContactByOwnerID(ctx context.Context,uid uint)([]model.Contact,error){
// 	var contacts []model.Contact
// 	err:=dao.db.WithContext(ctx).Where("owner_id=?",uid).Find(&contacts).Error
// 	return contacts,err
// }

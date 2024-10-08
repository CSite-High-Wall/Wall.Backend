package dao

import (
	"time"
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

// 创建评论
func (dao ReviewDao) CreateReview(UserId uuid.UUID, ExpressionId uint64, Content string) error {
	err := dao.db.Create(&model.Review{
		UserId:       UserId,
		ExpressionId: ExpressionId,
		Content:      Content,
	}).Error
	return err
}

// 删除评论
func (dao ReviewDao) DeleteReview(userId uuid.UUID, reviewId uint64) error {
	err := dao.db.Model(&model.Review{}).Where("user_id=? AND review_id=?", userId, reviewId).Update("DeletedAt", time.Now()).Error
	return err
}

// 更新评论内容
func (dao ReviewDao) UpdateReview(userId uuid.UUID, reviewId uint64, content string) error {
	err := dao.db.Model(&model.Review{}).Where("user_id=? AND review_id=?", userId, reviewId).Updates(map[string]interface{}{
		"content": content,
		"time":    time.Now(),
	}).Error
	return err
}

// 根据评论 Id 查找评论
func (dao ReviewDao) FindReviewByReviewId(reviewId uint64) (model.Review, error) {
	var review model.Review
	result := dao.db.First(&review, "review_id = ?", reviewId)

	return review, result.Error
}

// 根据表白 Id 查找评论
func (dao ReviewDao) FindReviewByExpressionId(expressionId uint64) ([]model.Review, error) {
	var reviews []model.Review
	result := dao.db.Find(&reviews, "expression_id = ?", expressionId)

	return reviews, result.Error
}

// // 根据id和评论，看是否是本人操作
// func (dao ReviewDao) FindReviewByUserId(ctx context.Context,ReviewId uint) (model.Review, error) {
// 	var review model.Review
// 	result := dao.db.First(&review, "user_id = ? AND review_id = ?", UserId, ReviewId)

// 	return review, result.Error
// }

// // SoftDeleteReport 软删除一个帖子
// func SoftDeleteReport(user_id uint, post_id uint) (*models.Post, error) {
// 	// 查询要删除的帖子
// 	var post models.Post
// 	result := database.DB.Where("id = ? AND post_id = ?", user_id, post_id).First(&post)
// 	if result.Error != nil {
// 		return nil, result.Error // 如果找不到记录，返回错误
// 	}

// 	// 更新 DeletedAt 字段来软删除帖子
// 	result = database.DB.Model(&post).Update("DeletedAt", time.Now())
// 	if result.Error != nil {
// 		return nil, result.Error // 更新失败，返回错误
// 	}

// 	return &post, nil // 返回更新后的帖子
// }

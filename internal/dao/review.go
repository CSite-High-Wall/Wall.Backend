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
		Time:         time.Now(),
	}).Error
	return err
}

// 删除评论
func (dao ReviewDao) DeleteReview(UserId uuid.UUID, ReviewId uint) error {
	err := dao.db.Where("user_id=? AND review_id=?", UserId, ReviewId).Update("DeletedAt", time.Now()).Error
	return err
}

//更新评论内容

func (dao ReviewDao) UpdateReview(UserId uuid.UUID, ReviewId uint, Content string) error {
	err := dao.db.Model(&model.Review{}).Where("user_id=? AND review_id=?", UserId, ReviewId).Updates(map[string]interface{}{
		"content": Content,
		"time":    time.Now(),
	}).Error
	return err
}

// 根据表白ID查询表白，看表白是否存在
func (dao ReviewDao) FindPostByPostId(ExpressionId uint) (model.Review, error) {
	var review model.Review
	result := dao.db.First(&review, "expression_id = ?", ExpressionId)

	return review, result.Error
}

// 根据评论ID查询评论，看评论是否存在
func (dao ReviewDao) FindReviewByReviewId(ReviewId uint) (model.Review, error) {
	var review model.Review
	result := dao.db.First(&review, "review_id = ?", ReviewId)

	return review, result.Error
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

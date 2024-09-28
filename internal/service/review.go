package service

import (
	"wall-backend/internal/dao"
	"wall-backend/internal/model"

	"github.com/google/uuid"
)

type ReviewService struct {
	reviewDao dao.ReviewDao
}

func NewReviewService(dao dao.ReviewDao) ReviewService {
	return ReviewService{
		reviewDao: dao,
	}
}

func (service ReviewService) Publish(userId uuid.UUID, requestBody model.ReviewCreateRequestJsonObject) error {
	return service.reviewDao.CreateReview(userId, requestBody.ExpressionId, requestBody.Content)
}

func (service ReviewService) Delete(userId uuid.UUID, requestBody model.ReviewDeleteRequestJsonObject) error {
	return service.reviewDao.DeleteReview(userId, requestBody.ID)
}

func (service ReviewService) Edit(userId uuid.UUID, requestBody model.ReviewUpdateRequestJsonObject) error {
	return service.reviewDao.UpdateReview(userId, requestBody.ID, requestBody.Content)
}

func (service ReviewService) FindReviewByReviewId(id uint) (model.Review, error) {
	return service.reviewDao.FindReviewByReviewId(id)
}

// func (service ReviewService) FindReviewByUserId(userId uuid.UUID, id uint) (model.Review, error) {
// 	return d.FindReviewByUserId(ctx, userId, id)
// }

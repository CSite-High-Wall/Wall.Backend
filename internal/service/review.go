package service

import (
	"context"
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

var ctx context.Context
var d dao.ReviewDao

func (service ReviewService) Publish(requestBody model.ReviewCreateRequestJsonObject) error {
	return d.CreateReview(ctx, requestBody.UserId, requestBody.ExpressionId, requestBody.Content)
}

func (service ReviewService) Delete(requestBody model.ReviewDeleteRequestJsonObject) error {
	return d.DeleteReview(ctx, requestBody.UserId, requestBody.ID)
}

func (service ReviewService) Edit(requestBody model.ReviewUpdateRequestJsonObject) error {
	return d.UpdateReview(ctx, requestBody.UserId, requestBody.ID, requestBody.Content)
}

func (service ReviewService) FindPostByPostId(expressionId uint) (model.Review, error) {
	return d.FindPostByPostId(ctx, expressionId)
}

func (service ReviewService) FindReviewByReviewId(id uint) (model.Review, error) {
	return d.FindReviewByReviewId(ctx, id)
}
func (service ReviewService) FindReviewByUserId(userId uuid.UUID, id uint) (model.Review, error) {
	return d.FindReviewByUserId(ctx, userId, id)
}

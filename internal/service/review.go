package service

import (
	"context"
	"wall-backend/internal/dao"
	"wall-backend/internal/model"

	"github.com/google/uuid"
)

var ctx context.Context
var d dao.ReviewDao

func CreateReview(UserId uuid.UUID, PostId uint, Content string) error {
	re := model.Review{
		UserId:  UserId,
		PostId:  PostId,
		Content: Content,
	}
	return d.CreateReview(ctx, re.UserId, re.PostId, re.Content)

}

func UpdateReview(UserId uuid.UUID, ReviewId uint, Content string) error {
	return d.UpdateReview(ctx, UserId, ReviewId, Content)
}

func DeleteReview(UserId uuid.UUID, ReviewId uint) error {
	return d.DeleteReview(ctx, UserId, ReviewId)
}

// func GetContactList(uid uint)([]model.Contact,error){
// 	return d.GetContactByOwnerID(ctx,uid)
// }

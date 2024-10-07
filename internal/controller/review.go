package controller

import (
	"errors"
	"wall-backend/internal/model"
	"wall-backend/internal/service"
	"wall-backend/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReviewController struct {
	userService       service.UserService
	expressionService service.ExpressionService
	reviewService     service.ReviewService
}

func NewReviewController(userService service.UserService, reviewService service.ReviewService, expressionService service.ExpressionService) ReviewController {
	return ReviewController{
		reviewService:     reviewService,
		expressionService: expressionService,
		userService:       userService,
	}
}

// 发布评论接口
func (controller ReviewController) Publish(c *gin.Context) {
	var requestBody model.ReviewCreateRequestJsonObject
	var userId = utils.ParseUserIdFromRequest(c)
	if err := c.BindJSON(&requestBody); err != nil {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	}
	_, error := controller.userService.FindUserByUserId(userId)
	if errors.Is(error, gorm.ErrRecordNotFound) {
		utils.ResponseFailWithoutData(c, "未找到该用户") // 检查用户
	} else if error != nil {
		utils.ResponseFailWithoutData(c, "获取用户信息失败") // 检查用户
	} else {
		_, error := controller.expressionService.FindExpressionByExpressionId(requestBody.ExpressionId)
		if errors.Is(error, gorm.ErrRecordNotFound) {
			utils.ResponseFailWithoutData(c, "未找到该表白帖子") // 检查表白帖子
		} else if error != nil {
			utils.ResponseFailWithoutData(c, "获取表白帖子信息失败") // 检查表白帖子
		} else if error := controller.reviewService.Publish(userId, requestBody); error != nil {
			utils.ResponseFailWithoutData(c, "发布评论失败")
		} else {
			utils.ResponseOkWithoutData(c)
		}
	}
}

// 删除评论接口
func (controller ReviewController) Delete(c *gin.Context) {
	var userId = utils.ParseUserIdFromRequest(c)
	exist, reviewId := utils.TryGetUInt64(c, "review_id")

	if !exist {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	}

	_, error := controller.userService.FindUserByUserId(userId)
	if errors.Is(error, gorm.ErrRecordNotFound) {
		utils.ResponseFailWithoutData(c, "未找到该用户") // 检查用户
	} else if error != nil {
		utils.ResponseFailWithoutData(c, "获取用户信息失败") // 检查用户
	} else {
		review, error := controller.reviewService.FindReviewByReviewId(reviewId)
		if error != nil {
			utils.ResponseFailWithoutData(c, "获取评论信息失败")
		} else if review.UserId != userId {
			utils.ResponseFailWithoutData(c, "您只能删除自己的评论") // 判断创建评论的用户和请求者是否为同一人
		} else if error = controller.reviewService.Delete(userId, reviewId); error != nil {
			utils.ResponseFailWithoutData(c, "删除表白失败") // 删除评论失败
		} else {
			utils.ResponseOkWithoutData(c) // 返回成功相应
		}
	}

}

// 更新评论接口
func (controller ReviewController) Edit(c *gin.Context) {
	var requestBody model.ReviewUpdateRequestJsonObject
	var userId = utils.ParseUserIdFromRequest(c)
	if err := c.BindJSON(&requestBody); err != nil {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	}

	_, error := controller.userService.FindUserByUserId(userId)
	if errors.Is(error, gorm.ErrRecordNotFound) {
		utils.ResponseFailWithoutData(c, "未找到该用户") // 检查用户
	} else if error != nil {
		utils.ResponseFailWithoutData(c, "获取用户信息失败") // 检查用户
	} else {
		review, error := controller.reviewService.FindReviewByReviewId(requestBody.ReviewId)
		if error != nil {
			utils.ResponseFailWithoutData(c, "获取评论信息失败")
		} else if review.UserId != userId {
			utils.ResponseFailWithoutData(c, "您只能修改自己的评论") // 判断创建评论的用户和请求者是否为同一人
		} else if error = controller.reviewService.Edit(userId, requestBody); error != nil {
			utils.ResponseFailWithoutData(c, "修改表白失败") // 修改评论失败
		} else {
			utils.ResponseOkWithoutData(c) // 返回成功相应
		}
	}
}

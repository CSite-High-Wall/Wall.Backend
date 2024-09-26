package controller

import (
	"wall-backend/internal/model"
	"wall-backend/internal/service"
	"wall-backend/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReviewController struct {
	reviewService service.ReviewService
}

func NewReviewController(service service.ReviewService) ReviewController {
	return ReviewController{
		reviewService: service,
	}
}

func (controller ReviewController) Publish(c *gin.Context) {
	var requestBody model.ReviewCreateRequestJsonObject
	if err := c.BindJSON(&requestBody); err != nil {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	}
	_, err := controller.reviewService.FindPostByPostId(requestBody.ExpressionId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ResponseFailWithoutData(c, "帖子不存在")
			return
		} else {
			utils.ResponseFailWithoutData(c, "评论失败")
			return
		}
	}
	utils.ResponseOkWithoutData(c)
}

func (controller ReviewController) Delete(c *gin.Context) {
	var requestBody model.ReviewDeleteRequestJsonObject
	if err := c.BindJSON(&requestBody); err != nil {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	}
	_, err := controller.reviewService.FindReviewByReviewId(requestBody.ID)
	if err != nil {
		utils.ResponseFailWithoutData(c, "评论不存在")
		return
	}
	_, err = controller.reviewService.FindReviewByUserId(requestBody.UserId, requestBody.ID)
	if err != nil {
		utils.ResponseFailWithoutData(c, "不是本人操作")
		return
	}
	utils.ResponseOkWithoutData(c)

}

func (controller ReviewController) Edit(c *gin.Context) {
	var requestBody model.ReviewUpdateRequestJsonObject
	if err := c.BindJSON(&requestBody); err != nil {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	}
	_, err := controller.reviewService.FindReviewByReviewId(requestBody.ID)
	if err != nil {
		utils.ResponseFailWithoutData(c, "评论不存在")
		return
	}
	utils.ResponseOkWithoutData(c)

}

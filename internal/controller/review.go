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
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	}

}

func (controller ReviewController) Edit(c *gin.Context) {
	var requestBody model.ReviewUpdateRequestJsonObject
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	}

}

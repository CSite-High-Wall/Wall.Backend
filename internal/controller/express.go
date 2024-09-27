package controller

import (
	"github.com/gin-gonic/gin"
	"wall-backend/internal/model"
	"wall-backend/internal/service"
	"wall-backend/pkg/utils"
)

type ExpressController struct {
	expressionService service.ExpressionService
}

func (controller ExpressController) Publish(c *gin.Context) {
	var requestBody model.ExpressionCreateRequestJsonObject

	if err := c.BindJSON(&requestBody); err != nil {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	}

	//判断是否匿名



	// 创建新帖子

	// 将新帖子保存到数据库

	// 返回成功响应
}

func (controller ExpressController) Delete(c *gin.Context) {

}

func (controller ExpressController) Edit(c *gin.Context) {

}

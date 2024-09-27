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

func NewExpressController(service service.ExpressionService) ExpressController {
	return ExpressController{
		expressionService: service,
	}
}

func (a ExpressController) Publish(c *gin.Context) {
	var requestBody model.ExpressionCreateRequestJsonObject

	if err := c.BindJSON(&requestBody); err != nil {
		utils.ResponseFail(c, "无效的请求参数", err)
		return
	}
	//判断是否匿名
	user, err := a.expressionService.FindUserByUserId(requestBody.UserId)
	
	// 创建新帖子

	// 将新帖子保存到数据库

	// 返回成功响应
}

func (a ExpressController) Delete(c *gin.Context) {

}

func (a ExpressController) Edit(c *gin.Context) {

}

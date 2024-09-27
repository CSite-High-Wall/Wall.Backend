package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	_, err := a.expressionService.FindUserByUserId(requestBody.UserId)

	//查找用户
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ResponseFailWithoutData(c, "该用户未授权")
		return
	}
	if err != nil {
		utils.ResponseFail(c, "获取用户信息失败", nil)
		return
	}

	//匿名
	if requestBody.Anonymity == model.Anonymity {
		err = a.expressionService.Publish2(requestBody)
		if err != nil {
			utils.ResponseFail(c, "保存表白到数据库失败", nil)
			return
		}
	}

	//不匿名
	if requestBody.Anonymity == model.NoAnonymity {
		err = a.expressionService.Publish1(requestBody)
		if err != nil {
			utils.ResponseFail(c, "保存表白到数据库失败", nil)
			return
		}
	}

	// 返回成功响应
	utils.ResponseOkWithoutData(c)
}

//删除表白

func (a ExpressController) Delete(c *gin.Context) {
	var requestBody model.ExpressionDeleteRequestJsonObject

	if err := c.BindJSON(&requestBody); err != nil {
		utils.ResponseFail(c, "无效的请求参数", err)
		return
	}
	user, err := a.expressionService.FindUserByUserId(requestBody.UserId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ResponseFailWithoutData(c, "该用户未授权")
		return
	}
	if err != nil {
		utils.ResponseFail(c, "获取用户信息失败", nil)
		return
	}

	var expression model.Expression
	expression, err = a.expressionService.FindExpressionByExpressionId(expression.ExpressionID)
	if err != nil {
		utils.ResponseFail(c, "获取表白信息失败", nil)
		return
	}

	//判断创建表白的用户和请求者是否为同一人
	if expression.UserID != user.UserId {
		utils.ResponseFail(c, "您只能删除自己的帖子", nil)
		return
	}

	//删除表白失败
	err = a.expressionService.Delete(requestBody)
	if err != nil {
		utils.ResponseFail(c, "删除表白失败", nil)
		return
	}

	//返回成功相应
	utils.ResponseOkWithoutData(c)
}

//修改表白

func (a ExpressController) Edit(c *gin.Context) {
	var requestBody model.ExpressionUpdateRequestJsonObject

	if err := c.BindJSON(&requestBody); err != nil {
		utils.ResponseFail(c, "无效的请求参数", err)
		return
	}
	user, err := a.expressionService.FindUserByUserId(requestBody.UserId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ResponseFailWithoutData(c, "该用户未授权")
		return
	}
	if err != nil {
		utils.ResponseFail(c, "获取用户信息失败", nil)
		return
	}

	var expression model.Expression

	expression, err = a.expressionService.FindExpressionByExpressionId(expression.ExpressionID)
	if err != nil {
		utils.ResponseFail(c, "获取表白信息失败", nil)
		return
	}

	//判断创建表白的用户和请求者是否为同一人
	if expression.UserID != user.UserId {
		utils.ResponseFail(c, "您只能修改自己的帖子", nil)
		return
	}

	//修改表白失败
	err = a.expressionService.Edit(requestBody)
	if err != nil {
		utils.ResponseFail(c, "修改表白失败", nil)
		return
	}

	//返回成功响应
	utils.ResponseOkWithoutData(c)
}

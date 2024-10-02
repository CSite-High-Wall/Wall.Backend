package controller

import (
	"errors"
	"strconv"
	"unicode/utf8"
	"wall-backend/internal/model"
	"wall-backend/internal/service"
	"wall-backend/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ExpressController struct {
	userService       service.UserService
	expressionService service.ExpressionService
}

func NewExpressController(userService service.UserService, expressionService service.ExpressionService) ExpressController {
	return ExpressController{
		expressionService: expressionService,
		userService:       userService,
	}
}

// 发布表白
func (controller ExpressController) Publish(c *gin.Context) {
	var requestBody model.ExpressionCreateRequestJsonObject
	var userId = utils.ParseUserIdFromRequest(c)

	if err := c.BindJSON(&requestBody); err != nil || userId == uuid.Nil {
		utils.ResponseFail(c, "无效的请求参数", err)
		return
	}

	_, error := controller.userService.FindUserByUserId(userId)

	if errors.Is(error, gorm.ErrRecordNotFound) {
		utils.ResponseFailWithoutData(c, "未找到该用户") // 检查用户
	} else if error != nil {
		utils.ResponseFailWithoutData(c, "获取用户信息失败") // 检查用户
	} else if utf8.RuneCountInString(requestBody.Content) > 800 {
		utils.ResponseFailWithoutData(c, "限制的文本，文本过长") // 防过长文本
	} else if error := controller.expressionService.Publish(userId, requestBody); error != nil {
		utils.ResponseFailWithoutData(c, "发布表白失败")
	} else {
		utils.ResponseOkWithoutData(c) // 返回成功响应
	}
}

// 修改表白
func (controller ExpressController) Edit(c *gin.Context) {
	var requestBody model.ExpressionUpdateRequestJsonObject
	var userId = utils.ParseUserIdFromRequest(c)

	if err := c.BindJSON(&requestBody); err != nil || userId == uuid.Nil {
		utils.ResponseFail(c, "无效的请求参数", err)
		return
	}

	_, error := controller.userService.FindUserByUserId(userId)

	if errors.Is(error, gorm.ErrRecordNotFound) {
		utils.ResponseFailWithoutData(c, "未找到该用户") // 检查用户
	} else if error != nil {
		utils.ResponseFailWithoutData(c, "获取用户信息失败") // 检查用户
	} else {
		expression, error := controller.expressionService.FindExpressionByExpressionId(requestBody.ExpressionId)

		if error != nil {
			utils.ResponseFailWithoutData(c, "获取表白信息失败")
		} else if expression.UserId != userId {
			utils.ResponseFailWithoutData(c, "您只能修改自己的帖子") // 判断创建表白的用户和请求者是否为同一人
		} else if utf8.RuneCountInString(requestBody.Content) > 800 {
			utils.ResponseFailWithoutData(c, "限制的文本，文本过长") // 防过长文本
		} else if error := controller.expressionService.Edit(userId, requestBody); error != nil {
			utils.ResponseFailWithoutData(c, "修改表白失败")
		} else {
			utils.ResponseOkWithoutData(c)
		}
	}
}

// 删除表白
func (controller ExpressController) Delete(c *gin.Context) {
	var userId = utils.ParseUserIdFromRequest(c)
	var expressionId uint64 = 0

	if expression_id, isUserIdExist := c.GetQuery("expression_id"); !isUserIdExist {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	} else {
		expressionId, _ = strconv.ParseUint(expression_id, 10, 32)
	}

	_, error := controller.userService.FindUserByUserId(userId)

	if errors.Is(error, gorm.ErrRecordNotFound) {
		utils.ResponseFailWithoutData(c, "未找到该用户") // 检查用户
	} else if error != nil {
		utils.ResponseFailWithoutData(c, "获取用户信息失败") // 检查用户
	} else {
		expression, error := controller.expressionService.FindExpressionByExpressionId(expressionId)

		if error != nil {
			utils.ResponseFailWithoutData(c, "获取表白信息失败")
		} else if expression.UserId != userId {
			utils.ResponseFailWithoutData(c, "您只能删除自己的帖子") // 判断创建表白的用户和请求者是否为同一人
		} else if error = controller.expressionService.Delete(userId, expressionId); error != nil {
			utils.ResponseFailWithoutData(c, "删除表白失败") // 删除表白失败
		} else {
			utils.ResponseOkWithoutData(c) // 返回成功相应
		}
	}
}

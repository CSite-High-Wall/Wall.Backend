package controller

import (
	"errors"
	"strconv"
	"wall-backend/internal/service"
	"wall-backend/pkg/utils"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type CommunityController struct {
	userService       service.UserService
	expressionService service.ExpressionService
}

func NewCommunityController(userService service.UserService, expressionService service.ExpressionService) CommunityController {
	return CommunityController{
		userService:       userService,
		expressionService: expressionService,
	}
}

// 获取所有表白
func (controller CommunityController) FetchAllExpression(c *gin.Context) {
	expressions, err := controller.expressionService.FetchAllExpression()

	if err != nil {
		utils.ResponseFailWithoutData(c, "获取表白列表失败") // 如果查询出错，返回内部服务器错误
	} else {
		var expressionList []gin.H               // 准备响应数据
		for _, expression := range expressions { // 遍历所有表白，将每个表白的信息添加到expressionList中

			user, error := controller.userService.FindUserByUserId(expression.UserId)
			if error != nil {
				continue
			}

			expressionList = append(expressionList, gin.H{
				"expression_id": expression.ExpressionId,
				"user_name":     user.UserName,
				"user_id":       expression.UserId,
				"content":       expression.Content,
				"title":         expression.Title,
				"time":          expression.CreatedAt.Format("2006-01-02 15:04:05"), // 格式化时间为易读格式
			})
		}

		utils.ResponseOk(c, gin.H{
			"expression_list": expressionList, // 准备最终响应
		}) // 返回成功响应，包含所有表白信息
	}
}

// 获取指定表白
func (controller CommunityController) FetchTargetedExpression(c *gin.Context) {
	var expressionId uint64 = 0

	if expression_id, isUserIdExist := c.GetQuery("expression_id"); !isUserIdExist {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	} else {
		expressionId, _ = strconv.ParseUint(expression_id, 10, 32)
	}

	expression, err := controller.expressionService.FindExpressionByExpressionId(expressionId)
	if err != nil {
		utils.ResponseFailWithoutData(c, "获取指定表白失败") // 如果查询出错，返回内部服务器错误
	} else {
		user, error := controller.userService.FindUserByUserId(expression.UserId)
		if error != nil {
			utils.ResponseFailWithoutData(c, "获取指定表白失败")
		} else {
			utils.ResponseOk(c, gin.H{
				"expression_id": expression.ExpressionId,
				"user_name":     user.UserName,
				"user_id":       expression.UserId,
				"content":       expression.Content,
				"title":         expression.Title,
				"time":          expression.CreatedAt.Format("2006-01-02 15:04:05"), // 格式化时间为易读格式
			})
		}
	}
}

// 获取自己的所有表白
func (controller CommunityController) FetchUserExpression(c *gin.Context) {
	var userId = utils.ParseUserIdFromRequest(c) //获取请求体地UserId
	user, error := controller.userService.FindUserByUserId(userId)

	if errors.Is(error, gorm.ErrRecordNotFound) {
		utils.ResponseFailWithoutData(c, "未找到该用户") // 检查用户
	} else if error != nil {
		utils.ResponseFailWithoutData(c, "获取用户信息失败")
	} else {
		expressions, err := controller.expressionService.FetchUserExpression(userId)

		if err != nil {
			utils.ResponseFailWithoutData(c, "获取个人表白列表失败") // 如果查询出错，返回内部服务器错误
		} else {
			var expressionList []gin.H
			for _, expression := range expressions { // 遍历表白，将特定表白的信息添加到expressionList中
				expressionList = append(expressionList, gin.H{
					"expression_id": expression.ExpressionId,
					"user_id":       expression.UserId,
					"user_name":     user.UserName,
					"content":       expression.Content,
					"title":         expression.Title,
					"time":          expression.CreatedAt,
				})
			}

			utils.ResponseOk(c, gin.H{
				"expression_list": expressionList,
			}) // 返回成功响应
		}
	}
}

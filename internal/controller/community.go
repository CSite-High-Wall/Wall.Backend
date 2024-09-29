package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"wall-backend/internal/model"
	"wall-backend/internal/service"
	"wall-backend/pkg/utils"
)

type CommunityController struct {
	community   service.CommunityService
	userService service.UserService
}

func (controller CommunityController) AllExpression(c *gin.Context) {
	var expressions []model.Expression
	expressions, err := controller.community.AllExpression()
	if err != nil {
		// 如果查询出错，返回内部服务器错误
		utils.ResponseFail(c, "服务器内部错误，获取表白失败", nil)
		return
	}
	// 准备响应数据
	var expressionList []gin.H
	for _, expression := range expressions {
		// 遍历所有表白，将每个表白的信息添加到expressionList中
		expressionList = append(expressionList, gin.H{
			"expression_id": expression.ExpressionId,
			"user_id":       expression.UserId,
			"content":       expression.Content,
			"title":         expression.Title,
			"time":          expression.CreatedAt,
		})
	}
	// 准备最终响应
	responseData := gin.H{
		"expression_list": expressionList,
	}

	// 返回成功响应，包含所有表白信息
	utils.ResponseOk(c, responseData)
}

// 根据ExpressionId获取表白
func (controller CommunityController) GetExpressionById(c *gin.Context) {
	//获取ExpressionId
	expressionId, err := strconv.ParseUint(c.Param("expression_id"), 10, 32)
	if err != nil {
		utils.ResponseFail(c, "无效的表白ID", nil)
		return
	}

	expression, err := controller.community.GetExpressionById(uint(expressionId))
	if err != nil {
		utils.ResponseFail(c, "服务器内部错误，获取表白失败", nil)
		return
	}

	if expression.ExpressionId == 0 {
		utils.ResponseFail(c, "该表白不存在", nil)
		return
	}

	responseData := gin.H{
		"expression_id": expression.ExpressionId,
		"user_id":       expression.UserId,
		"content":       expression.Content,
		"title":         expression.Title,
		"time":          expression.CreatedAt,
	}

	// 返回成功响应，包含特定表白信息
	utils.ResponseOk(c, responseData)
}

// 获取自己的所有表白
func (controller CommunityController) GetMyExpressionById(c *gin.Context) {

	//获取请求体地UserId
	var userId = utils.ParseUserIdFromRequest(c)

	_, error := controller.userService.FindUserByUserId(userId)

	if errors.Is(error, gorm.ErrRecordNotFound) {
		utils.ResponseFailWithoutData(c, "未找到该用户") // 检查用户
	} else if error != nil {
		utils.ResponseFailWithoutData(c, "获取用户信息失败")
		if error != nil {
			utils.ResponseFail(c, "无效的表白ID", nil)
			return
		}
	}

	var expressionList []gin.H
	for _, expression := range expression {
		// 遍历表白，将特定表白的信息添加到expressionList中
		expressionList = append(expressionList, gin.H{
			"expression_id": expression.ExpressionId,
			"user_id":       expression.UserId,
			"content":       expression.Content,
			"title":         expression.Title,
			"time":          expression.CreatedAt,
		})
	}

	if len(expressionList) == 0 {
		utils.ResponseFail(c, "该表白不存在", nil)
		return
	}

	responseData := gin.H{
		"expression_list": expressionList,
	}
	// 返回成功响应，包含特定表白信息
	utils.ResponseOk(c, responseData)
}

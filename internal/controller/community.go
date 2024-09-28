package controller

import (
	"github.com/gin-gonic/gin"
	"wall-backend/internal/model"
	"wall-backend/internal/service"
	"wall-backend/pkg/utils"
)

type CommunityController struct {
	community service.CommunityService
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
			"time":          expression.CreatedAt.Format("2006-01-02 15:04:05"), // 格式化时间为易读格式
		})
	}
	// 准备最终响应
	responseData := gin.H{
		"expression_list": expressionList,
	}

	// 返回成功响应，包含所有表白信息
	utils.ResponseOk(c, responseData)

}

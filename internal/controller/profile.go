package controller

import (
	"errors"
	"wall-backend/internal/service"
	"wall-backend/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProfileController struct {
	userService       service.UserService
	expressionService service.ExpressionService
}

func NewProfileController(userService service.UserService, expressionService service.ExpressionService) ProfileController {
	return ProfileController{
		userService:       userService,
		expressionService: expressionService,
	}
}

func (controller ProfileController) GetUserInfo(c *gin.Context) {
	var userId = utils.ParseUserIdFromRequest(c)
	response, error := controller.userService.GetUserInfoByUserId(userId)

	if error != nil {
		utils.ResponseFailWithoutData(c, "获取用户信息失败")
	} else {
		utils.ResponseOk(c, response)
	}
}

// 获取自己的所有表白
func (controller ProfileController) FetchUserExpressions(c *gin.Context) {
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

				var displayUserName string = "匿名用户"
				var displayAvatar string = ""

				if !expression.Anonymity {
					displayUserName = user.UserName
					displayAvatar = user.AvatarUrl
				}

				expressionList = append(expressionList, gin.H{
					"expression_id": expression.ExpressionId,
					"user_id":       expression.UserId,
					"user_name":     displayUserName,
					"avatar_url":    displayAvatar,
					"content":       expression.Content,
					"title":         expression.Title,
					"time":          expression.CreatedAt.Format("2006-01-02 15:04:05"),
				})
			}

			if len(expressionList) == 0 {
				var expressionList [0]gin.H

				utils.ResponseOk(c, gin.H{
					"expression_list": expressionList, // 准备最终响应
				}) // 返回成功响应，包含所有表白信息
			} else {
				utils.ResponseOk(c, gin.H{
					"expression_list": expressionList, // 准备最终响应
				}) // 返回成功响应，包含所有表白信息
			}
		}
	}
}

// 上传用户头像 Url
func (controller ProfileController) UploadUserAvatarUrl(c *gin.Context) {
	var userId = utils.ParseUserIdFromRequest(c)
	var avatarUrl string

	if avatar_url, isUserIdExist := c.GetQuery("avatar_url"); !isUserIdExist {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	} else {
		avatarUrl = avatar_url
	}

	error := controller.userService.UploadUserAvatarUrl(userId, avatarUrl)

	if error != nil {
		utils.ResponseFailWithoutData(c, "上传用户头像 Url 失败")
	} else {
		utils.ResponseOkWithoutData(c)
	}
}

package controller

import (
	"errors"
	"wall-backend/internal/service"
	"wall-backend/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BlacklistController struct {
	userService       service.UserService
	blacklistService  service.BlacklistService
	expressionService service.ExpressionService
}

func NewBlacklistController(userService service.UserService, blacklistService service.BlacklistService, expressionService service.ExpressionService) BlacklistController {
	return BlacklistController{
		blacklistService:  blacklistService,
		userService:       userService,
		expressionService: expressionService,
	}
}

func (controller BlacklistController) GetUserBlacklist(c *gin.Context) {
	var userId = utils.ParseUserIdFromRequest(c)

	_, error := controller.userService.FindUserByUserId(userId)
	if errors.Is(error, gorm.ErrRecordNotFound) {
		utils.ResponseFailWithoutData(c, "未找到该用户") // 检查用户
	} else if error != nil {
		utils.ResponseFailWithoutData(c, "获取用户信息失败") // 检查用户
	} else {
		blacklist, err := controller.blacklistService.GetUserBlacklistByUserId(userId)
		if err != nil {
			utils.ResponseFailWithoutData(c, "服务器内部错误，获取拉黑名单失败")
			return
		}

		var blacklistList []gin.H
		for _, blacklistItem := range blacklist {
			blockedUser, error := controller.userService.FindUserByUserId(blacklistItem.BlockedUserId)
			if error != nil {
				continue
			}

			blacklistList = append(blacklistList, gin.H{
				"owner_user_id":           blacklistItem.OwnerUserId,
				"blocked_user_id":         blacklistItem.BlockedUserId,
				"blocked_user_name":       blockedUser.UserName,
				"blocked_user_avatar_url": blockedUser.AvatarUrl,
			})
		}

		if len(blacklist) == 0 {
			utils.ResponseOk(c, gin.H{
				"blacklist": [0]gin.H{}, // 准备最终响应
			})
		} else {
			utils.ResponseOk(c, gin.H{
				"blacklist": blacklistList, // 准备最终响应
			})
		}
	}

}

func (controller BlacklistController) AddUserIntoBlacklist(c *gin.Context) {
	var userId = utils.ParseUserIdFromRequest(c)
	exist, blockedUserId := utils.TryGetUuid(c, "blocked_user_id")

	if !exist {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	}

	_, error := controller.userService.FindUserByUserId(userId)
	if errors.Is(error, gorm.ErrRecordNotFound) {
		utils.ResponseFailWithoutData(c, "未找到该用户") // 检查用户
	} else if error != nil {
		utils.ResponseFailWithoutData(c, "获取用户信息失败") // 检查用户
	} else if blockedUserId == userId {
		utils.ResponseFailWithoutData(c, "你不能屏蔽你自己") // 检查用户
	} else {
		_, error := controller.userService.FindUserByUserId(blockedUserId)

		if errors.Is(error, gorm.ErrRecordNotFound) {
			utils.ResponseFailWithoutData(c, "未找到该被屏蔽用户")
		} else if error != nil {
			utils.ResponseFailWithoutData(c, "获取被屏蔽用户信息失败")
		} else if error := controller.blacklistService.AddBlacklistUserItem(userId, blockedUserId); error != nil {
			utils.ResponseFailWithoutData(c, "添加被屏蔽用户失败")
		} else {
			utils.ResponseOkWithoutData(c)
		}
	}
}

func (controller BlacklistController) RemoveUserFromBlacklist(c *gin.Context) {
	var userId = utils.ParseUserIdFromRequest(c)
	exist, blockedUserId := utils.TryGetUuid(c, "blocked_user_id")

	if !exist {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	}

	_, error := controller.userService.FindUserByUserId(userId)
	if errors.Is(error, gorm.ErrRecordNotFound) {
		utils.ResponseFailWithoutData(c, "未找到该被屏蔽用户")
	} else if error != nil {
		utils.ResponseFailWithoutData(c, "获取被屏蔽用户信息失败")

	} else {
		if error := controller.blacklistService.RemoveBlacklistUserItem(userId, blockedUserId); error != nil {
			utils.ResponseFailWithoutData(c, "添加被屏蔽用户失败")
		} else {
			utils.ResponseOkWithoutData(c)
		}
	}
}

func (controller BlacklistController) GetExpressionBlacklist(c *gin.Context) {
	var userId = utils.ParseUserIdFromRequest(c)

	_, error := controller.userService.FindUserByUserId(userId)
	if errors.Is(error, gorm.ErrRecordNotFound) {
		utils.ResponseFailWithoutData(c, "未找到该用户") // 检查用户
	} else if error != nil {
		utils.ResponseFailWithoutData(c, "获取用户信息失败") // 检查用户
	} else {
		blacklist, err := controller.blacklistService.GetExpressionBlacklistByUserId(userId)
		if err != nil {
			utils.ResponseFailWithoutData(c, "服务器内部错误，获取拉黑名单失败")
			return
		}

		var blacklistList []gin.H
		for _, blacklistExpression := range blacklist {
			blockedEexpression, error := controller.expressionService.FindExpressionByExpressionId(blacklistExpression.ExpressionId)
			if error != nil {
				continue
			}

			blacklistList = append(blacklistList, gin.H{
				"owner_user_id":                 blockedEexpression.UserId,
				"blocked_expression_id":         blockedEexpression.ExpressionId,
				"blocked_expression_title":      blockedEexpression.Title,
				"blocked_expression_content":    blockedEexpression.Content,
				"blocked_expression_anonymity":  blockedEexpression.Anonymity,
				"blocked_expression_created_at": blockedEexpression.CreatedAt,
				"blocked_expression_updated_at": blockedEexpression.UpdatedAt,
				"blocked_expression_deleted_at": blockedEexpression.DeletedAt,
			})
		}

		if len(blacklist) == 0 {
			utils.ResponseOk(c, gin.H{
				"blacklist": [0]gin.H{}, // 准备最终响应
			})
		} else {
			utils.ResponseOk(c, gin.H{
				"blacklist": blacklistList, // 准备最终响应
			})
		}
	}
}

func (controller BlacklistController) AddExpressionIntoBlacklist(c *gin.Context) {
	var userId = utils.ParseUserIdFromRequest(c)
	exist, blockedExpressionId := utils.TryGetUInt64(c, "blocked_expression_id")

	if !exist {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	}

	_, error := controller.userService.FindUserByUserId(userId)
	if errors.Is(error, gorm.ErrRecordNotFound) {
		utils.ResponseFailWithoutData(c, "未找到该用户") // 检查用户
	} else if error != nil {
		utils.ResponseFailWithoutData(c, "获取用户信息失败") // 检查用户
	} else {
		_, error := controller.expressionService.FindExpressionByExpressionId(blockedExpressionId)

		if error != nil {
			utils.ResponseFailWithoutData(c, "获取表白信息失败")
		} else if error := controller.blacklistService.AddBlacklistExpressionItem(userId, blockedExpressionId); error != nil {
			utils.ResponseFailWithoutData(c, "增加屏蔽表白失败")
		} else {
			utils.ResponseOkWithoutData(c)
		}
	}
}

func (controller BlacklistController) RemoveExpressionFromBlacklist(c *gin.Context) {
	var userId = utils.ParseUserIdFromRequest(c)
	exist, blockedExpressionId := utils.TryGetUInt64(c, "blocked_expression_id")

	if !exist {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	}

	_, error := controller.userService.FindUserByUserId(userId)
	if errors.Is(error, gorm.ErrRecordNotFound) {
		utils.ResponseFailWithoutData(c, "未找到该用户") // 检查用户
	} else if error != nil {
		utils.ResponseFailWithoutData(c, "获取用户信息失败") // 检查用户
	} else {
		if error := controller.blacklistService.RemoveBlacklistExpressionItem(userId, blockedExpressionId); error != nil {
			utils.ResponseFailWithoutData(c, "添加被屏蔽用户失败")
		} else {
			utils.ResponseOkWithoutData(c)
		}
	}
}

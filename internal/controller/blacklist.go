package controller

import (
	"errors"
	"github.com/google/uuid"
	"wall-backend/internal/service"
	"wall-backend/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BlacklistController struct {
	userService      service.UserService
	blacklistService service.BlacklistService
}

func NewBlacklistController(userService service.UserService, blacklistService service.BlacklistService) BlacklistController {
	return BlacklistController{
		blacklistService: blacklistService,
		userService:      userService,
	}
}

func (controller BlacklistController) GetBlacklistOfUser(c *gin.Context) {
	var userId = utils.ParseUserIdFromRequest(c)

	_, error := controller.userService.FindUserByUserId(userId)
	if errors.Is(error, gorm.ErrRecordNotFound) {
		utils.ResponseFailWithoutData(c, "未找到该用户") // 检查用户
	} else if error != nil {
		utils.ResponseFailWithoutData(c, "获取用户信息失败") // 检查用户
	} else {
		blacklist, err := controller.blacklistService.FindBlacklistItemsByUserId(userId)
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
func (controller BlacklistController) Add(c *gin.Context) {
	var userId = utils.ParseUserIdFromRequest(c)
	var blockedUserId uuid.UUID = uuid.Nil

	if _blockedUserId, isUserIdExist := c.GetQuery("blocked_user_id"); !isUserIdExist {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	} else {
		blockedUserId, _ = uuid.Parse(_blockedUserId)
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
		} else if error := controller.blacklistService.Add(userId, blockedUserId); error != nil {
			utils.ResponseFailWithoutData(c, "添加被屏蔽用户失败")
		} else {
			utils.ResponseOkWithoutData(c)
		}
	}
}

func (controller BlacklistController) Remove(c *gin.Context) {
	var userId = utils.ParseUserIdFromRequest(c)
	var blockedUserId uuid.UUID = uuid.Nil

	if _blockedUserId, isUserIdExist := c.GetQuery("blocked_user_id"); !isUserIdExist {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	} else {
		blockedUserId, _ = uuid.Parse(_blockedUserId)
	}

	_, error := controller.userService.FindUserByUserId(userId)
	if errors.Is(error, gorm.ErrRecordNotFound) {
		utils.ResponseFailWithoutData(c, "未找到该被屏蔽用户")
	} else if error != nil {
		utils.ResponseFailWithoutData(c, "获取被屏蔽用户信息失败")

	} else {
		if error := controller.blacklistService.Remove(userId, blockedUserId); error != nil {
			utils.ResponseFailWithoutData(c, "添加被屏蔽用户失败")
		} else {
			utils.ResponseOkWithoutData(c)
		}
	}
}

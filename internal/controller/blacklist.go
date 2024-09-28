package controller

import (
	"errors"
	"wall-backend/internal/model"
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

func (controller BlacklistController) Get(c *gin.Context) {
	var requestBody []model.Blacklist
	var userId = utils.ParseUserIdFromRequest(c)

	if error := c.BindJSON(&requestBody); error != nil {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	}

	_, error := controller.userService.FindUserByUserId(userId)
	if errors.Is(error, gorm.ErrRecordNotFound) {
		utils.ResponseFailWithoutData(c, "未找到该用户") // 检查用户
	} else if error != nil {
		utils.ResponseFailWithoutData(c, "获取用户信息失败") // 检查用户
	} else {
		blacklist, err := controller.blacklistService.AllBlacklist()
		if err != nil {
			utils.ResponseFail(c, "服务器内部错误，获取拉黑名单失败", nil)
			return
		}
		var blacklistList []gin.H
		for _, blacklist := range blacklist {
			blacklistList = append(blacklistList, gin.H{
				"blacklist_id": blacklist.ID,
				"bing_id":      blacklist.BeingId,
				"pull_id":      blacklist.PullId,
				"time":         blacklist.Time.Format("2006-01-02 15:04:05"), // 格式化时间为易读格式
			})
		}
		// 准备最终响应
		responseData := gin.H{
			"blacklist_list": blacklistList,
		}

		// 返回成功响应，包含所有表白信息
		utils.ResponseOk(c, responseData)

	}

}
func (controller BlacklistController) Add(c *gin.Context) {
	var requestBody model.BlacklistCreateRequestJsonObject
	var userId = utils.ParseUserIdFromRequest(c)

	if err := c.BindJSON(&requestBody); err != nil {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	}

	_, error := controller.userService.FindUserByUserId(userId)
	if errors.Is(error, gorm.ErrRecordNotFound) {
		utils.ResponseFailWithoutData(c, "未找到该用户") // 检查用户
	} else if error != nil {
		utils.ResponseFailWithoutData(c, "获取用户信息失败") // 检查用户
	} else {
		_, error := controller.userService.FindUserByUserId(requestBody.BeingId)
		if errors.Is(error, gorm.ErrRecordNotFound) {
			utils.ResponseFailWithoutData(c, "未找到该被拉黑用户")
		} else if error != nil {
			utils.ResponseFailWithoutData(c, "获取被拉黑用户信息失败")
		} else if error := controller.blacklistService.Add(userId, requestBody); error != nil {
			utils.ResponseFailWithoutData(c, "拉黑失败")
		} else {
			utils.ResponseOkWithoutData(c)
		}
	}
}
func (controller BlacklistController) Remove(c *gin.Context) {
	var requestBody model.BlacklistDeleteRequestJsonObject
	var userId = utils.ParseUserIdFromRequest(c)

	if err := c.BindJSON(&requestBody); err != nil {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	}

	_, error := controller.userService.FindUserByUserId(userId)
	if errors.Is(error, gorm.ErrRecordNotFound) {
		utils.ResponseFailWithoutData(c, "未找到该用户") // 检查用户
	} else if error != nil {
		utils.ResponseFailWithoutData(c, "获取用户信息失败") // 检查用户

	} else {
		blacklist, error := controller.blacklistService.FindBlacklistById(requestBody.ID)
		if error != nil {
			utils.ResponseFailWithoutData(c, "获取拉黑信息失败")
		} else if blacklist.PullId != userId {
			utils.ResponseFailWithoutData(c, "您只能取消自己的拉黑")
		} else if error = controller.blacklistService.Remove(userId, requestBody); error != nil {
			utils.ResponseFailWithoutData(c, "删除表白失败") // 删除评论失败
		} else {
			utils.ResponseOkWithoutData(c) // 返回成功相应
		}
	}
}

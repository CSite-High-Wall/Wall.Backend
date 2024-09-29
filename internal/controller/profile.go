package controller

import (
	"wall-backend/internal/service"
	"wall-backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	userService service.UserService
}

func NewProfileController(userService service.UserService) ProfileController {
	return ProfileController{
		userService: userService,
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

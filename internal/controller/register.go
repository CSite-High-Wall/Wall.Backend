package controller

import (
	"wall-backend/internal/model"
	"wall-backend/internal/service"
	"wall-backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

type RegisterController struct {
	userService service.UserService
}

func NewRegisterController(service service.UserService) RegisterController {
	return RegisterController{
		userService: service,
	}
}

func (controller RegisterController) Register(c *gin.Context) {
	var requestBody model.RegisterRequestJsonObject

	if error := c.BindJSON(&requestBody); error != nil {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	}

	if contains, _ := controller.userService.ContainsUserName(requestBody.UserName); contains {
		utils.ResponseFailWithoutData(c, "用户名已存在")
	} else if vaild, message := utils.CheckRegisterRequest(requestBody); !vaild {
		utils.ResponseFailWithoutData(c, message)
	} else if error := controller.userService.RegisterUser(requestBody); error != nil {
		utils.ResponseFailWithoutData(c, "注册用户失败")
	} else {
		utils.ResponseOkWithoutData(c)
	}
}

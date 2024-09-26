package controller

import (
	"errors"
	"wall-backend/internal/model"
	"wall-backend/internal/service"
	"wall-backend/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	authService service.AuthService
	userService service.UserService
}

func NewAuthController(authService service.AuthService, userService service.UserService) AuthController {
	return AuthController{
		authService: authService,
		userService: userService,
	}
}

func (controller AuthController) Authenticate(c *gin.Context) {
	var requestBody model.LoginRequestJsonObject

	if error := c.BindJSON(&requestBody); error != nil {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	}

	user, error := controller.userService.FindUserByUserName(requestBody.UserName)

	if errors.Is(error, gorm.ErrRecordNotFound) {
		utils.ResponseFailWithoutData(c, "用户名不存在")
	} else if user.Password != requestBody.Password {
		utils.ResponseFailWithoutData(c, "密码错误")
	} else {
		response, error := controller.authService.Authenticate(user.UserId)

		if error != nil {
			utils.ResponseFailWithoutData(c, "登录验证失败")
		} else {
			utils.ResponseOk(c, response)
		}
	}
}

func (controller AuthController) Refresh(c *gin.Context) {
	var requestBody model.AuthTokenRequestJsonObject

	if error := c.BindJSON(&requestBody); error != nil {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	}

	if vaild, userId := controller.authService.VerifyAccessToken(requestBody.AccessToken); vaild {
		response, error := controller.authService.Authenticate(userId)

		if error != nil {
			utils.ResponseFailWithoutData(c, "刷新验证失败")
		} else {
			utils.ResponseOk(c, response)
		}
	} else {
		utils.ResponseFailWithoutData(c, "登录验证失败")
	}
}

func (controller AuthController) Signout(c *gin.Context) {
	var requestBody model.AuthTokenRequestJsonObject

	if error := c.BindJSON(&requestBody); error != nil {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	}

	if vaild, userId := controller.authService.VerifyAccessToken(requestBody.AccessToken); vaild {
		error := controller.authService.Signout(userId)

		if error != nil {
			utils.ResponseFailWithoutData(c, "登出验证失败")
		} else {
			utils.ResponseOkWithoutData(c)
		}
	} else {
		utils.ResponseFailWithoutData(c, "登录验证失败")
	}
}

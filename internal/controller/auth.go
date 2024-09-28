package controller

import (
	"errors"
	"net/http"
	"wall-backend/internal/model"
	"wall-backend/internal/service"
	"wall-backend/pkg/utils"

	"golang.org/x/crypto/bcrypt"

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

// 登录验证接口
func (controller AuthController) Authenticate(c *gin.Context) {
	var requestBody model.LoginRequestJsonObject

	if error := c.BindJSON(&requestBody); error != nil {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	}

	user, error := controller.userService.FindUserByUserName(requestBody.UserName)

	if errors.Is(error, gorm.ErrRecordNotFound) {
		utils.ResponseFailWithoutData(c, "用户名不存在")
	} else if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password)) != nil {
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

// 刷新令牌接口
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

// 登出验证接口
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

// 验证令牌接口
func (controller AuthController) Validate(c *gin.Context) {
	var requestBody model.AuthTokenRequestJsonObject

	if error := c.BindJSON(&requestBody); error != nil {
		utils.ResponseFailWithoutData(c, "missing parameters")
		return
	}

	if vaild, _ := controller.authService.VerifyAccessToken(requestBody.AccessToken); vaild {
		utils.ResponseFrom(c, http.StatusOK, "令牌有效", nil)
	} else {
		utils.ResponseFailWithoutData(c, "令牌失效")
	}
}

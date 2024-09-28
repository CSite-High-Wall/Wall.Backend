package utils

import (
	"regexp"
	"wall-backend/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CheckRegisterRequest(requestBody model.RegisterRequestJsonObject) (bool, string) {
	var vaild bool = true
	var message string

	if regex := regexp.MustCompile(`^[a-zA-Z0-9_]{3,30}$`); !regex.MatchString(requestBody.UserName) {
		vaild = false
		message = "用户名必须是：3-30位字符，只允许数字、大小写字母、下划线"
	} else if regex := regexp.MustCompile(`^[a-zA-Z0-9@#$%^&*]{8,30}$`); !regex.MatchString(requestBody.Password) {
		vaild = false
		message = "密码必须是：8-30位字符，只允许数字、大小写字母、以及 @ # $ % ^ & * 字符"
	}

	return vaild, message
}

func ParseUserIdFromRequest(c *gin.Context) uuid.UUID {
	var userId uuid.UUID = uuid.Nil
	value, exist := c.Get("user_id")

	if exist {
		userId = value.(uuid.UUID)
	}

	return userId
}

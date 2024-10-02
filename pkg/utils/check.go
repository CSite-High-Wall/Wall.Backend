package utils

import (
	"regexp"
	"unicode/utf8"
	"wall-backend/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CheckRegisterRequest(requestBody model.RegisterRequestJsonObject) (bool, string) {
	var valid bool = true
	var message string

	if regex := regexp.MustCompile(`^[a-zA-Z0-9_]{3,30}$`); !regex.MatchString(requestBody.UserName) {
		valid = false
		message = "用户名必须是：3-30位字符，只允许数字、大小写字母、下划线"
	} else if regex := regexp.MustCompile(`^[a-zA-Z0-9@#$%^&*]{8,30}$`); !regex.MatchString(requestBody.Password) {
		valid = false
		message = "密码必须是：8-30位字符，只允许数字、大小写字母、以及 @ # $ % ^ & * 字符"
	}

	return valid, message
}

func TruncateText(text string, limit int) string {
	charCount := utf8.RuneCountInString(text) // 获取文本字符数

	// 如果字符数小于等于限制，直接返回原文本
	if charCount <= limit {
		return text
	}

	// 将字符串转换为 []rune，方便按字符截取
	runes := []rune(text)

	// 返回截取后的文本
	return string(runes[:limit]) + "..."
}

func ParseUserIdFromRequest(c *gin.Context) uuid.UUID {
	var userId uuid.UUID = uuid.Nil
	value, exist := c.Get("user_id")

	if exist {
		userId = value.(uuid.UUID)
	}

	return userId
}

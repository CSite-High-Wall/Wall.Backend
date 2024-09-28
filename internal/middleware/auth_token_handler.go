package middleware

import (
	"net/http"
	"wall-backend/internal/service"
	"wall-backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

var AuthService service.AuthService

func AuthToken(c *gin.Context) {
	accessToken := c.Request.Header.Get("Authorization")
	vaild, uuid := AuthService.VerifyAccessToken(accessToken)

	if !vaild {
		utils.ResponseFrom(c, http.StatusUnauthorized, "未授权的验证令牌或验证令牌已过期", nil)
		c.Abort()
	} else {
		c.Set("user_id", uuid)
		c.Next()
	}
}

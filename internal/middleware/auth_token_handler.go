package middleware

import (
	"net/http"
	"strings"
	"wall-backend/internal/service"
	"wall-backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

var AuthService service.AuthService

func AuthToken(c *gin.Context) {
	authorization := c.Request.Header.Get("Authorization")

	parts := strings.SplitN(authorization, " ", 2)

	if !(len(parts) == 2 && parts[0] == "Bearer") {
		utils.ResponseFrom(c, http.StatusBadRequest, "请求头中的 Authorization 格式错误", nil)
		return
	}

	accessToken := parts[1]
	vaild, uuid := AuthService.VerifyAccessToken(accessToken)

	if !vaild {
		utils.ResponseFrom(c, http.StatusUnauthorized, "未授权的验证令牌或验证令牌已过期", nil)
		c.Abort()
	} else {
		c.Set("user_id", uuid)
		c.Next()
	}
}

package middleware

import (
	"net/http"
	"wall-backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

func NotFoundHandler(c *gin.Context) {
	utils.ResponseFrom(c, 404, http.StatusText(http.StatusNotFound), nil)
	c.Abort()
}

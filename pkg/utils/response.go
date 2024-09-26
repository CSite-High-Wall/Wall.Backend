package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseFrom(c *gin.Context, httpStatus int, message string, data interface{}) {
	c.JSON(httpStatus, gin.H{
		"code":    httpStatus,
		"message": message,
		"data":    data,
	})
}

func ResponseOk(c *gin.Context, data interface{}) {
	ResponseFrom(c, http.StatusOK, "success", data)
}

func ResponseOkWithoutData(c *gin.Context) {
	ResponseOk(c, nil)
}

func ResponseFail(c *gin.Context, msg string, data interface{}) {
	ResponseFrom(c, http.StatusBadRequest, msg, nil)
}

func ResponseFailWithoutData(c *gin.Context, msg string) {
	ResponseFail(c, msg, nil)
}

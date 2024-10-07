package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strconv"
	"strings"
	"wall-backend/internal/service"
)

func TryGetUserId(c *gin.Context, service service.AuthService) (bool, uuid.UUID) {
	authorization := c.Request.Header.Get("Authorization")
	parts := strings.SplitN(authorization, " ", 2)

	if len(parts) == 2 && parts[0] == "Bearer" {
		accessToken := parts[1]
		valid, uuid := service.VerifyAccessToken(accessToken)

		if valid {
			return true, uuid
		}
	}

	return false, uuid.Nil
}

func TryGetUInt64(c *gin.Context, key string) (bool, uint64) {
	if _value, isExist := c.GetQuery(key); !isExist {
		return false, 0
	} else if value, error := strconv.ParseUint(_value, 10, 32); error == nil {
		return true, value
	}

	return false, 0
}

func TryGetUuid(c *gin.Context, key string) (bool, uuid.UUID) {
	if _value, isExist := c.GetQuery(key); !isExist {
		return false, uuid.Nil
	} else if value, error := uuid.Parse(_value); error == nil {
		return true, value
	}

	return false, uuid.Nil
}

func TryGetBool(c *gin.Context, key string) (bool, bool) {
	if _value, isExist := c.GetQuery(key); !isExist {
		return false, false
	} else if value, error := strconv.ParseBool(_value); error == nil {
		return true, value
	}

	return false, false
}

func TryGetString(c *gin.Context, key string) (bool, string) {
	if _value, isExist := c.GetQuery(key); !isExist {
		return false, ""
	} else {
		return true, _value
	}
}

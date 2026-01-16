package base

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

const (
	ReqKeyUserId = "user_id"
)

func MustAuth(c *gin.Context) (uuid.UUID, error) {
	userID, ok := GetUserID(c)
	if !ok {
		return uuid.Nil, errors.New("unauthorized")
	}
	return userID, nil
}

func GetUserID(c *gin.Context) (uuid.UUID, bool) {
	userID, exists := c.Get(ReqKeyUserId)
	if !exists {
		return uuid.UUID{}, false
	}

	if uid, ok := userID.(uuid.UUID); ok {
		return uid, true
	}

	return uuid.UUID{}, false
}

func StdErr(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}

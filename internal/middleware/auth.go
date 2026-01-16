package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/todo-app/pkg/base"
	"github.com/todo-app/pkg/util"
	"net/http"
	"strings"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")
		claims, err := util.ParseToken(tokenStr)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userID, err := uuid.Parse(claims.Subject)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid user id in token",
			})
			return
		}

		c.Set(base.ReqKeyUserId, userID)
		c.Next()
	}
}

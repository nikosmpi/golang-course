package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikosmpi/gorestapi/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No token provided"})
		return
	}
	userId, err := utils.ValidateToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}
	context.Set("userId", userId)
	context.Next()
}

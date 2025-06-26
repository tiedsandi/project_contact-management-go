package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tiedsandi/project_contact-management-go/utils"
)

func Authenticate(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	if authHeader == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": "Missing Authorization header"})
		return
	}

	token := strings.TrimSpace(authHeader)
	claims, err := utils.ValidateToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": "Invalid token"})
		return
	}

	context.Set("userId", claims.UserID)
	context.Set("username", claims.Username)
	context.Set("name", claims.Name)

	context.Next()
}

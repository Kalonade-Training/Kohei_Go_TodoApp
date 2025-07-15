package middleware
import (
	"strings"
	"net/http"
	"goTodoApp/domain/services"
	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware(tokenService services.ITokenService) gin.HandlerFunc{
	return func(c *gin.Context) {
		//Authorizationヘッダーからトークンを取得
		authHeader := c.GetHeader("Authorization")
		if authHeader == ""{
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
			c.Abort()
			return
		}
		//Bearer トークン形式を検証
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer"{
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			c.Abort()
			return
		}
		//トークンを検証
		token := parts[1]
		userID, err := tokenService.ValidateJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}
		//userIDをコンテキストに設定
		c.Set("userID", userID)
		c.Next()
	}
}
package server

import (
	"jobTracker/internal/utils"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddlware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("Authorization")
		if tokenStr == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Autohrization header required."})
			ctx.Abort()
			return
		}

		if utils.IsBlacklisted(tokenStr) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or Unauthorized."})
			ctx.Abort()
			return
		}

		claims, err := utils.ValidateJWT(tokenStr)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token!"})
			ctx.Abort()
			return
		}

		ctx.Set("username", claims.Username)
		ctx.Next()
	}

}

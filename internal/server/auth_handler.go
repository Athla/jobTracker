package server

import (
	"jobTracker/internal/auth"
	"jobTracker/internal/utils"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s *Server) LoginHandler(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errorf("Unable to bind json due: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	authSvc := auth.New(s.db)
	if err := authSvc.ValidateCredentials(req.Username, req.Password); err != nil {
		log.Errorf("Unable to validate credentials due: %s", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := authSvc.GenerateToken(req.Username)
	if err != nil {
		log.Errorf("Unable to generate token for '%s' due: %s", req.Username, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"type":  "Bearer",
	})
}

func (s *Server) LogoutHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
		return
	}

	claims, err := auth.ValidateToken(token)
	if err != nil {
		log.Errorf("Invalid token provided due: %s", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	utils.AddTokenToBlackList(token, claims.ExpiresAt)

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged out",
	})
}

package server

import (
	"jobTracker/internal/models"
	"jobTracker/internal/utils"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterUserHandler(ctx *gin.Context) {
	var input models.User
	if err := ctx.ShouldBindJSON(&input); err != nil {
		log.Errorf("Unalbe to read input due: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input."})
		return
	}

	if input.Username == "" || input.Password == "" {
		log.Errorf("Unalbe to read input due it being empty.")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required."})
		return
	}
	err := models.CreateUser(s.db, input.Username, input.Password)
	if err != nil {
		log.Errorf("Unable to create user due: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unable to create user."})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully."})
}

func (s *Server) LoginUserHandler(ctx *gin.Context) {
	var input models.User

	if err := ctx.ShouldBindJSON(&input); err != nil {
		log.Errorf("Unalbe to read input due: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input."})
		return
	}

	user, err := models.GetUserByUsername(s.db, input.Username)
	if err != nil || user.CheckPwd(input.Password) != nil {
		log.Errorf("Unable to proceed due invalid credentials.")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials."})
		return
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		log.Errorf("Unable to generate token due: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to generate token."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (s *Server) LogoutUserHandler(ctx *gin.Context) {
	tokenStr := ctx.GetHeader("Authorization")
	if tokenStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Auth header required."})
		return
	}

	claims, err := utils.ValidateJWT(tokenStr)
	if err != nil {
		log.Errorf("Unable to validate JSON due: %v", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Auth token invalid or expired."})
		return
	}

	utils.AddTokenToBlackList(tokenStr, claims.ExpiresAt)

	ctx.JSON(http.StatusOK, gin.H{"message": "Logged out successfully!"})
}

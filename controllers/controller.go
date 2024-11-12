package controllers

import (
	domain "valorInsight/domain"
	interfaces "valorInsight/domain/interfaces"
	"valorInsight/infrastructure"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	UserUsecase interfaces.UserUsecase
}

func NewController(userUsecase interfaces.UserUsecase) *Controller {
	return &Controller{
		UserUsecase: userUsecase,
	}
}

func (c *Controller) RegisterUser(ctx *gin.Context) {
	var user domain.User
	ctx.BindJSON(&user)
	token, err := c.UserUsecase.RegisterUser(user)
	if err.Message != "" {
		ctx.JSON(400, gin.H{"error": err.Message})
		return
	}
	ctx.JSON(200, gin.H{"token": token})
}

func (c *Controller) RefreshToken(ctx *gin.Context) {
	var refresh_token = gin.Param("refresh_token")
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	token, err := infrastructure.ValidateJWT(req.RefreshToken)
	if err != nil || !token.Valid {
		ctx.JSON(401, gin.H{"error": "Invalid or expired refresh token"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["user_id"] == nil {
		ctx.JSON(500, gin.H{"error": "Invalid token claims"})
		return
	}

	user := domain.User{
		UserId:    claims["user_id"].(string),
		Email:     claims["email"].(string),
		FirstName: claims["first_name"].(string),
		LastName:  claims["last_name"].(string),
		Username:  claims["username"].(string),
		Role:      claims["role"].(string),
	}

	newAccessToken, _, tokenErr := infrastructure.GenerateJWT(user)
	if tokenErr.Message != "" {
		ctx.JSON(500, gin.H{"error": tokenErr.Message})
		return
	}

	ctx.JSON(200, gin.H{"access_token": newAccessToken})
}

func (c *Controller) Login(ctx *gin.Context) {
	

package controllers

import (
	"fmt"
	"net/http"
	domain "valorInsight/domain"
	interfaces "valorInsight/domain/interfaces"
	"valorInsight/infrastructure"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	UserUsecase  interfaces.UserUsecase
	EmailUsecase interfaces.EmailUsecase
}

func NewController(userUsecase interfaces.UserUsecase, emailUsecase interfaces.EmailUsecase) *Controller {
	return &Controller{
		UserUsecase:  userUsecase,
		EmailUsecase: emailUsecase,
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

func (c *Controller) SendCode(ctx *gin.Context) {
	var request struct {
		Email string `json:"email"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	fmt.Println("Controller", request.Email)

	if err := c.EmailUsecase.SendVerificationCode(ctx, request.Email); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Verification code sent"})
}

func (c *Controller) VerifyCode(ctx *gin.Context) {
	var request struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := c.EmailUsecase.VerifyCode(ctx, request.Email, request.Code); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Email verified"})
}

func (c *Controller) RefreshToken(ctx *gin.Context) {
	var refresh_token = ctx.Param("refresh_token")
	if refresh_token == "" {
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	token, err := infrastructure.ValidateJWT(refresh_token)
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
	var user domain.User
	ctx.BindJSON(&user)

	token, err := c.UserUsecase.Login(user)

	if err.Message != "" {
		ctx.JSON(400, gin.H{"error": err.Message})
		return
	}
	ctx.JSON(200, gin.H{"token": token})
}

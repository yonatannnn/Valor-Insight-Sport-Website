package controllers

import (
	domain "valorInsight/domain"
	interfaces "valorInsight/domain/interfaces"

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

// RegisterUser godoc
func (c *Controller) RegisterUser(ctx *gin.Context) {
	var user domain.User
	ctx.BindJSON(&user)
	token, err := c.UserUsecase.RegisterUser(user)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"token": token})
}

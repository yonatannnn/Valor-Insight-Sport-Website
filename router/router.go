package router

import (
	"time"
	"valorInsight/controllers"
	"valorInsight/infrastructure"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(tc *controllers.Controller) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Correct field for specifying allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "Origin", "Accept"},
		ExposeHeaders:    []string{"Content-Length"}, // Optional, exposes headers to the frontend
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/register", tc.RegisterUser)
	r.POST("/login", tc.Login)
	r.POST("/refresh", tc.RefreshToken)
	r.POST("/send-code", tc.SendCode)
	r.POST("/verify-code", tc.VerifyCode)
	auth := r.Group("/")
	auth.Use(infrastructure.JWTMiddleware())
	{
		// auth.GET("/tasks", tc.GetAllTasks)
		// auth.GET("/tasks/:id", tc.GetTaskById)
		// auth.Use(middleware.AdminOnly())
		// {
		// 	auth.POST("/tasks", tc.CreateTask)
		// 	auth.PUT("/tasks/:id", tc.UpdateTask)
		// 	auth.DELETE("/tasks/:id", tc.DeleteTask)
		// 	auth.POST("/promote/:id", tc.PromoteUser)
		// }
	}

	return r
}

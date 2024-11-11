package routers

import (
	"valorInsight/controllers"
	"valorInsight/infrastructure"

	"github.com/gin-gonic/gin"
)

func SetupRouter(tc *controllers.Controller) *gin.Engine {
	r := gin.Default()
	r.POST("/register", tc.RegisterUser)

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

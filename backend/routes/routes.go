package routes

import (
	"pharmacy/controllers"
	"pharmacy/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	protected := r.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	{
		// Добавьте защищенные маршруты здесь
	}
}

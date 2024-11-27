package main

import (
	"pharmacy/config"
	"pharmacy/database"
	"pharmacy/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitAuthDB()
	database.InitPharmacyDB()

	r := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:3001"}, // Разрешаем только фронтенд на порту 3001
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}

	// Применяем CORS middleware
	r.Use(cors.New(corsConfig))
	routes.SetupRoutes(r)

	r.Run(":3000")
}

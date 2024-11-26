package main

import (
	"pharmacy/config"
	"pharmacy/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	r := gin.Default()

	r.Use(cors.Default())
	routes.SetupRoutes(r)

	r.Run(":3000")
}

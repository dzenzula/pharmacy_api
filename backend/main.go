package main

import (
	"log"
	"pharmacy/blockchain"
	"pharmacy/database"
	"pharmacy/routes"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitAuthDB()
	database.InitPharmacyDB()

	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatalf("Failed to connect to Ganache: %v", err)
	}
	blockchain.SetClient(client)

	r := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:3001", "http://192.168.0.102"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}

	// Применяем CORS middleware
	r.Use(cors.New(corsConfig))
	routes.SetupRoutes(r)

	r.Run("192.168.0.101:3000")
}

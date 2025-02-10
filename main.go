package main

import (
	"log"
	"receipt-processor/config"
	"receipt-processor/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	r := gin.Default()
	r.POST("/receipts/process", handlers.ProcessReceipt)
	r.GET("/receipts/:id/points", handlers.GetPoints)

	log.Printf("Starting server on %s", cfg.ServerPort)
	r.Run(cfg.ServerPort)
}

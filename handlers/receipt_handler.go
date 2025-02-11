package handlers

import (
	"net/http"
	"regexp"
	"fmt"

	"receipt-processor/models"
	"receipt-processor/services"
	"receipt-processor/storage"
	"receipt-processor/utils"
	"receipt-processor/config"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// validateReceipt checks if the receipt structure adheres to the required format
func validateReceipt(receipt models.Receipt) error {
	retailerPattern := regexp.MustCompile(`^[\w\s\-&]+$`)
	totalPattern := regexp.MustCompile(`^\d+\.\d{2}$`)
	itemPattern := regexp.MustCompile(`^[\w\s\-]+$`)

	if !retailerPattern.MatchString(receipt.Retailer) {
		return fmt.Errorf("Invalid retailer format")
	}
	if !totalPattern.MatchString(receipt.Total) {
		return fmt.Errorf("Invalid total format")
	}
	if len(receipt.Items) < 1 {
		return fmt.Errorf("At least one item is required")
	}
	for _, item := range receipt.Items {
		if !itemPattern.MatchString(item.ShortDescription) {
			return fmt.Errorf("Invalid item description format")
		}
		if !totalPattern.MatchString(item.Price) {
			return fmt.Errorf("Invalid item price format")
		}
	}
	return nil
}

// ProcessReceipt handles POST /receipts/process
func ProcessReceipt(c *gin.Context) {
	var receipt models.Receipt

	if err := c.ShouldBindJSON(&receipt); err != nil {
		if utils.IsGeneratedByLLM() {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Please verify input."})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "The receipt is invalid."})
		}
		return
	}

	if err := validateReceipt(receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uuid.New().String()
	cfg := config.LoadConfig()
	store := storage.NewRedisStore(cfg.RedisAddr)

	if err := store.SaveReceipt(id, receipt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save receipt. Please try again."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// GetPoints handles GET /receipts/{id}/points
func GetPoints(c *gin.Context) {
	id := c.Param("id")
	cfg := config.LoadConfig()
	store := storage.NewRedisStore(cfg.RedisAddr)

	receipt, err := store.GetReceipt(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No receipt found for that ID."})
		return
	}

	points := services.CalculatePoints(receipt)
	if points < 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error calculating points. Please retry."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"points": points})
}
package tests

import (
	"receipt-processor/models"
	"receipt-processor/services"
	"testing"
)

// TestCalculatePoints - Validate point calculation logic
func TestCalculatePoints(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Total:        "35.35",
		Items: []models.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
		},
	}

	expectedPoints := 28
	calculatedPoints := services.CalculatePoints(receipt)

	if calculatedPoints != expectedPoints {
		t.Errorf("Expected %d points, got %d", expectedPoints, calculatedPoints)
	}
}
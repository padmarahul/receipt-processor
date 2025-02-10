package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"receipt-processor/handlers"
	"receipt-processor/models"
	"testing"

	"github.com/gin-gonic/gin"
)

// setupRouter initializes a test router
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/receipts/process", handlers.ProcessReceipt)
	r.GET("/receipts/:id/points", handlers.GetPoints)
	return r
}

// TestProcessReceipt - Validate receipt processing API
func TestProcessReceipt(t *testing.T) {
	router := setupRouter()

	receipt := models.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Total:        "35.35",
		Items: []models.Item{
			{ShortDescription: "Milk", Price: "6.49"},
		},
	}

	jsonData, _ := json.Marshal(receipt)
	req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)

	if _, exists := response["id"]; !exists {
		t.Errorf("Response does not contain 'id'")
	}
}

// TestGetPoints - Validate retrieving points for a processed receipt
func TestGetPoints(t *testing.T) {
	router := setupRouter()

	// Process a receipt first to generate an ID
	receipt := models.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Total:        "35.35",
		Items: []models.Item{
			{ShortDescription: "Milk", Price: "6.49"},
		},
	}

	jsonData, _ := json.Marshal(receipt)
	req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	receiptID := response["id"]

	// Now test retrieving points
	req, _ = http.NewRequest("GET", "/receipts/"+receiptID+"/points", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var pointsResponse map[string]int
	json.Unmarshal(w.Body.Bytes(), &pointsResponse)

	if _, exists := pointsResponse["points"]; !exists {
		t.Errorf("Response does not contain 'points'")
	}
}
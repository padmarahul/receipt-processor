package services

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"receipt-processor/models"
	"receipt-processor/utils"
)

// CalculatePoints applies business rules to compute receipt points
func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// Retailer Name Alphanumeric Count
	reg := regexp.MustCompile("[^a-zA-Z0-9]")
	cleanedRetailer := reg.ReplaceAllString(receipt.Retailer, "")
	points += len(cleanedRetailer)

	// Total-Based Points
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err == nil {
		if total == math.Floor(total) {
			points += 50
		}
		if math.Mod(total, 0.25) == 0 {
			points += 25
		}
		if utils.IsGeneratedByLLM() && total > 10.00 {
			points += 5
		}
	}

	// Item-Based Points
	points += (len(receipt.Items) / 2) * 5 // 5 points for every two items

	for _, item := range receipt.Items {
		// Trim all spaces, including excess spaces inside the description
		description := strings.ReplaceAll(strings.TrimSpace(item.ShortDescription), "  ", " ")

		// If description length is a multiple of 3, apply price multiplier
		if len(description)%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err == nil {
				points += int(math.Ceil(price * 0.2)) // Multiply by 0.2 and round up
			}
		}
	}

	// Purchase Date Check (Odd Day)
	dateParts := strings.Split(receipt.PurchaseDate, "-")
	if len(dateParts) == 3 {
		day, err := strconv.Atoi(dateParts[2])
		if err == nil && day%2 == 1 {
			points += 6
		}
	}

	// Purchase Time Check (Between 2:00 PM and 4:00 PM)
	timeParts := strings.Split(receipt.PurchaseTime, ":")
	if len(timeParts) == 2 {
		hour, err := strconv.Atoi(timeParts[0])
		if err == nil && hour >= 14 && hour < 16 {
			points += 10
		}
	}

	return points
}
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
    "github.com/google/uuid"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Item struct {
	ShortDescription string `json:"shortDescription" binding:"required"`
	Price            string `json:"price" binding:"required"`
}

type Receipt struct {
	Retailer     string `json:"retailer" binding:"required"`
	PurchaseDate string `json:"purchaseDate" binding:"required"`
	PurchaseTime string `json:"purchaseTime" binding:"required"`
	Items        []Item `json:"items" binding:"required"`
	Total        string `json:"total" binding:"required"`
}

type ReceiptStore struct {
	sync.Mutex
	Receipts map[string]Receipt
	Points   map[string]int
}

var store = ReceiptStore{
	Receipts: make(map[string]Receipt),
	Points:   make(map[string]int),
}

func main() {
	router := gin.Default()

	router.POST("/receipts/process", processReceipt)
	router.GET("/receipts/:id/points", getPoints)

	router.Run(":8080")
}

// Endpoint: /receipts/process
func processReceipt(c *gin.Context) {
	var receipt Receipt
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := generateID() // Function to generate a unique ID
	store.Lock()
	store.Receipts[id] = receipt
	store.Points[id] = calculatePoints(receipt)
	store.Unlock()

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// Endpoint: /receipts/:id/points
func getPoints(c *gin.Context) {
	id := c.Param("id")

	store.Lock()
	points, exists := store.Points[id]
	store.Unlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "No receipt found for that id"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"points": points})
}

func generateID() string{
    return uuid.New().String()
}

// Function to calculate points for a receipt
func calculatePoints(receipt Receipt) int {
	points := 0

	points += countAlphanumeric(receipt.Retailer)

	if isRoundDollar(receipt.Total) {
		points += 50
	}

	if isMultipleOfQuarter(receipt.Total) {
		points += 25
	}

	points += (len(receipt.Items) / 2) * 5

	for _, item := range receipt.Items {
		points += pointsforDesc(item)
	}

	if isOddDay(receipt.PurchaseDate) {
		points += 6
	}

	if isBetweenTwoAndFour(receipt.PurchaseTime) {
		points += 10
	}

	return points
}


func countAlphanumeric(s string) int {
	alphanumeric := regexp.MustCompile("[a-zA-Z0-9]")
	return len(alphanumeric.FindAllString(s, -1))
}

func isRoundDollar(total string) bool {
	parts := strings.Split(total, ".")
	return parts[1] == "00"
}

func isMultipleOfQuarter(total string) bool {
	amount, _ := strconv.ParseFloat(total, 64)
	return math.Mod(amount, 0.25) == 0
}

func pointsforDesc(item Item) int {
	length := len(strings.TrimSpace(item.ShortDescription))
	if length%3 == 0 {
		price, _ := strconv.ParseFloat(item.Price, 64)
		return int(math.Ceil(price * 0.2))
	}
	return 0
}

//print statement is for my debugging case
func isOddDay(date string) bool {
    day, err := strconv.Atoi(strings.Split(date, "-")[2])
    if err != nil {
        fmt.Println("Date parse error:", err)
        return false
    }
    return day%2 != 0
}

//print statement is for my debugging case
func isBetweenTwoAndFour(timeStr string) bool {
    t, err := time.Parse("15:04", timeStr)
    if err != nil {
        fmt.Println("Time parse error:", err)
        return false
    }
    return t.Hour() >= 14 && t.Hour() < 16
}


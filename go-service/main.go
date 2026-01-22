package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Sensor struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	Location    string      `json:"location"`
	Value       interface{} `json:"value"`
	Unit        string      `json:"unit"`
	Status      string      `json:"status"`
	LastReading string      `json:"last_reading"`
}

type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

type ItemsResponse struct {
	Sensors []Sensor `json:"sensors"`
	Count   int      `json:"count"`
}

var dataFile = "/app/data/sensors.json"

func loadSensors() ([]Sensor, error) {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Sensor{}, nil
		}
		return nil, err
	}

	var sensors []Sensor
	if err := json.Unmarshal(data, &sensors); err != nil {
		return nil, err
	}

	return sensors, nil
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, HealthResponse{
		Status:  "ok",
		Service: "go",
	})
}

func itemsHandler(c *gin.Context) {
	sensors, err := loadSensors()
	if err != nil {
		log.Printf("Error loading sensors: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load sensors"})
		return
	}

	c.JSON(http.StatusOK, ItemsResponse{
		Sensors: sensors,
		Count:   len(sensors),
	})
}

func main() {
	router := gin.Default()

	router.GET("/health", healthHandler)
	router.GET("/items", itemsHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting Go IoT Sensor Service on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

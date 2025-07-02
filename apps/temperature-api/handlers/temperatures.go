package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TemperatureHandler struct {
}

type TemperatureResponse struct {
	Value       float64   `json:"value"`
	Unit        string    `json:"unit"`
	Timestamp   time.Time `json:"timestamp"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	SensorID    string    `json:"sensor_id"`
	SensorType  string    `json:"sensor_type"`
	Description string    `json:"description"`
}

func NewTemperatureHandler() *TemperatureHandler {
	return &TemperatureHandler{}
}

func (h *TemperatureHandler) RegisterRoutes(router *gin.RouterGroup) {
	temperature := router.Group("/temperature")
	{
		temperature.GET("", h.GetTemperatureByLocation)
		temperature.GET("/:sensorID", h.GetTemperatureByLocation)
	}
}

func (h *TemperatureHandler) GetTemperatureByLocation(c *gin.Context) {
	location := c.Param("location")
	sensorID := c.Param("sensorID")

	var tempRes TemperatureResponse
	log.Printf("Location is %s", location)
	log.Printf("sensorID is %s", sensorID)

	c.JSON(http.StatusOK, tempRes)
}

func (h *TemperatureHandler) GetTemperatureBySensorID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetTemperatureBySensorID"})
}

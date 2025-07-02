package handlers

import (
	"math/rand/v2"
	"net/http"
	"temperature-api/utils"
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
	location := c.Query("location")

	var tempRes TemperatureResponse

	tempRes.SensorID = utils.GenerateSensorID(location)
	tempRes.Location = location
	tempRes.Timestamp = time.Now()
	tempRes.Value = randomTemperature()

	c.JSON(http.StatusOK, tempRes)
}

func (h *TemperatureHandler) GetTemperatureBySensorID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetTemperatureBySensorID"})
}

func randomTemperature() float64 {
	min := 17
	max := 35
	return float64(rand.IntN(max-min) + min)
}

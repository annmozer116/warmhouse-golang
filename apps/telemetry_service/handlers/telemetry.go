package handlers

import (
	"math/rand/v2"
	"net/http"
	"strconv"
	"telemetry_service/models"
	"telemetry_service/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type TelemetryHandler struct {
}

type TelemetryResponse struct {
	Value float64 `json:"value"`
	// Unit        string    `json:"unit"`
	Status     string    `json:"status"`
	MetricType string    `json:"metric_type"`
	Timestamp  time.Time `json:"timestamp"`
	DeviceID   string    `json:"device_id"`
}

type RegisterDeviceRequest struct {
	DeviceType string `json:"device_type"`
	DeviceID   int    `json:"device_id"`
}

func NewTelemetryHandler() *TelemetryHandler {
	return &TelemetryHandler{}
}

func (h *TelemetryHandler) RegisterRoutes(router *gin.RouterGroup) {
	telemetry := router.Group("/telemetry")
	{
		telemetry.GET("", h.GetTelemetryByDeviceID)
	}
	// devices := router.Group("/devices")
	// {
	// 	devices.POST("", h.RegisterNewDevice)
	// }
}

// [open_close_status, on_off_status, temperature, luminance]

func (h *TelemetryHandler) GetTelemetryByDeviceID(c *gin.Context) {
	deviceTypes := []models.DeviceType{
		models.Light,
		models.Smoke,
		models.Temperature,
		models.Video,
	}
	limit_param := c.Query("limit")
	var limit int
	if limit_param != "" {
		limit_i, err := strconv.Atoi(limit_param)
		if err != nil || limit_i <= 0 {
			limit = 10
		} else {
			limit = limit_i
		}
	}

	deviceType := deviceTypes[rand.IntN(len(deviceTypes))]
	metrics := utils.GenerateDeviceStatus(deviceType, limit)

	c.JSON(http.StatusOK, metrics)
}

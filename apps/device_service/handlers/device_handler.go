package handlers

import (
	"context"
	"device_service/db"
	"device_service/models"
	"device_service/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeviceHandler struct {
	DB               *db.DB
	TelemetryService *services.TelemetryService
}

// NewDeviceHandler creates a new DeviceHandler
// telemetryService *services.TelemetryService
func NewDeviceHandler(db *db.DB, telemetryService *services.TelemetryService) *DeviceHandler {
	return &DeviceHandler{
		DB:               db,
		TelemetryService: telemetryService,
	}
}

// RegisterRoutes registers the Device routes
func (h *DeviceHandler) RegisterRoutes(router *gin.RouterGroup) {
	devices := router.Group("/devices")
	{
		devices.GET("", h.GetDevices)
		// devices.GET("/:id", h.GetDeviceByID)
		devices.POST("", h.CreateDevice)
		devices.POST("/transfer", h.CreateDeviceTransfer)
		// devices.PUT("/:id", h.UpdateDevice)
		// devices.DELETE("/:id", h.DeleteDevice)
		// devices.PATCH("/:id/value", h.UpdateDeviceValue)
		// devices.GET("/:id/telemetry", h.GetTelemetryByDeviceID)
	}
}

// GetDevices handles GET /api/v1/devices
func (h *DeviceHandler) GetDevices(c *gin.Context) {
	userIDHeader := c.GetHeader("X-user-id")
	if userIDHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing X-user-Id header"})
		return
	}

	userID, err := strconv.Atoi(userIDHeader)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}
	// Парсинг параметра location_id (опциональный)
	locationID := c.Query("location_id")
	var locID int
	if locationID != "" {
		id, err := strconv.Atoi(locationID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid location ID"})
			return
		}
		locID = id
	}
	type_code := c.Query("type")

	devices, _ := h.DB.GetDevices(context.Background(), userID, locID, type_code)
	for i, device := range devices {
		teleData, err := h.TelemetryService.GetTelemetryByDeviceID(device.ID, device.Type)
		log.Println("GetDevices handler, GetTelemetryByDeviceID finished")
		if err == nil {
			// Update device with real-time data
			devices[i].Value = float32(teleData.Value)
			devices[i].Status = teleData.Status
			devices[i].Unit = teleData.Unit
			devices[i].MetricType = teleData.MetricType
			log.Printf("Updated telemetry data for device %d from external API", device.ID)
		} else {
			log.Printf("Failed to fetch telemetry data for device %d: %v", device.ID, err)
		}
	}

	c.JSON(http.StatusOK, devices)

	log.Print("Get devices")
}

func (h *DeviceHandler) CreateDevice(c *gin.Context) {
	userIDHeader := c.GetHeader("X-user-id")
	if userIDHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing X-user-Id header"})
		return
	}

	userID, err := strconv.Atoi(userIDHeader)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	var deviceCreate models.DeviceCreate
	if err := c.ShouldBindJSON(&deviceCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	type_code, exists := models.GetDeviceType(deviceCreate.DeviceModelID)
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Device not supported"})
		return
	}
	deviceCreate.Type = type_code
	deviceCreate.OwnerID = userID

	device, err := h.DB.CreateDevice(context.Background(), deviceCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, device)
}

// CreateDeviceTransfer is used for monolyth integration. In: location name, convert to id, get type by model, save
func (h *DeviceHandler) CreateDeviceTransfer(c *gin.Context) {
	userIDHeader := c.GetHeader("X-user-id")
	if userIDHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing X-user-Id header"})
		return
	}

	userID, err := strconv.Atoi(userIDHeader)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	var deviceCreateTransfer models.DeviceCreateV0
	if err := c.ShouldBindJSON(&deviceCreateTransfer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	location_id := models.GetLocationId(deviceCreateTransfer.Location)

	deviceCreate := models.DeviceCreate{
		Name:          deviceCreateTransfer.Name,
		DeviceModelID: deviceCreateTransfer.DeviceModelID,
		Serial_Number: deviceCreateTransfer.Serial_Number,
		LocationID:    location_id,
	}
	type_code, exists := models.GetDeviceType(deviceCreate.DeviceModelID)
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Device not supported"})
		return
	}

	deviceCreate.LocationID = location_id
	deviceCreate.Type = type_code
	deviceCreate.OwnerID = userID

	device, err := h.DB.CreateDevice(context.Background(), deviceCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, device)
}

// GetDeviceByID handles GET /api/v1/devices/:id
// func (h *DeviceHandler) GetDeviceByID(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	log.Printf("Get devices by ID: %s", id)
// }

// CreateDevice handles POST /api/v1/devices
// func (h *DeviceHandler) CreateDevice(c *gin.Context) {
// 	log.Printf("CreateDevice device")
// }

// GetTelemetryByDeviceID handles GET /api/v1/devices/:id/telemetry
// func (h *DeviceHandler) GetTelemetryByDeviceID(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	log.Printf("GetTelemetryByDeviceID by ID: %s", id)
// }

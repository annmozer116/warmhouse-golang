package models

import (
	"errors"
	"time"
)

// SensorType represents the type of sensor
type DeviceType string

// [lighting, temperature, videocam, smoke_detector]
const (
	Temperature DeviceType = "temperature"
	Light       DeviceType = "lighting"
	Video       DeviceType = "videocam"
	Smoke       DeviceType = "smoke_detector"
)

type DeviceModelID int

// Константы идентификаторов моделей
const (
	Temp1  DeviceModelID = 1
	Temp2  DeviceModelID = 2
	Temp3  DeviceModelID = 3
	Smoke1 DeviceModelID = 4
	Smoke2 DeviceModelID = 5
	Cam1   DeviceModelID = 6
	Cam2   DeviceModelID = 7
	Cam3   DeviceModelID = 8
	Light1 DeviceModelID = 9
	Light2 DeviceModelID = 10
	Light3 DeviceModelID = 11
)

// маппинг моделей и типов
var modelTypeMapping = map[DeviceModelID]DeviceType{
	Temp1:  Temperature,
	Temp2:  Temperature,
	Temp3:  Temperature,
	Smoke1: Smoke,
	Smoke2: Smoke,
	Cam1:   Video,
	Cam2:   Video,
	Cam3:   Video,
	Light1: Light,
	Light2: Light,
	Light3: Light,
}

// GetDeviceType возвращает тип устройства по ID модели
func GetDeviceType(modelID DeviceModelID) (DeviceType, bool) {
	deviceType, exists := modelTypeMapping[modelID]
	return deviceType, exists
}

// Device for smart home
type Device struct {
	ID            int           `json:"id"`
	Name          string        `json:"name"`
	Type          DeviceType    `json:"type"`
	Status        string        `json:"status"`
	Serial_Number string        `json:"serial_number"`
	LocationID    string        `json:"location_id"`
	DeviceModelID DeviceModelID `json:"model_id"`
	LastUpdated   time.Time     `json:"last_updated"`
	CreatedAt     time.Time     `json:"created_at"`
}

// DeviceCreate represents the data needed to create a new device
type DeviceCreate struct {
	Name          string `json:"name" binding:"required"`
	LocationID    int    `json:"location_id" binding:"required"`
	Type          DeviceType
	DeviceModelID DeviceModelID `json:"model_id" binding:"required"`
	Serial_Number string        `json:"serial_number" binding:"required"`
	OwnerID       int
}

func NewDeviceCreate(name string, location_id int, model_id DeviceModelID, serial_number string, owner_id int) (*DeviceCreate, error) {
	type_code, exists := GetDeviceType(model_id)
	if !exists {
		return nil, errors.New("Device not supported")
	}
	device := &DeviceCreate{
		name,
		location_id,
		type_code,
		model_id,
		serial_number,
		owner_id,
	}
	return device, nil
}

// DeviceUpdate represents the data that can be updated for a device
type DeviceUpdate struct {
	Name       string `json:"name"`
	LocationID string `json:"location_id" binding:"required"`
	Status     string `json:"status"`
}

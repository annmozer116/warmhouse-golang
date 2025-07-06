package models

import (
	"fmt"
	"strings"
)

// DeviceType represents the type of device
type DeviceType string

// [lighting, temperature, videocam, smoke_detector]
const (
	Temperature DeviceType = "temperature"
	Light       DeviceType = "lighting"
	Video       DeviceType = "videocam"
	Smoke       DeviceType = "smoke_detector"
)

type Device struct {
	DeviceID   int        `json:"device_id"`
	DeviceType DeviceType `json:"device_type"`
}

// DeviceMetric содержит метрику устройства
type DeviceMetric struct {
	Type   string  `json:"type"`
	Value  float64 `json:"value"`
	Unit   string  `json:"unit"`
	Status string  `json:"status"`
}

func ConvertToDeviceType(s string) (DeviceType, error) {
	switch strings.ToLower(s) {
	case "temperature", "sensor", "thermostat":
		return Temperature, nil
	case "lighting", "light":
		return Light, nil
	case "videocam", "video":
		return Video, nil
	case "smoke_detector", "smoke":
		return Smoke, nil
	default:
		return "", fmt.Errorf("unknown device type: %s", s)
	}
}

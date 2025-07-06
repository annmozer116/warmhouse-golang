package models

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

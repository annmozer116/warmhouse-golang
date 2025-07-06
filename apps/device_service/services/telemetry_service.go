package services

import (
	"net/http"
	"time"
)

// TemperatureService handles fetching telemetry data from external API
type TelemetryService struct {
	BaseURL    string
	HTTPClient *http.Client
}

// TelemetryResponse represents the response from the telemetry API
type TelemetryResponse struct {
	Value      float64   `json:"value"`
	Unit       string    `json:"unit"`
	Timestamp  time.Time `json:"timestamp"`
	Status     string    `json:"status"`
	DeviceID   string    `json:"device_id"`
	DeviceType string    `json:"device_type"`
}

func NewTelemetryService(baseURL string) *TelemetryService {
	return &TelemetryService{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func main() {

}

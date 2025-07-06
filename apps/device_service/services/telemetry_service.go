package services

import (
	"device_service/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// TemperatureService handles fetching telemetry data from external API
type TelemetryService struct {
	BaseURL    string
	HTTPClient *http.Client
}

// TelemetryResponse represents the response from the telemetry API
type TelemetryResponse struct {
	// Timestamp  time.Time `json:"timestamp"`
	MetricType string  `json:"type"`
	Value      float64 `json:"value"`
	Unit       string  `json:"unit"`
	Status     string  `json:"status"`
}

func NewTelemetryService(baseURL string) *TelemetryService {
	return &TelemetryService{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *TelemetryService) GetTelemetryByDeviceID(device_id int, device_type models.DeviceType) (*TelemetryResponse, error) {
	url := fmt.Sprintf("%s/api/v1/telemetry?limit=1&device_type=%s&device_id=%s", s.BaseURL, device_type, strconv.Itoa(device_id))
	resp, err := s.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching telemetry data: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var telemetryResp TelemetryResponse
	if err := json.NewDecoder(resp.Body).Decode(&telemetryResp); err != nil {
		return nil, fmt.Errorf("error decoding telemetry response: %w", err)
	}
	return &telemetryResp, nil
}

package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"smarthome/models"
	"strconv"
	"time"
)

// TemperatureService handles fetching temperature data from external API
type DeviceAPIService struct {
	BaseURL    string
	HTTPClient *http.Client
}

// TemperatureResponse represents the response from the temperature API
type DeviceAPIResponse struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Type          string    `json:"type"`
	Status        string    `json:"status"`
	Serial_Number string    `json:"serial_number"`
	LocationID    string    `json:"location_id"`
	DeviceModelID string    `json:"model_id"`
	LastUpdated   time.Time `json:"last_updated"`
	CreatedAt     time.Time `json:"created_at"`
}

// DeviceAPICreateRequest represents the request for device creating in the device-service
type DeviceAPICreateRequest struct {
	Name          string `json:"name"`
	Location      string `json:"location"`
	DeviceModelID int    `json:"model_id"`
	Serial_Number string `json:"serial_number"`
}

// NewDeviceAPIService creates a new device api service
func NewDeviceAPIService(baseURL string) *DeviceAPIService {
	return &DeviceAPIService{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
func (s *DeviceAPIService) CreateDeviceForSensor(ctx context.Context, sensor models.SensorCreate, sensor_id int) (*DeviceAPIResponse, error) {
	url := fmt.Sprintf("%s/api/v1/devices/transfer", s.BaseURL)

	request := DeviceAPICreateRequest{
		Name:          sensor.Name,
		Location:      sensor.Location,
		DeviceModelID: 2,
		Serial_Number: "000-HCC-687" + strconv.Itoa(sensor_id),
	}
	// Подготовка тела запроса
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-user-id", "1")

	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var deviceResponse DeviceAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&deviceResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &deviceResponse, nil
}

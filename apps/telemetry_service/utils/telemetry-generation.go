package utils

import (
	"math/rand"
	"telemetry_service/models"
)

// GenerateDeviceStatus генерирует случайные данные для устройства
func GenerateDeviceStatus(deviceType models.DeviceType, limit int) []models.DeviceMetric {

	metrics := make([]models.DeviceMetric, limit)

	switch deviceType {
	case models.Light:
		for i := range limit {
			metrics[i] = randomLightingMetric()
		}

	case models.Smoke:
		for i := range limit {
			metrics[i] = randomSmokeMetric()
		}

	case models.Temperature:
		for i := range limit {
			metrics[i] = randomTemperatureMetric()
		}

	case models.Video:
		for i := range limit {
			metrics[i] = randomMotionMetric()
		}

	default:
		for i := range limit {
			metrics[i] = models.DeviceMetric{Status: "unknown"}
		}
	}
	return metrics
}

// Вспомогательные функции для генерации статусов
func randomLightingMetric() models.DeviceMetric {
	statuses := []string{"on", "off", "dimmed", "color_change"}
	lightStatus := statuses[rand.Intn(len(statuses))]
	metricObject := models.DeviceMetric{
		Status: lightStatus,
	}

	// Добавляем яркость только если устройство включено
	if lightStatus != "off" {
		brightness := float64(rand.Intn(101))
		metricObject.Value = brightness
		metricObject.Unit = "%"
		metricObject.Type = "Lums"
	}
	return metricObject
}

func randomSmokeMetric() models.DeviceMetric {
	smokeLevel := rand.Float32()
	status := "normal"
	if smokeLevel > 0.95 {
		status = "alert"
	}

	metricObject := models.DeviceMetric{
		Status: status,
		Value:  float64(smokeLevel),
		Type:   "concentration",
		Unit:   "mmol",
	}
	return metricObject
}

func randomTemperatureMetric() models.DeviceMetric {
	statuses := []string{"on", "off"}
	status := statuses[rand.Intn(len(statuses))]
	temperature := 15 + rand.Float64()*30 // 15-45°C

	metricObject := models.DeviceMetric{
		Status: status,
		Value:  temperature,
		Type:   "temperature",
		Unit:   "°C",
	}
	return metricObject
}

func randomMotionMetric() models.DeviceMetric {
	var status string
	var value int
	if rand.Float32() < 0.7 {
		status = "no_motion"
		value = 0
	} else {
		status = "motion_detected"
		value = 1
	}

	metricObject := models.DeviceMetric{
		Status: status,
		Value:  float64(value),
		Type:   "motion_detection",
		Unit:   "flag",
	}
	return metricObject
}

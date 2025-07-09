package utils

func GenerateLocation(sensorID string) string {
	var location string
	switch sensorID {
	case "1":
		location = "Living Room"
	case "2":
		location = "Bedroom"
	case "3":
		location = "Kitchen"
	default:
		location = "Unknown"
	}
	return location
}

func GenerateSensorID(location string) string {
	var sensorID string
	switch location {
	case "Living Room":
		sensorID = "1"
	case "Bedroom":
		sensorID = "2"
	case "Kitchen":
		sensorID = "3"
	default:
		sensorID = "0"
	}
	return sensorID
}

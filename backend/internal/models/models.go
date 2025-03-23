package models

import (
	"time"
)

// User представляет модель пользователя.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Device представляет модель устройства.
type Device struct {
	ID           int    `json:"id"`
	SerialNumber string `json:"serialNumber"`
	Model        string `json:"model"`
	Address      string `json:"address"`
}

// Telemetry представляет модель телеметрии.
type Telemetry struct {
	ID          int       `json:"id"`
	DeviceID    int       `json:"deviceId"`
	Timestamp   time.Time `json:"timestamp"`
	Temperature float64   `json:"temperature"`
	Humidity    float64   `json:"humidity"`
}

// LoginRequest представляет модель запроса на авторизацию.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse представляет модель ответа на авторизацию.
type LoginResponse struct {
	Token string `json:"token"`
}

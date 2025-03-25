package models

import "time"

// User представляет модель пользователя системы
// swagger:model User
type User struct {
	// ID пользователя
	// example: 1
	ID int `json:"id"`

	// Имя пользователя (логин)
	// example: admin
	Username string `json:"username"`

	// Пароль пользователя
	// example: admin123
	Password string `json:"password"`
}

// Device представляет модель устройства
// swagger:model Device
type Device struct {
	// ID устройства
	// example: 1
	ID int `json:"id"`

	// Серийный номер устройства
	// example: SN123456
	SerialNumber string `json:"serialNumber"`

	// Модель устройства
	// example: Model X
	Model string `json:"model"`

	// Адрес установки устройства
	// example: ул. Примерная, 123
	Address string `json:"address"`
}

// Telemetry представляет модель телеметрии
// swagger:model Telemetry
type Telemetry struct {
	// ID записи телеметрии
	// example: 1
	ID int `json:"id"`

	// ID устройства
	// example: 1
	DeviceID int `json:"deviceId"`

	// Временная метка записи
	// example: 2023-01-01T12:00:00Z
	Timestamp time.Time `json:"timestamp"`

	// Температура
	// example: 25.5
	Temperature float64 `json:"temperature"`

	// Влажность
	// example: 60.0
	Humidity float64 `json:"humidity"`
}

// LoginRequest представляет запрос на авторизацию
// swagger:model LoginRequest
type LoginRequest struct {
	// Имя пользователя
	// example: admin
	Username string `json:"username"`

	// Пароль пользователя
	// example: password123
	Password string `json:"password"`
}

// LoginResponse представляет ответ с токеном
// swagger:model LoginResponse
type LoginResponse struct {
	// JWT токен для аутентификации
	// example: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
	Token string `json:"token"`
}

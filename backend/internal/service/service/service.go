package service

import (
	"backend/internal/models"
	"backend/internal/repository"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type DeviceService interface {
	GetDevices(filter map[string]string, page, limit int) ([]models.Device, error)
	GetDeviceByID(id int) (*models.Device, error)
	AddDevice(device models.Device) error
	UpdateDevice(device models.Device) error
	DeleteDevice(id int) error
}

// TelemetryService определяет методы для работы с телеметрией.
type TelemetryService interface {
	GetTelemetry(deviceID int, start, end string) ([]models.Telemetry, error)
	AddTelemetry(telemetry models.Telemetry) error
	DeleteTelemetry(id int) error
}

// UserService определяет методы для работы с пользователями.
type UserService interface {
	Authenticate(username, password string) (string, error)
}

// deviceService реализует интерфейс DeviceService.
type deviceService struct {
	repo repository.DeviceRepository
}

// NewDeviceService создает новый экземпляр deviceService.
func NewDeviceService(repo repository.DeviceRepository) DeviceService {
	return &deviceService{repo: repo}
}

// GetDevices возвращает список устройств с фильтрацией и пагинацией.
func (s *deviceService) GetDevices(filter map[string]string, page, limit int) ([]models.Device, error) {
	return s.repo.GetDevices(filter, page, limit)
}

// GetDeviceByID возвращает устройство по его ID.
func (s *deviceService) GetDeviceByID(id int) (*models.Device, error) {
	return s.repo.GetDeviceByID(id)
}

// AddDevice добавляет новое устройство.
func (s *deviceService) AddDevice(device models.Device) error {
	return s.repo.AddDevice(device)
}

// UpdateDevice обновляет существующее устройство.
func (s *deviceService) UpdateDevice(device models.Device) error {
	return s.repo.UpdateDevice(device)
}

// DeleteDevice удаляет устройство по его ID.
func (s *deviceService) DeleteDevice(id int) error {
	return s.repo.DeleteDevice(id)
}

// telemetryService реализует интерфейс TelemetryService.
type telemetryService struct {
	repo repository.TelemetryRepository
}

// NewTelemetryService создает новый экземпляр telemetryService.
func NewTelemetryService(repo repository.TelemetryRepository) TelemetryService {
	return &telemetryService{repo: repo}
}

// GetTelemetry возвращает телеметрию устройства за указанный период.
func (s *telemetryService) GetTelemetry(deviceID int, start, end string) ([]models.Telemetry, error) {
	return s.repo.GetTelemetry(deviceID, start, end)
}

// AddTelemetry добавляет новую запись телеметрии.
func (s *telemetryService) AddTelemetry(telemetry models.Telemetry) error {
	return s.repo.AddTelemetry(telemetry)
}

// DeleteTelemetry удаляет запись телеметрии по её ID.
func (s *telemetryService) DeleteTelemetry(id int) error {
	return s.repo.DeleteTelemetry(id)
}

// userService реализует интерфейс UserService.
type userService struct {
	repo repository.UserRepository
}

// NewUserService создает новый экземпляр userService.
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// Authenticate выполняет аутентификацию пользователя и возвращает JWT-токен.
func (s *userService) Authenticate(username, password string) (string, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if user.Password != password {
		return "", errors.New("invalid credentials")
	}

	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Токен действителен 24 часа
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT secret key not configured")
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return tokenString, nil
}

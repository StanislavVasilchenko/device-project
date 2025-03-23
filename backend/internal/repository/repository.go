package repository

import (
	"backend/internal/models"
	"database/sql"
	"fmt"
)

// DeviceRepository определяет методы для работы с устройствами.
type DeviceRepository interface {
	GetDevices(filter map[string]string, page, limit int) ([]models.Device, error)
	GetDeviceByID(id int) (*models.Device, error)
	AddDevice(device models.Device) error
	UpdateDevice(device models.Device) error
	DeleteDevice(id int) error
}

// TelemetryRepository определяет методы для работы с телеметрией.
type TelemetryRepository interface {
	GetTelemetry(deviceID int, start, end string) ([]models.Telemetry, error)
	AddTelemetry(telemetry models.Telemetry) error
	DeleteTelemetry(id int) error
}

// UserRepository определяет методы для работы с пользователями.
type UserRepository interface {
	GetUserByUsername(username string) (*models.User, error)
}

// deviceRepository реализует интерфейс DeviceRepository.
type deviceRepository struct {
	db *sql.DB
}

// NewDeviceRepository создает новый экземпляр deviceRepository.
func NewDeviceRepository(db *sql.DB) DeviceRepository {
	return &deviceRepository{db: db}
}

// GetDevices возвращает список устройств с фильтрацией и пагинацией.
func (r *deviceRepository) GetDevices(filter map[string]string, page, limit int) ([]models.Device, error) {
	query := `SELECT id, serial_number, model, address FROM devices`
	var args []interface{}
	var conditions []string

	// Добавляем фильтры
	if serialNumber, ok := filter["serialNumber"]; ok {
		conditions = append(conditions, "serial_number = $1")
		args = append(args, serialNumber)
	}

	if len(conditions) > 0 {
		query += " WHERE " + conditions[0]
	}

	// Добавляем пагинацию
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, (page-1)*limit)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query devices: %w", err)
	}
	defer rows.Close()

	var devices []models.Device
	for rows.Next() {
		var device models.Device
		if err := rows.Scan(&device.ID, &device.SerialNumber, &device.Model, &device.Address); err != nil {
			return nil, fmt.Errorf("failed to scan device row: %w", err)
		}
		devices = append(devices, device)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over device rows: %w", err)
	}

	return devices, nil
}

// GetDeviceByID возвращает устройство по его ID.
func (r *deviceRepository) GetDeviceByID(id int) (*models.Device, error) {
	var device models.Device
	query := `SELECT id, serial_number, model, address FROM devices WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&device.ID, &device.SerialNumber, &device.Model, &device.Address)
	if err != nil {
		return nil, fmt.Errorf("failed to get device: %w", err)
	}
	return &device, nil
}

// AddDevice добавляет новое устройство.
func (r *deviceRepository) AddDevice(device models.Device) error {
	query := `INSERT INTO devices (serial_number, model, address) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, device.SerialNumber, device.Model, device.Address)
	if err != nil {
		return fmt.Errorf("failed to add device: %w", err)
	}
	return nil
}

// UpdateDevice обновляет существующее устройство.
func (r *deviceRepository) UpdateDevice(device models.Device) error {
	query := `UPDATE devices SET serial_number = $1, model = $2, address = $3 WHERE id = $4`
	_, err := r.db.Exec(query, device.SerialNumber, device.Model, device.Address, device.ID)
	if err != nil {
		return fmt.Errorf("failed to update device: %w", err)
	}
	return nil
}

// DeleteDevice удаляет устройство по его ID.
func (r *deviceRepository) DeleteDevice(id int) error {
	query := `DELETE FROM devices WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete device: %w", err)
	}
	return nil
}

// telemetryRepository реализует интерфейс TelemetryRepository.
type telemetryRepository struct {
	db *sql.DB
}

// NewTelemetryRepository создает новый экземпляр telemetryRepository.
func NewTelemetryRepository(db *sql.DB) TelemetryRepository {
	return &telemetryRepository{db: db}
}

// GetTelemetry возвращает телеметрию устройства за указанный период.
func (r *telemetryRepository) GetTelemetry(deviceID int, start, end string) ([]models.Telemetry, error) {
	query := `
        SELECT id, device_id, timestamp, temperature, humidity
        FROM telemetry
        WHERE device_id = $1 AND timestamp BETWEEN $2 AND $3
        ORDER BY timestamp
    `
	rows, err := r.db.Query(query, deviceID, start, end)
	if err != nil {
		return nil, fmt.Errorf("failed to query telemetry: %w", err)
	}
	defer rows.Close()

	var telemetry []models.Telemetry
	for rows.Next() {
		var t models.Telemetry
		if err := rows.Scan(&t.ID, &t.DeviceID, &t.Timestamp, &t.Temperature, &t.Humidity); err != nil {
			return nil, fmt.Errorf("failed to scan telemetry row: %w", err)
		}
		telemetry = append(telemetry, t)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over telemetry rows: %w", err)
	}

	return telemetry, nil
}

// AddTelemetry добавляет новую запись телеметрии.
func (r *telemetryRepository) AddTelemetry(telemetry models.Telemetry) error {
	query := `
        INSERT INTO telemetry (device_id, timestamp, temperature, humidity)
        VALUES ($1, $2, $3, $4)
    `
	_, err := r.db.Exec(query, telemetry.DeviceID, telemetry.Timestamp, telemetry.Temperature, telemetry.Humidity)
	if err != nil {
		return fmt.Errorf("failed to insert telemetry: %w", err)
	}
	return nil
}

// DeleteTelemetry удаляет запись телеметрии по её ID.
func (r *telemetryRepository) DeleteTelemetry(id int) error {
	query := `DELETE FROM telemetry WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete telemetry: %w", err)
	}
	return nil
}

// userRepository реализует интерфейс UserRepository.
type userRepository struct {
	db *sql.DB
}

// NewUserRepository создает новый экземпляр userRepository.
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

// GetUserByUsername возвращает пользователя по его имени.
func (r *userRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, password FROM users WHERE username = $1`
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}

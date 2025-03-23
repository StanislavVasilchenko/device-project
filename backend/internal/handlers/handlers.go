package handlers

import (
	"backend/internal/models"
	"backend/internal/service/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeviceHandler обрабатывает запросы, связанные с устройствами.
type DeviceHandler struct {
	service service.DeviceService
}

// NewDeviceHandler создает новый экземпляр DeviceHandler.
func NewDeviceHandler(service service.DeviceService) *DeviceHandler {
	return &DeviceHandler{service: service}
}

// GetDevices возвращает список устройств.
func (h *DeviceHandler) GetDevices(c *gin.Context) {
	// Получаем параметры запроса
	filter := make(map[string]string)
	for key, values := range c.Request.URL.Query() {
		filter[key] = values[0]
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	// Вызываем сервис для получения данных
	devices, err := h.service.GetDevices(filter, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем результат
	c.JSON(http.StatusOK, devices)
}

// GetDeviceByID возвращает устройство по его ID.
func (h *DeviceHandler) GetDeviceByID(c *gin.Context) {
	id := c.Param("id")
	deviceID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid device ID"})
		return
	}

	device, err := h.service.GetDeviceByID(deviceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, device)
}

// AddDevice добавляет новое устройство.
func (h *DeviceHandler) AddDevice(c *gin.Context) {
	var device models.Device
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := h.service.AddDevice(device); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "device added"})
}

// UpdateDevice обновляет существующее устройство.
func (h *DeviceHandler) UpdateDevice(c *gin.Context) {
	id := c.Param("id")
	deviceID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid device ID"})
		return
	}

	var device models.Device
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	device.ID = deviceID

	if err := h.service.UpdateDevice(device); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "device updated"})
}

// DeleteDevice удаляет устройство по его ID.
func (h *DeviceHandler) DeleteDevice(c *gin.Context) {
	id := c.Param("id")
	deviceID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid device ID"})
		return
	}

	if err := h.service.DeleteDevice(deviceID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// TelemetryHandler обрабатывает запросы, связанные с телеметрией.
type TelemetryHandler struct {
	service service.TelemetryService
}

// NewTelemetryHandler создает новый экземпляр TelemetryHandler.
func NewTelemetryHandler(service service.TelemetryService) *TelemetryHandler {
	return &TelemetryHandler{service: service}
}

// GetTelemetry возвращает телеметрию устройства за указанный период.
func (h *TelemetryHandler) GetTelemetry(c *gin.Context) {
	id := c.Param("id")
	deviceID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid device ID"})
		return
	}

	start := c.Query("start")
	end := c.Query("end")

	telemetry, err := h.service.GetTelemetry(deviceID, start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, telemetry)
}

// AddTelemetry добавляет новую запись телеметрии.
func (h *TelemetryHandler) AddTelemetry(c *gin.Context) {
	var telemetry models.Telemetry
	if err := c.ShouldBindJSON(&telemetry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := h.service.AddTelemetry(telemetry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "telemetry added"})
}

// DeleteTelemetry удаляет запись телеметрии по её ID.
func (h *TelemetryHandler) DeleteTelemetry(c *gin.Context) {
	id := c.Param("id")
	telemetryID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid telemetry ID"})
		return
	}

	if err := h.service.DeleteTelemetry(telemetryID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// AuthHandler обрабатывает запросы, связанные с аутентификацией.
type AuthHandler struct {
	service service.UserService
}

// NewAuthHandler создает новый экземпляр AuthHandler.
func NewAuthHandler(service service.UserService) *AuthHandler {
	return &AuthHandler{service: service}
}

// Login выполняет аутентификацию пользователя и возвращает JWT-токен.
func (h *AuthHandler) Login(c *gin.Context) {
	var loginRequest models.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	token, err := h.service.Authenticate(loginRequest.Username, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

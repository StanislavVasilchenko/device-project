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

// GetDevices godoc
// @Summary Получить список устройств
// @Description Получить список устройств с пагинацией и фильтрацией
// @Tags Устройства
// @Accept json
// @Produce json
// @Param serialNumber query string false "Фильтр по серийному номеру"
// @Param page query int false "Номер страницы" default(1)
// @Param limit query int false "Количество записей на странице" default(10)
// @Security ApiKeyAuth
// @Success 200 {array} models.Device
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /devices [get]
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

// GetDeviceByID godoc
// @Summary Получить устройство по ID
// @Description Получить детальную информацию об устройстве
// @Tags Устройства
// @Accept json
// @Produce json
// @Param id path int true "ID устройства"
// @Security ApiKeyAuth
// @Success 200 {object} models.Device
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /devices/{id} [get]
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

// AddDevice godoc
// @Summary Добавить новое устройство
// @Description Создать новое устройство в системе
// @Tags Устройства
// @Accept json
// @Produce json
// @Param device body models.Device true "Данные устройства"
// @Security ApiKeyAuth
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /devices [post]
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

// UpdateDevice godoc
// @Summary Обновить устройство
// @Description Обновить информацию об устройстве
// @Tags Устройства
// @Accept json
// @Produce json
// @Param id path int true "ID устройства"
// @Param device body models.Device true "Обновленные данные устройства"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /devices/{id} [put]
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

// DeleteDevice godoc
// @Summary Удалить устройство
// @Description Удалить устройство из системы
// @Tags Устройства
// @Accept json
// @Produce json
// @Param id path int true "ID устройства"
// @Security ApiKeyAuth
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /devices/{id} [delete]
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

// GetTelemetry godoc
// @Summary Получить телеметрию устройства
// @Description Получить данные телеметрии за указанный период
// @Tags Телеметрия
// @Accept json
// @Produce json
// @Param id path int true "ID устройства"
// @Param start query string true "Начальная дата (YYYY-MM-DD)"
// @Param end query string true "Конечная дата (YYYY-MM-DD)"
// @Security ApiKeyAuth
// @Success 200 {array} models.Telemetry
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /devices/{id}/telemetry [get]
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

// AddTelemetry godoc
// @Summary Добавить данные телеметрии
// @Description Добавить новую запись телеметрии для устройства
// @Tags Телеметрия
// @Accept json
// @Produce json
// @Param telemetry body models.Telemetry true "Данные телеметрии"
// @Security ApiKeyAuth
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /devices/{id}/telemetry [post]
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

// DeleteTelemetry godoc
// @Summary Удалить запись телеметрии
// @Description Удалить запись телеметрии по ID
// @Tags Телеметрия
// @Accept json
// @Produce json
// @Param id path int true "ID записи телеметрии"
// @Security ApiKeyAuth
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /devices/{id}/telemetry/{telemetryId} [delete]
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

// Login godoc
// @Summary Аутентификация пользователя
// @Description Вход в систему и получение JWT токена
// @Tags Аутентификация
// @Accept json
// @Produce json
// @Param credentials body models.LoginRequest true "Данные для входа"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/login [post]
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

package main

import (
	"backend/internal/database"
	"backend/internal/handlers"
	"backend/internal/middleware"
	"backend/internal/repository"
	"backend/internal/service/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// Параметры подключения к базе данных

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Подключаемся к базе данных
	db, err := database.ConnectToDB(host, port, user, password, dbname)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Инициализируем репозитории
	deviceRepo := repository.NewDeviceRepository(db)
	telemetryRepo := repository.NewTelemetryRepository(db)
	userRepo := repository.NewUserRepository(db)

	// Инициализируем сервисы
	deviceService := service.NewDeviceService(deviceRepo)
	telemetryService := service.NewTelemetryService(telemetryRepo)
	userService := service.NewUserService(userRepo)

	// Инициализируем обработчики
	deviceHandler := handlers.NewDeviceHandler(deviceService)
	telemetryHandler := handlers.NewTelemetryHandler(telemetryService)
	authHandler := handlers.NewAuthHandler(userService)

	// Настраиваем маршрутизацию
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})
	// Группа маршрутов с префиксом /api
	api := router.Group("/api")
	{
		// Группа маршрутов для устройств
		deviceRoutes := api.Group("/devices")
		deviceRoutes.Use(middleware.JwtMiddleware())
		{
			deviceRoutes.GET("", deviceHandler.GetDevices)
			deviceRoutes.GET("/:id", deviceHandler.GetDeviceByID)
			deviceRoutes.POST("", deviceHandler.AddDevice)
			deviceRoutes.PUT("/:id", deviceHandler.UpdateDevice)
			deviceRoutes.DELETE("/:id", deviceHandler.DeleteDevice)
		}

		// Группа маршрутов для телеметрии
		telemetryRoutes := api.Group("/devices/:id/telemetry").Use(middleware.JwtMiddleware())
		{
			telemetryRoutes.GET("", telemetryHandler.GetTelemetry)
			telemetryRoutes.POST("", telemetryHandler.AddTelemetry)
			telemetryRoutes.DELETE("/:telemetryId", telemetryHandler.DeleteTelemetry)
		}

		// Маршрут для аутентификации
		api.POST("/auth/login", authHandler.Login)
	}

	// Запускаем сервер
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

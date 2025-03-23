package main

import (
	"backend/internal/database"
	"backend/internal/handlers"
	"backend/internal/middleware"
	"backend/internal/repository"
	"backend/internal/service/service"

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

	// Группа маршрутов для устройств
	deviceRoutes := router.Group("/devices")
	deviceRoutes.Use(middleware.JwtMiddleware())
	{
		deviceRoutes.GET("", deviceHandler.GetDevices)
		deviceRoutes.GET("/:id", deviceHandler.GetDeviceByID)
		deviceRoutes.POST("", deviceHandler.AddDevice)
		deviceRoutes.PUT("/:id", deviceHandler.UpdateDevice)
		deviceRoutes.DELETE("/:id", deviceHandler.DeleteDevice)
	}

	// Группа маршрутов для телеметрии
	telemetryRoutes := router.Group("/devices/:id/telemetry").Use(middleware.JwtMiddleware())
	{
		telemetryRoutes.GET("", telemetryHandler.GetTelemetry)
		telemetryRoutes.POST("", telemetryHandler.AddTelemetry)
		telemetryRoutes.DELETE("/:telemetryId", telemetryHandler.DeleteTelemetry)
	}

	// Маршрут для аутентификации
	router.POST("/auth/login", authHandler.Login)

	// Запускаем сервер
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

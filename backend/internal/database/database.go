package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Импортируем драйвер PostgreSQL
)

// ConnectToDB создает подключение к базе данных PostgreSQL.
func ConnectToDB(host, port, user, password, dbname string) (*sql.DB, error) {
	// Формируем строку подключения
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Открываем подключение к базе данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Проверяем подключение
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

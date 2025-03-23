package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func JwtMiddleware() gin.HandlerFunc {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	var jwtKey = []byte(os.Getenv("SECRET_JWT_KEY"))

	return func(c *gin.Context) {
		var err = godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
		// Получаем заголовок Authorization
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Удаляем префикс "Bearer " из заголовка
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		claims := &Claims{}

		// Парсим токен
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Проверяем, что используется правильный алгоритм
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil {
			fmt.Println("Error parsing token:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		if !token.Valid {
			fmt.Println("Invalid token")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}

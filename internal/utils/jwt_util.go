package utils

import (
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

type JWTUtil struct {
	secretKey []byte
}

// NewJWTUtil создает экземпляр JWTUtil
func NewJWTUtil(secret string) JWTUtil {
	return JWTUtil{
		secretKey: []byte(secret),
	}
}

func (j *JWTUtil) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи токена
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		//return j.secretKey, nil
		return j.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Проверяем валидность токена и извлекаем claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Проверяем истечение срока действия токена
	if exp, ok := claims["exp"].(float64); ok {
		if time.Unix(int64(exp), 0).Before(time.Now()) {
			return nil, errors.New("token has expired")
		}
	} else {
		return nil, errors.New("invalid token structure")
	}

	return claims, nil
}

// GenerateToken создает JWT токен с указанным сроком действия
func (j *JWTUtil) GenerateToken(username string, expiration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(expiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secretKey)
}

// GenerateRefreshToken создает refresh-токен с указанным сроком действия
func (j *JWTUtil) GenerateRefreshToken(username string, expiration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(expiration).Unix(),
		"type":     "refresh",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secretKey)
}

// ValidateToken проверяет валидность токена и возвращает claims
func JSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json") // Установка заголовка Content-Type
	w.WriteHeader(statusCode)                          // Установка HTTP-статуса

	// Кодирование данных в JSON и отправка их в ResponseWriter
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}
}

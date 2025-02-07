package usecase

import (
	"errors"
	"motorcycle-sales/internal/domain/repositories"
	"motorcycle-sales/internal/utils"
	"time"
)

type AuthUseCase struct {
	userRepo *repositories.UserRepository
	jwtUtil  utils.JWTUtil
}

// NewAuthUseCase создает новый экземпляр AuthUseCase
func NewAuthUseCase(userRepo *repositories.UserRepository, util utils.JWTUtil) *AuthUseCase {
	return &AuthUseCase{
		userRepo: userRepo,
		jwtUtil:  util,
	}
}

func (a *AuthUseCase) GenerateToken(username string, duration time.Duration) (string, error) {
	return a.jwtUtil.GenerateToken(username, duration)
}

// RefreshToken обновляет токен доступа
func (uc *AuthUseCase) RefreshToken(refreshToken string) (string, error) {
	// Проверяем валидность refresh-токена
	claims, err := uc.jwtUtil.ValidateToken(refreshToken)
	if err != nil {
		return "", err
	}

	// Проверяем тип токена
	if claims["type"] != "refresh" {
		return "", errors.New("invalid token type")
	}

	// Генерируем новый токен
	username := claims["username"].(string)
	newToken, err := uc.jwtUtil.GenerateToken(username, time.Hour*1) // Новый токен на 1 час
	if err != nil {
		return "", err
	}

	return newToken, nil
}

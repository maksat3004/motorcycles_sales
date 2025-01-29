package services

import (
	"errors"
	"motorcycle-sales/internal/domain/models"
	"motorcycle-sales/internal/domain/repositories"
	"motorcycle-sales/internal/utils"
	"time"
)

type UserService struct {
	UserRepo  repositories.UserRepository
	userRepos repositories.UserRepositor
	JWTUtil   *utils.JWTUtil // Изменено на указатель для корректного вызова методов
}

// Регистрация пользователя
func (us *UserService) Register(user models.User) error {
	if _, err := us.userRepos.GetByUsername(user.Username); err == nil {
		return errors.New("user already exists")
	}

	// Создаем пользователя
	err := us.UserRepo.CreateUser(user)
	if err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

// Авторизация пользователя
func (us *UserService) Login(creds models.Credentials) (map[string]string, error) {
	// Получаем пользователя по имени
	user, err := us.userRepos.GetByUsername(creds.Username)
	if err != nil || user.PasswordHash != creds.Password {
		return nil, errors.New("invalid username or password")
	}

	// Генерация access-токена
	accessToken, err := us.JWTUtil.GenerateToken(user.Username, time.Minute*15)
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	// Генерация refresh-токена
	refreshToken, err := us.JWTUtil.GenerateToken(user.Username, time.Hour*24*7)
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	// Возврат токенов
	return map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil
}

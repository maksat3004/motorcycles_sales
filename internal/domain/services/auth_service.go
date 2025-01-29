// /internal/service/auth_service.go
package services

import (
	"errors"
	"motorcycle-sales/internal/domain/repositories"
	"motorcycle-sales/internal/utils"
	"time"
)

type AuthService struct {
	userRepo  repositories.UserRepository
	userRepos repositories.UserRepositor
	jwtUtil   utils.JWTUtil
}

func NewAuthService(userRepo repositories.UserRepository, jwtUtil utils.JWTUtil) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		jwtUtil:  jwtUtil,
	}
}

func (a *AuthService) Login(username, password string) (accessToken, refreshToken string, err error) {
	user, err := a.userRepos.GetByUsername(username)
	if err != nil || user.PasswordHash != password {
		return "", "", errors.New("invalid username or password")
	}

	accessToken, err = a.jwtUtil.GenerateToken(username, 15*time.Minute)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = a.jwtUtil.GenerateRefreshToken(username, 24*time.Hour)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (a *AuthService) ValidateAccessToken(token string) (string, error) {
	claims, err := a.jwtUtil.ValidateToken(token)
	if err != nil {
		return "", err
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", errors.New("invalid token payload")
	}

	return username, nil
}

func (a *AuthService) RefreshAccessToken(refreshToken string) (newAccessToken string, err error) {
	claims, err := a.jwtUtil.ValidateToken(refreshToken)
	if err != nil {
		return "", err
	}

	tokenType, ok := claims["type"].(string)
	if !ok || tokenType != "refresh" {
		return "", errors.New("invalid token type")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", errors.New("invalid token payload")
	}

	return a.jwtUtil.GenerateToken(username, 15*time.Minute)
}

package handlers

import (
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"motorcycle-sales/internal/domain/models"
	"motorcycle-sales/internal/domain/repositories"
	"motorcycle-sales/internal/utils"
	"net/http"
	"time"
)

type AuthHandler struct {
	UserRepo  repositories.UserRepository
	UserRepos repositories.UserRepositor
}

func NewAuthHandler(userRepo repositories.UserRepository) *AuthHandler {
	return &AuthHandler{UserRepo: userRepo}
}

func (h *AuthHandler) RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	// Пример логики обновления токена
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}

	newToken, err := h.NewAuthUseCase.RefreshToken(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	utils.JSONResponse(w, map[string]string{"token": newToken}, http.StatusOK)
}

// LoginHandler - обработчик для входа пользователя
func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Проверка существования пользователя
	user, err := h.UserRepo.FindByUsername(input.Username)
	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	// Проверка пароля
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Генерация JWT токенов
	jwtUtil := utils.NewJWTUtil("your_secret_key")

	accessToken, err := jwtUtil.GenerateToken(user.Username, 15*time.Minute)
	if err != nil {
		http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
		return
	}

	refreshToken, err := jwtUtil.GenerateToken(user.Username, 7*24*time.Hour) // Убрали GenerateRefreshToken
	if err != nil {
		http.Error(w, "Failed to generate refresh token", http.StatusInternalServerError)
		return
	}

	// Установка токенов в cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(15 * time.Minute),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		HttpOnly: true,
	})

	// Ответ пользователю
	utils.JSONResponse(w, map[string]string{"message": "Login successful"}, http.StatusOK)
}

// RegisterHandler - обработчик для регистрации пользователя
func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if input.Username == "" || input.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	// Проверка существования пользователя
	_, err := h.UserRepo.FindByUsername(input.Username)
	if err == nil {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	} else if !errors.Is(err, repositories.ErrNotFound) {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Хэширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	// Создание пользователя
	user := models.User{
		Username:     input.Username,
		PasswordHash: string(hashedPassword),
		Role:         input.Role,
	}

	if err := h.UserRepo.CreateUser(user); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, map[string]string{"message": "User registered successfully"}, http.StatusCreated)
}

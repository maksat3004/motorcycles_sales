package main

import (
	"fmt"
	"log"
	"net/http"

	"motorcycle-sales/internal/domain/handlers"
	"motorcycle-sales/internal/domain/repositories"
	"motorcycle-sales/internal/domain/usecase"
	"motorcycle-sales/internal/infrastructure/database"
	"motorcycle-sales/internal/infrastructure/middleware"
	"motorcycle-sales/internal/utils"
)

func main() {
	// Секретный ключ для JWT
	const jwtSecretKey = "qwerty"

	// Подключение к базе данных
	db, err := database.ConnectPostgres("postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	// Создание репозиториев
	userRepo := repositories.NewPostgresUserRepository(db) // Проверьте, что эта функция корректно создаёт репозиторий
	motorcycleRepo := repositories.NewPostgresMotorcycleRepository(db)
	orderRepo := repositories.NewPostgresOrderRepository(db)

	// Инициализация утилит
	jwtUtil := utils.NewJWTUtil(jwtSecretKey)

	// Инициализация UseCase слоев
	authUseCase := usecase.NewAuthUseCase(userRepo, jwtUtil) // Проверь, что userRepo подходит для usecase
	motorcycleUseCase := usecase.NewMotorcycleUseCase(motorcycleRepo)
	orderUseCase := usecase.NewOrderUseCase(orderRepo, motorcycleRepo)

	// Инициализация хендлеров
	authHandler := handlers.NewAuthHandler(authUseCase)
	motorcycleHandler := handlers.NewMotorcycleHandler(motorcycleUseCase)
	orderHandler := handlers.NewOrderHandler(orderUseCase)

	// Настройка маршрутов
	http.HandleFunc("/register", authHandler.RegisterHandler)    // Регистрация
	http.HandleFunc("/login", authHandler.LoginHandler)          // Логин
	http.HandleFunc("/refresh", authHandler.RefreshTokenHandler) // Обновление токена

	// Защищенные маршруты (требуют авторизации)
	http.HandleFunc("/motorcycles", middleware.AuthMiddleware(jwtUtil, motorcycleHandler.GetAllMotorcyclesHandler)) // Получение всех мотоциклов
	http.HandleFunc("/motorcycle", middleware.AuthMiddleware(jwtUtil, motorcycleHandler.AddMotorcycleHandler))      // Добавление мотоцикла
	http.HandleFunc("/orders", middleware.AuthMiddleware(jwtUtil, orderHandler.CreateOrderHandler))                 // Создание заказа

	// Запуск сервера
	serverAddr := ":8080"
	fmt.Printf("Сервер запущен на http://localhost%s\n", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}

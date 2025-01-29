package usecase

import (
	"motorcycle-sales/internal/domain/models"
	"motorcycle-sales/internal/domain/repositories"
)

type OrderUseCase struct {
	orderRepo      repositories.OrderRepository
	motorcycleRepo repositories.MotorcycleRepository
}

func NewOrderUseCase(orderRepo repositories.OrderRepository, motorcycleRepo repositories.MotorcycleRepository) OrderUseCase {
	return OrderUseCase{orderRepo: orderRepo, motorcycleRepo: motorcycleRepo}
}

func (u *OrderUseCase) CreateOrder(order models.Order) error {
	// Валидация или бизнес-логика для создания заказа
	_, err := u.motorcycleRepo.GetByID(order.MotorcycleID)
	if err != nil {
		return err // Ошибка, если мотоцикл не найден
	}

	return u.orderRepo.Create(order)
}

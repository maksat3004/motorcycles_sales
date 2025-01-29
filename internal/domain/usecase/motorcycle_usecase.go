package usecase

import (
	"motorcycle-sales/internal/domain/models"
	"motorcycle-sales/internal/domain/repositories"
)

type MotorcycleUseCase struct {
	motorcycleRepo repositories.MotorcycleRepository
}

func NewMotorcycleUseCase(repo repositories.MotorcycleRepository) MotorcycleUseCase {
	return MotorcycleUseCase{motorcycleRepo: repo}
}

func (u *MotorcycleUseCase) GetAllMotorcycles() ([]models.Motorcycle, error) {
	return u.motorcycleRepo.GetAll()
}

func (u *MotorcycleUseCase) AddMotorcycle(motorcycle models.Motorcycle) error {
	return u.motorcycleRepo.Add(motorcycle)
}

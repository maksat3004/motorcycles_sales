package repositories

import (
	"database/sql"
	"motorcycle-sales/internal/domain/models"
)

type MotorcycleRepository interface {
	GetAll() ([]models.Motorcycle, error)
	Create(motorcycle models.Motorcycle) error
	Add(motorcycle models.Motorcycle) error
	GetByID(id int) (models.Motorcycle, error)
}

type PostgresMotorcycleRepository struct {
	db *sql.DB
}

func NewPostgresMotorcycleRepository(db *sql.DB) *PostgresMotorcycleRepository {
	return &PostgresMotorcycleRepository{db: db}
}

func (r *PostgresMotorcycleRepository) GetAll() ([]models.Motorcycle, error) {
	rows, err := r.db.Query("SELECT id, name, brand, price, created_at FROM motorcycles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var motorcycles []models.Motorcycle
	for rows.Next() {
		var m models.Motorcycle
		if err := rows.Scan(&m.ID, &m.Name, &m.Brand, &m.Price, &m.CreatedAt); err != nil {
			return nil, err
		}
		motorcycles = append(motorcycles, m)
	}
	return motorcycles, nil
}

func (r *PostgresMotorcycleRepository) Add(motorcycle models.Motorcycle) error {
	_, err := r.db.Exec("INSERT INTO motorcycles (brand, model, price) VALUES ($1, $2, $3)", motorcycle.Brand, motorcycle.Model, motorcycle.Price)
	return err
}

func (r *PostgresMotorcycleRepository) GetByID(id int) (models.Motorcycle, error) {
	var m models.Motorcycle
	err := r.db.QueryRow("SELECT id, brand, model, price FROM motorcycles WHERE id = $1", id).Scan(&m.ID, &m.Brand, &m.Model, &m.Price)
	if err != nil {
		return models.Motorcycle{}, err
	}
	return m, nil
}

func (r *PostgresMotorcycleRepository) Create(motorcycle models.Motorcycle) error {
	_, err := r.db.Exec(
		"INSERT INTO motorcycles (name, brand, price) VALUES ($1, $2, $3)",
		motorcycle.Name, motorcycle.Brand, motorcycle.Price,
	)
	return err
}

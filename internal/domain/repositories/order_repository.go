package repositories

import (
	"database/sql"
	"motorcycle-sales/internal/domain/models"
)

type OrderRepository interface {
	Create(order models.Order) error
	GetByUserID(userID int) ([]models.Order, error)
}

type PostgresOrderRepository struct {
	db *sql.DB
}

func NewPostgresOrderRepository(db *sql.DB) OrderRepository {
	return &PostgresOrderRepository{db: db}
}

func (r *PostgresOrderRepository) Create(order models.Order) error {
	_, err := r.db.Exec(
		"INSERT INTO orders (user_id, motorcycle_id, total_price) VALUES ($1, $2, $3)",
		order.UserID, order.MotorcycleID, order.TotalPrice, order.MotorcycleID, order.Quantity,
	)
	return err
}

func (r *PostgresOrderRepository) GetByUserID(userID int) ([]models.Order, error) {
	rows, err := r.db.Query(
		"SELECT id, user_id, motorcycle_id, order_date, total_price FROM orders WHERE user_id = $1",
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var o models.Order
		if err := rows.Scan(&o.ID, &o.UserID, &o.MotorcycleID, &o.OrderDate, &o.TotalPrice); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}

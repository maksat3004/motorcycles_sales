package repositories

import (
	"database/sql"
	"errors"
	_ "errors"
	"motorcycle-sales/internal/domain/models"
)

var ErrNotFound = errors.New("not found")

type UserRepository struct {
	db *sql.DB
}
type UserRepositor interface {
	GetByUsername(username string) (models.User, error)
	Create(user models.User) error
	FindByUsername(username string) (models.User, error)
	CreateUser(user models.User) error
	ErrNotFound(err error)
	IsUserExists(username string) (bool, error)
	GetUserByUsername(username string) (*models.User, error)
}

func NewPostgresUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// FindByUsername ищет пользователя по имени
func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, password_hash, role FROM users WHERE username = $1`
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser создает нового пользователя
func (r *UserRepository) CreateUser(user models.User) error {
	query := `INSERT INTO users (username, password_hash, role) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, user.Username, user.PasswordHash, user.Role)
	if err != nil {
		return err
	}
	return nil
}

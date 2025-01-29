package repositories

import (
	"errors"
	_ "errors"
	"motorcycle-sales/internal/domain/models"

	"github.com/jmoiron/sqlx"
)

var ErrNotFound = errors.New("not found")

type UserRepository struct {
	db *sqlx.DB
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

func NewPostgresUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

// FindByUsername ищет пользователя по имени
func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, password_hash, role FROM users WHERE username = $1`
	err := r.db.Get(&user, query, username)
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

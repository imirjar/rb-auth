package storage

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/imirjar/rb-auth/internal/domain/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// UserStorage - репозиторий пользователей
type UserStorage struct {
	pool *pgxpool.Pool
}

func New(DBConn string) (*UserStorage, error) {
	// Создаем конфигурацию пула
	config, err := pgxpool.ParseConfig(DBConn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse connection string: %w", err)
	}

	// Настраиваем пул
	config.MaxConns = 25
	config.MinConns = 5
	config.MaxConnLifetime = 30 * time.Minute
	config.MaxConnIdleTime = 5 * time.Minute
	config.HealthCheckPeriod = 1 * time.Minute

	// Создаем пул
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Проверяем соединение
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &UserStorage{
		pool: pool,
	}, nil
}

// Close закрывает пул соединений
func (s *UserStorage) Close() {
	s.pool.Close()
}

func (s *UserStorage) AddUser(ctx context.Context, user models.User) error {
	query := `
        INSERT INTO users (
            email, 
            username, 
            password_hash, 
            is_active, 
            roles
        ) VALUES ($1, $2, $3, $4, $5)
        RETURNING 
            id, 
            email, 
            username, 
            password_hash, 
            is_active, 
            roles, 
            created_at, 
            updated_at, 
            last_login_at
    `

	var createdUser models.User
	var lastLoginAt *time.Time

	err := s.pool.QueryRow(ctx, query,
		user.Email,
		user.Username,
		user.PasswordHash,
		user.IsActive,
		user.Roles,
	).Scan(
		&createdUser.ID,
		&createdUser.Email,
		&createdUser.Username,
		&createdUser.PasswordHash,
		&createdUser.IsActive,
		&createdUser.Roles,
		&createdUser.CreatedAt,
		&createdUser.UpdatedAt,
		&lastLoginAt,
	)

	if err != nil {
		// Обработка ошибок уникальности (если email или username уже существуют)
		if strings.Contains(err.Error(), "unique") ||
			strings.Contains(err.Error(), "duplicate") {
			if strings.Contains(err.Error(), "email") {
				return fmt.Errorf("user with email %s already exists", user.Email)
			}
			if strings.Contains(err.Error(), "username") {
				return fmt.Errorf("user with username %s already exists", user.Username)
			}
			return fmt.Errorf("user with this data already exists")
		}
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (s *UserStorage) FindByEmail(ctx context.Context, email string) (models.User, error) {
	query := `
        SELECT 
            id, 
            email, 
            username, 
            password_hash, 
            is_active, 
            roles, 
            created_at, 
            updated_at, 
            last_login_at
        FROM users 
        WHERE email = $1
    `

	var user models.User
	var lastLoginAt *time.Time // pgx автоматически обрабатывает NULL

	err := s.pool.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.PasswordHash,
		&user.IsActive,
		&user.Roles,
		&user.CreatedAt,
		&user.UpdatedAt,
		&lastLoginAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user, fmt.Errorf("user with email %s not found", email)
		}
		return user, fmt.Errorf("failed to find user by email: %w", err)
	}

	user.LastLoginAt = lastLoginAt

	return user, nil
}

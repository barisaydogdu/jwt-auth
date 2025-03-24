package repository

import (
	"context"
	"fmt"
	"github.com/barisaydogdu/jwt-auth/domain"
	"github.com/jackc/pgx/v5"
	"log"
)

type UserRepository interface {
	FindByEmail(email string) (*domain.User, error)
	Create(user *domain.User) error
}

type userRepository struct {
	ctx context.Context
	DB  *pgx.Conn
}

func NewUserRepository(ctx context.Context, db *pgx.Conn) UserRepository {
	return &userRepository{ctx: ctx, DB: db}
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	user := &domain.User{}

	query := `
		SELECT id, firstname, lastname, username, email, password, role, created_at, updated_at 
		FROM users 
		WHERE email = $1
	`

	err := r.DB.QueryRow(r.ctx, query, email).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error fetching user by email: %w", err)
	}

	return user, nil
}

func (r *userRepository) Create(user *domain.User) error {
	query := `INSERT INTO users (firstname, lastname, username, email, password, role) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	if err := r.DB.QueryRow(r.ctx, query, user.FirstName, user.LastName, user.Username, user.Email, user.Password, user.Role).Scan(&user.ID); err != nil {
		return fmt.Errorf("error while creating user: %w", err)
	}
	log.Printf("user created successfully %d", user.ID)
	return nil
}

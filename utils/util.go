package utils

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"time"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type Utils interface {
	VerifyToken(string) (*jwt.Token, error)
	Cleanup()
}
type utils struct {
	db  *pgx.Conn
	ctx context.Context
}

func NewUtils(ctx context.Context, db *pgx.Conn) *utils {
	return &utils{
		ctx: ctx,
		db:  db,
	}
}

func (u *utils) CreateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().UTC().Add(time.Minute * 15).Unix(),
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (u *utils) VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}

func (u *utils) Cleanup() {
	_, err := u.db.Exec(u.ctx, "DELETE FROM users")
	if err != nil {
		log.Fatal(err)
	}
}

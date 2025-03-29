package service

import (
	"context"
	"github.com/barisaydogdu/jwt-auth/config"
	"github.com/barisaydogdu/jwt-auth/domain"
	"github.com/barisaydogdu/jwt-auth/infrastructure/postgres"
	repo "github.com/barisaydogdu/jwt-auth/repository"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var db *pgx.Conn
var ctx context.Context
var userRepo repo.UserRepository

func TestMain(m *testing.M) {
	Setup(m)
}

func Setup(m *testing.M) {
	var err error

	ctx = context.Background()

	err = godotenv.Load("../env/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envConfig, err := config.NewEnvDBConfig()
	if err != nil {
		log.Panic(err)
	}

	pg, err := postgres.NewPostgres(ctx, envConfig)
	if err != nil {
		log.Panic(err)
	}

	err = pg.ConnectDB()
	if err != nil {
		log.Panic(err)
	}

	db = pg.Conn

	userRepo = repo.NewUserRepository(ctx, db)

	m.Run()
}

func TestLogin(t *testing.T) {
	user := &domain.User{}

	user.FirstName = "baris"
	user.LastName = "aydogdu"
	user.Username = "john"
	user.Email = "baris@mail.com"
	user.Password = "123456"
	user.Role = "user"

	service := NewUserService(ctx, userRepo)

	registerToken, err := service.Register(user.FirstName, user.LastName, user.Username, user.Email, user.Password)
	if err != nil {
		t.Error(err)
	}

	t.Log(registerToken)

	loginToken, err := service.Login(user.Email, user.Password)
	if err != nil {
		t.Error(err)
	}

	t.Log(loginToken)

	assert.Equal(t, registerToken, loginToken)

	Cleanup()
}

func TestRegister(t *testing.T) {
	user := domain.User{}

	user.FirstName = "baris1"
	user.LastName = "aydogdu1"
	user.Username = "barisaydogdu2"
	user.Email = "baris1@mail.com"
	user.Password = "1234567"
	user.Role = "user"

	service := NewUserService(ctx, userRepo)

	registerToken, err := service.Register(user.FirstName, user.LastName, user.Username, user.Email, user.Password)
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, registerToken)

}

func Cleanup() {
	_, err := db.Exec(ctx, "DELETE FROM users")
	if err != nil {
		log.Fatal(err)
	}
}

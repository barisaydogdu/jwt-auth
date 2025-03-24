package repository

import (
	"context"
	config "github.com/barisaydogdu/jwt-auth/config"
	domain "github.com/barisaydogdu/jwt-auth/domain"
	postgres "github.com/barisaydogdu/jwt-auth/infrastructure/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var db *pgx.Conn
var ctx context.Context

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
		panic(err)
	}

	pg, err := postgres.NewPostgres(ctx, envConfig)
	if err != nil {
		panic(err)
	}

	err = pg.ConnectDB()
	if err != nil {
		panic(err)
	}
	db = pg.Conn
	m.Run()

}

func TestCreateUser(test *testing.T) {
	user := domain.User{}

	user.FirstName = "Baris"
	user.LastName = "Aydogdu"
	user.Email = "baris.aydogdu@gmail.com"
	user.Username = "barisaydogdu"
	user.Password = "123456"
	user.Role = "admin"

	repo := NewUserRepository(ctx, db)

	defer Cleanup()

	err := repo.Create(&user)
	if err != nil {
		test.Fatal(err)
	}
}

func TestFindUser(t *testing.T) {
	user := &domain.User{}

	user.FirstName = "Baris"
	user.LastName = "Aydogdu"
	user.Email = "barisaydogdu2@mail.com"
	user.Username = "barisaydogdu22"
	user.Password = "123456"
	user.Role = "admin"

	repo := NewUserRepository(ctx, db)

	err := repo.Create(user)
	if err != nil {
		t.Fatal(err)
	}

	user2, err := repo.FindByEmail("barisaydogdu2@mail.com")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, user2.Email, "barisaydogdu2@mail.com")

	Cleanup()
}

func Cleanup() {
	_, err := db.Exec(ctx, "DELETE FROM users")
	if err != nil {
		log.Fatal(err)
	}

}

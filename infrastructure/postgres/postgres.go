package postgres

import (
	"context"
	"fmt"
	"github.com/barisaydogdu/jwt-auth/config"
	"github.com/jackc/pgx/v5"
	"os"
)

type Postgres struct {
	Ctx    context.Context
	Config config.EnvDBConfig
}

func NewPostgres(ctx context.Context, config *config.EnvDBConfig) (*Postgres, error) {
	return &Postgres{
		Ctx:    ctx,
		Config: *config,
	}, nil
}

func (p *Postgres) ConnectDB() error {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", p.Config.User, p.Config.Password, p.Config.Host, p.Config.Port, p.Config.DBName)
	conn, err := pgx.Connect(p.Ctx, connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully connected to database")
	defer conn.Close(context.Background())

	return nil
}

package main

import (
	"context"
	"github.com/barisaydogdu/jwt-auth/config"
	"github.com/barisaydogdu/jwt-auth/infrastructure/postgres"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())

	err := godotenv.Load("../env/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	envConfig, err := config.NewEnvDBConfig()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pgSql, err := postgres.NewPostgres(ctx, envConfig)
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	
	err = pgSql.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
}

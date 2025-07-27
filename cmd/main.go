package main

import (
	"context"
	"fmt"
	"github.com/barisaydogdu/jwt-auth/config"
	"github.com/barisaydogdu/jwt-auth/handlers"
	"github.com/barisaydogdu/jwt-auth/infrastructure/postgres"
	"github.com/barisaydogdu/jwt-auth/repository"
	"github.com/barisaydogdu/jwt-auth/service"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file1")
	}
	envConfig, err := config.NewEnvDBConfig()
	if err != nil {
		log.Fatal("Error loading .env file2")
	}

	pgSql, err := postgres.NewPostgres(ctx, envConfig)
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	err = pgSql.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(ctx, pgSql.Conn)
	userService := service.NewUserService(ctx, userRepo)
	userHandler := handlers.NewUserHandler(ctx, userService)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/login", userHandler.Login)
	mux.HandleFunc("/api/register", userHandler.Register)

	fmt.Println("Listening on port " + envConfig.HttpPort)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})
	srv := &http.Server{
		Addr:         ":" + envConfig.HttpPort,
		Handler:      c.Handler(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Println("Starting server on " + envConfig.HttpPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := srv.Shutdown(ctxShutdown); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}

package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sashindionicus/DBDocument"
	"github.com/sashindionicus/DBDocument/pkg/config"
	"github.com/sashindionicus/DBDocument/pkg/handler"
	"github.com/sashindionicus/DBDocument/pkg/repository"
	"github.com/sashindionicus/DBDocument/pkg/service"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatalf("Error loading config file, err: %s", err.Error())
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.postgres.host"),
		Port:     viper.GetString("db.postgres.port"),
		Username: viper.GetString("db.postgres.username"),
		DBName:   viper.GetString("db.postgres.dbname"),
		SSLMode:  viper.GetString("db.postgres.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("failed to connect to db: %s", err.Error())
	}

	repositories := repository.NewRepositories(db)
	services := service.NewServices(repositories)
	handlers := handler.NewHandler(services)

	srv := DBDocument.NewServer()
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.Init()); err != nil {
			log.Printf("Error occurred while running server: %s\n", err.Error())
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		log.Printf("error occurred while shutting down http server: %s\n", err.Error())
	}
}

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"parser/internal/client"
	"parser/internal/config"
	"parser/internal/controller"
	"parser/internal/provider"
	"parser/internal/routes"
	"parser/internal/usecase/domain"
)

func main() {
	ctx := context.Background()

	// 1 конфигурация
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// 2 подключение к базе данных
	pgClient, err := client.NewPGClient(ctx, cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	// 3 инициализация провайдера
	pgProvider := provider.NewPgProvider(pgClient)

	// 4 инициализация domain
	domainUseCase := domain.NewDomainUseCase(pgProvider)

	// 5 инициализация контроллера
	domainController := controller.NewDomainController(domainUseCase)

	// 6 инициализация роутера
	router := routes.NewRouter(domainController)

	// 7 запуск сервера
	log.Printf("Server starting on port %s", cfg.ServerPort)
	err = http.ListenAndServe(fmt.Sprintf(":%s", cfg.ServerPort), router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

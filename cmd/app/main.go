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

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	pgClient, err := client.NewPGClient(ctx, cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	pgProvider := provider.NewPgProvider(pgClient)
	domainUseCase := domain.NewDomainUseCase(pgProvider)
	domainController := controller.NewDomainController(domainUseCase)

	router := routes.NewRouter(domainController)

	log.Printf("Server starting on port %s", cfg.ServerPort)
	err = http.ListenAndServe(fmt.Sprintf(":%s", cfg.ServerPort), router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

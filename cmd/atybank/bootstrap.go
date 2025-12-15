package main

import (
	"database/sql"

	"github.com/go-chi/chi/v5"

	"atybank/cmd/atybank/handlers"
	"atybank/cmd/atybank/repositories"
	"atybank/cmd/atybank/routes"
	"atybank/cmd/atybank/services"
	configRepository "atybank/internal/infrastructure/config"
)

func Bootstrap(db *sql.DB, cfg configRepository.Config) (*chi.Mux, error) {
	// Repositories
	wiredRepositories, err := repositories.New(db)
	if err != nil {
		return nil, err
	}

	// Services
	wiredServices, err := services.New(wiredRepositories)
	if err != nil {
		return nil, err
	}

	// Handlers
	httpHandlers, err := handlers.New(db, cfg, wiredServices)
	if err != nil {
		return nil, err
	}

	return routes.New(cfg, *wiredServices, httpHandlers)
}

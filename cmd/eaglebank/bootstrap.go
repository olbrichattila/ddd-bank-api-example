package main

import (
	"database/sql"

	"github.com/go-chi/chi/v5"

	"eaglebank/cmd/eaglebank/handlers"
	"eaglebank/cmd/eaglebank/repositories"
	"eaglebank/cmd/eaglebank/routes"
	"eaglebank/cmd/eaglebank/services"
	configRepository "eaglebank/internal/infrastructure/config"
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

package main

import (
	"database/sql"
	"fmt"

	"github.com/go-chi/chi/v5"

	"eaglebank/cmd/eaglebank/handlers"
	"eaglebank/cmd/eaglebank/repositories"
	"eaglebank/cmd/eaglebank/services"
	"eaglebank/internal/api/middleware"
	configRepository "eaglebank/internal/infrastructure/config"
)

const (
	apiVersion = "v1"
)

func NewRouter(db *sql.DB, cfg configRepository.Config) (*chi.Mux, error) {
	r := chi.NewRouter()
	r.Use(middleware.JSON())

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
	httpHandlers, err := handlers.New(wiredServices)
	if err != nil {
		return nil, err
	}

	// Public endpoints
	r.Post(routeVersionDecorator("users"), httpHandlers.User.Create)

	// Authenticated routes
	r.Group(func(auth chi.Router) {
		auth.Use(middleware.JWT(cfg.GetJWTSecret()))

		// User routes
		auth.Group(func(auth chi.Router) {
			auth.Use(middleware.UserRole())
			auth.Get(routeVersionDecorator("users/{userId}"), httpHandlers.User.Get)
			auth.Patch(routeVersionDecorator("users/{userId}"), httpHandlers.User.Update)
			auth.Delete(routeVersionDecorator("users/{userId}"), httpHandlers.User.Delete)
		})

		// Account routes
		auth.Post(routeVersionDecorator("accounts"), httpHandlers.Account.Create)
		auth.Get(routeVersionDecorator("accounts"), httpHandlers.Account.List)
		auth.Group(func(auth chi.Router) {
			auth.Use(middleware.AccountRole())
			auth.Get(routeVersionDecorator("accounts/{accountNumber}"), httpHandlers.Account.Get)
			auth.Patch(routeVersionDecorator("accounts/{accountNumber}"), httpHandlers.Account.Update)
			auth.Delete(routeVersionDecorator("accounts/{accountNumber}"), httpHandlers.Account.Delete)
		})

		// Transaction routes
		auth.Post(routeVersionDecorator("accounts/{accountNumber}/transactions"), httpHandlers.Transaction.Create)
		auth.Get(routeVersionDecorator("accounts/{accountNumber}/transactions"), httpHandlers.Transaction.List)
		auth.Get(routeVersionDecorator("accounts/{accountNumber}/transactions/{transactionNumber}"), httpHandlers.Transaction.Get)
	})

	return r, nil
}

func routeVersionDecorator(route string) string {
	return fmt.Sprintf("/%s/%s", apiVersion, route)
}

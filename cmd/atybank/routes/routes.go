package routes

import (
	"fmt"

	"atybank/cmd/atybank/handlers"
	"atybank/cmd/atybank/services"
	"atybank/internal/api/middleware"
	configRepository "atybank/internal/infrastructure/config"

	"github.com/go-chi/chi/v5"
)

const (
	apiVersion = "v1"
)

func New(
	cfg configRepository.Config,
	services services.Services,
	httpHandlers *handlers.Handlers,
) (*chi.Mux, error) {
	r := chi.NewRouter()
	r.Use(middleware.JSON())

	// Public endpoints
	r.Post(routeVersionDecorator("login"), httpHandlers.User.Login)
	r.Post(routeVersionDecorator("users"), httpHandlers.User.Create)
	r.Get("/health", httpHandlers.Health.Health)
	r.Get("/ready", httpHandlers.Health.Ready)

	// Authenticated routes
	r.Group(func(auth chi.Router) {
		auth.Use(middleware.JWT(cfg.GetJWTSecret()))

		// User routes
		auth.Group(func(auth chi.Router) {
			auth.Use(middleware.UserGuard(services.Logger))
			auth.Get(routeVersionDecorator("users/{userId}"), httpHandlers.User.Get)
			auth.Patch(routeVersionDecorator("users/{userId}"), httpHandlers.User.Update)
			auth.Delete(routeVersionDecorator("users/{userId}"), httpHandlers.User.Delete)
		})

		// Account routes
		auth.Post(routeVersionDecorator("accounts"), httpHandlers.Account.Create)
		auth.Get(routeVersionDecorator("accounts"), httpHandlers.Account.List)
		auth.Group(func(auth chi.Router) {
			auth.Use(middleware.AccountGuard(services.Account))
			auth.Get(routeVersionDecorator("accounts/{accountNumber}"), httpHandlers.Account.Get)
			auth.Patch(routeVersionDecorator("accounts/{accountNumber}"), httpHandlers.Account.Update)
			auth.Delete(routeVersionDecorator("accounts/{accountNumber}"), httpHandlers.Account.Delete)
		})

		// Transaction routes
		auth.Group(func(auth chi.Router) {
			auth.Use(middleware.AccountGuard(services.Account))
			auth.Post(routeVersionDecorator("accounts/{accountNumber}/transactions"), httpHandlers.Transaction.Create)
			auth.Get(routeVersionDecorator("accounts/{accountNumber}/transactions"), httpHandlers.Transaction.List)

			auth.Group(func(auth chi.Router) {
				auth.Use(middleware.TransactionGuard(services.Transaction))
				auth.Get(routeVersionDecorator("accounts/{accountNumber}/transactions/{transactionNumber}"), httpHandlers.Transaction.Get)
			})
		})

	})

	return r, nil
}

func routeVersionDecorator(route string) string {
	return fmt.Sprintf("/%s/%s", apiVersion, route)
}

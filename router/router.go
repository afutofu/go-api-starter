package router

import (
	"github.com/afutofu/go-api-starter/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"
)

func SetupRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Post("/register", handlers.Register)
	router.Post("/login", handlers.Login)
	router.Post("/logout", handlers.Logout)

	log.Info("Router setup complete")

	return router
}

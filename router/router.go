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

	router.Post("/todos", handlers.CreateTodo)
	router.Get("/todos", handlers.GetTodos)
	router.Get("/todos/{id}", handlers.GetTodo)
	router.Put("/todos/{id}", handlers.UpdateTodo)
	router.Delete("/todos/{id}", handlers.DeleteTodo)

	log.Info("Router setup complete")

	return router
}

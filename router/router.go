package router

import (
	"github.com/afutofu/go-api-starter/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	log "github.com/sirupsen/logrus"
)

func SetupRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// CORS middleware setup
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Replace with your frontend URL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	router.Use(cors.Handler)

	router.Post("/todos", handlers.CreateTodo)
	router.Get("/todos", handlers.GetTodos)
	router.Get("/todos/{id}", handlers.GetTodo)
	router.Put("/todos/{id}", handlers.UpdateTodo)
	router.Delete("/todos/{id}", handlers.DeleteTodo)

	log.Info("Router setup complete")

	return router
}

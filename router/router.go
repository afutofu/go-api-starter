package router

import (
	"net/http"

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

	// Serve OpenAPI YAML file
	router.Get("/docs/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/openapi.yaml")
	})

	// Serve Swagger UI files
	router.Get("/swagger/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swagger-ui-dist"))).ServeHTTP(w, r)
	})

	// Serve modified index.html
	router.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./swagger-ui-dist/index.html")
	})
	log.Info("Router setup complete")

	return router
}

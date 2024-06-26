package router

import (
	"net/http"

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

	router.Post("/register", handlers.Register)
	router.Post("/login", handlers.Login)
	router.Post("/logout", handlers.Logout)

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

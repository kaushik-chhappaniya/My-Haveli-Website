package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kaushik-chhappnaiya/myHaweli/internal/handlers"
	"github.com/kaushik-chhappnaiya/myHaweli/internal/middleware"
)

func RegisterRoutes(app *handlers.App) http.Handler {
	// Define your route registrations here
	mux := chi.NewRouter()
	mux.Use(middleware.LoggingMiddleware)

	// Static files (adjust path to your static dir)
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// Routes
	mux.HandleFunc("/", app.IndexPageHandler)
	mux.HandleFunc("/about-us", app.AboutPageHandler)

	// Admin routes
	mux.HandleFunc("/admin", app.AdminHandler)
	mux.HandleFunc("/admin/dashboard", app.AdminDashboardHandler)

	// Custom 404 handler - this should be the last route
	mux.NotFound(app.NotFoundHandler)

	return mux
}

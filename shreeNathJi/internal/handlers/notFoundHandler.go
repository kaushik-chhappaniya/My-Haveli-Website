package handlers

import (
	"net/http"

	logger "github.com/kaushik-chhappnaiya/myHaweli/internal/middleware/logger"
)

// NotFoundHandler handles 404 errors with a custom page
func (a *App) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	logger.Warnf("404 Not Found: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)

	// Set 404 status code
	w.WriteHeader(http.StatusNotFound)

	// Prepare template data
	data := map[string]any{
		"title":       "Page Not Found - श्री नाथजी हवेली",
		"currentPath": r.URL.Path,
		"method":      r.Method,
	}

	// Render the 404 template
	name, tmpl := a.Render("notFound", data)
	if err := tmpl.ExecuteTemplate(w, name, data); err != nil {
		logger.Errorf("Error rendering 404 template: %v", err)
		// Fallback to basic HTTP error if template fails
		http.Error(w, "404 - Page Not Found", http.StatusNotFound)
		return
	}

	logger.Debugf("Rendered 404 page for path: %s", r.URL.Path)
}

// Custom404Middleware wraps handlers to catch 404s
func Custom404Middleware(notFoundHandler http.HandlerFunc) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Create a response writer that captures the status code
			rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

			// Call the next handler
			next.ServeHTTP(rw, r)

			// If no route was found (status is still 200 and no content was written)
			// or if it was explicitly set to 404, call our custom handler
			if rw.statusCode == http.StatusNotFound {
				notFoundHandler(w, r)
			}
		})
	}
}

// responseWriter wraps http.ResponseWriter to capture status codes
type responseWriter struct {
	http.ResponseWriter
	statusCode int
	written    bool
}

func (rw *responseWriter) WriteHeader(code int) {
	if !rw.written {
		rw.statusCode = code
		rw.written = true
		rw.ResponseWriter.WriteHeader(code)
	}
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	if !rw.written {
		rw.WriteHeader(http.StatusOK)
	}
	return rw.ResponseWriter.Write(b)
}

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/kaushik-chhappnaiya/myHaweli/internal/handlers"
	logger "github.com/kaushik-chhappnaiya/myHaweli/middleware/logger"
	"github.com/kaushik-chhappnaiya/myHaweli/routes"
)

func init() {
	// Any necessary initialization can be done here
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		logger.Warn("Error loading .env file, using system environment variables")
	}

}
func main() {

	fmt.Println("!! Jai Shree Nath Ji !!")
	// r := chi.NewRouter()

	app := &handlers.App{
		TemplateDir: "./ui/templates",
	}

	// Parse base + partials once
	app.MustLoadBase()

	// Precompute clones per page (no per-request parsing)
	app.MustPrecomputePages()

	srv := &http.Server{
		Addr:         ":8081",
		Handler:      routes.RegisterRoutes(app),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		logger.Error(err.Error())
	}

	// chiRouter := app.Routes()

	// 1. Setting up middlewares
	// Custom debug logger to indicate server start
	// chiRouter.Use(customMiddleware.LoggingMiddleware)
	logger.Info("Server started")

}

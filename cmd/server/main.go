package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mrityunjayvashisth/mini-api-demo/generated/categories"
	"github.com/mrityunjayvashisth/mini-api-demo/generated/todos"
	"github.com/mrityunjayvashisth/mini-api-demo/handlers"
)

func main() {
	// Parse command-line flags
	port := flag.Int("port", 8080, "Server port")
	flag.Parse()

	// Create Echo instance
	e := echo.New()

	// Add middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Register API handlers
	todoHandler := handlers.NewTodoHandler()
	todos.RegisterHandlers(e, todoHandler)

	categoryHandler := handlers.NewCategoryHandler()
	categories.RegisterHandlers(e, categoryHandler)

	// Add health check endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	// Add OpenAPI documentation endpoints
	e.GET("/api-docs/todos", func(c echo.Context) error {
		return c.File("api-specs/todos-api.yaml")
	})

	e.GET("/api-docs/categories", func(c echo.Context) error {
		return c.File("api-specs/categories-api.yaml")
	})

	// Start server
	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Starting server on %s", addr)
	if err := e.Start(addr); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}
}

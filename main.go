package main

import (
	"github.com/labstack/echo/v4"
	"webapp/pkg/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Set up logging with JSON format for Loki
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)

	// Get port from environment variable, default to 3000
	port := config.GetEnv("PORT", "3000")

	// Initialize Echo
	e := echo.New()
	e.Static("/", "public")

	// Define root route
	e.GET("/", func(c echo.Context) error {
		log.WithFields(log.Fields{
			"method": c.Request().Method,
			"path":   c.Request().URL.Path,
		}).Info("Serving webapp.html")
		return c.File("public/views/webapp.html")
	})

	// Start the server
	log.WithField("port", port).Info("Starting web application")
	if err := e.Start(":" + port); err != nil {
		log.WithError(err).Fatal("Failed to start server")
	}
}

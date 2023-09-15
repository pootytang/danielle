package main

import (
	"github.com/labstack/echo/v4"
	"github.com/pootytang/danielleapi/handlers"
)

func setupRoutes(app *echo.Echo) {
	public_routes := app.Group("/api/v1/public")
	private_routes := app.Group("/api/v1/private")

	// BASE ROUTES
	app.GET("/api/v1", handlers.ShowAllRoutes) // Show all routes. Maybe send a redirect to the public routes

	// PUBLIC ROUTES
	public_routes.GET("/", handlers.ShowPublicRoutes)
	public_routes.GET("/ultrasounds", handlers.GetUltrasounds)
	public_routes.GET("/ultrasounds/:image", handlers.GetUltrasoundImage)
	public_routes.GET("/growingmommy", handlers.GetGrowingMommyImages)
	public_routes.GET("/growingmommy/:image", handlers.GetGrowingMommyImage)
	public_routes.GET("/shower_reveal", handlers.GetShowerRevealImages)
	public_routes.GET("/shower_reveal/:image", handlers.GetShowerRevealImage)

	// PRIVATE ROUTES
	private_routes.GET("/", handlers.ShowPrivateRoutes)
}

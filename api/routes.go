package main

import (
	"github.com/labstack/echo/v4"
	"github.com/pootytang/danielleapi/handlers"
	"github.com/pootytang/danielleapi/initializers"
	"github.com/pootytang/danielleapi/middleware"
)

func setupRoutes(app *echo.Echo) {
	/**** BASE ROUTES ****/
	app.GET("/api/v1", handlers.ShowAllRoutes) // Show all routes. Maybe send a redirect to the public routes

	/**** PUBLIC ROUTES ****/
	public_routes := app.Group("/api/v1/public")
	public_routes.GET("/", handlers.ShowPublicRoutes)
	public_routes.GET("/ultrasounds", handlers.GetUltrasounds)
	public_routes.GET("/ultrasounds/:image", handlers.GetUltrasoundImage)
	public_routes.GET("/growingmommy", handlers.GetGrowingMommyImages)
	public_routes.GET("/growingmommy/:image", handlers.GetGrowingMommyImage)
	public_routes.GET("/shower_reveal", handlers.GetShowerRevealImages)
	public_routes.GET("/shower_reveal/:image", handlers.GetShowerRevealImage)
	public_routes.GET("/birth_day", handlers.GetBirthDayImages)
	public_routes.GET("/dob/:image", handlers.GetBirthDayImage)

	// Stages
	public_routes.GET("/zero-one", handlers.GetInfantImages)
	public_routes.GET("/zero-one/:image", handlers.GetInfantImage)
	public_routes.GET("/one-two", handlers.GetToddlerImages)
	public_routes.GET("/one-two/:image", handlers.GetToddlerImage)
	public_routes.GET("/three-six", handlers.GetEarlyChildhoodImages)
	public_routes.GET("/three-six/:image", handlers.GetEarlyChildhoodImage)
	public_routes.GET("/seven-ten", handlers.GetLateChildhoodImages)
	public_routes.GET("/seven-ten/:image", handlers.GetLateChildhoodImage)
	public_routes.GET("/eleven-nineteen", handlers.GetAdolescenceImages)
	public_routes.GET("/eleven-nineteen/:image", handlers.GetAdolescenceImage)
	public_routes.GET("/twenty-fortyfour", handlers.GetEarlyAdulthoodImages)
	public_routes.GET("/twenty-fortyfour/:image", handlers.GetEarlyAdulthoodImage)
	public_routes.GET("/fortyfive-sixtyfour", handlers.GetMiddleAdulthoodImages)
	public_routes.GET("/fortyfive-sixtyfour/:image", handlers.GetMiddleAdulthoodImage)
	public_routes.GET("/sixtyfiveplus", handlers.GetLateAdulthoodImages)
	public_routes.GET("/sixtyfiveplus/:image", handlers.GetLateAdulthoodImage)

	/**** PRIVATE ROUTES ****/
	private_routes := app.Group("/api/v1/private")
	private_routes.Use(middleware.DeserializeUser)
	private_routes.GET("/", handlers.ShowPrivateRoutes)
	private_routes.GET("/delivery", handlers.GetDeliveryImages)
	private_routes.GET("/delivery/:image", handlers.GetDeliveryImage)

	/**** AUTH ROUTES ****/
	auth_routes := app.Group("/api/v1/auth")
	ac := handlers.NewAuthController(initializers.DB)
	auth_routes.POST("/register", ac.RegisterUser)
	auth_routes.POST("/login", ac.SignInUser)
	auth_routes.POST("/refresh", ac.RefreshAccessToken)
	auth_routes.POST("/user", ac.GetUser)
	auth_routes.GET("/user", ac.FindUser)
	auth_routes.GET("/logout", ac.LogoutUser, middleware.DeserializeUser)
}

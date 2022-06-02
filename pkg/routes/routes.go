package routes

import (
	"github.com/gofiber/fiber/v2"

	"shopping-mono/app/controllers"
	"shopping-mono/pkg/middlewares"
)

func SetupRoutes(c *controllers.Controller, app *fiber.App, middleware *middlewares.Middleware) {
	v1 := app.Group("/api/v1")
	// user
	{
		v1.Post("/user", c.CreateUser)
		v1.Get("/user/:id", middleware.JWTProtected, c.GetUserById)
		v1.Get("/users", middleware.JWTProtected, c.ListUsers)
		v1.Get("/users/all", middleware.JWTProtected, c.ListAllUsers)
		v1.Put("/user/:id", middleware.JWTProtected, c.UpdateUserById)
		v1.Delete("/user/:id", middleware.JWTProtected, c.DeleteUserById)
	}
	// auth
	{
		v1.Post("/auth/login", c.Login)
		v1.Get("/auth/profile", middleware.JWTProtected, c.GetProfile)
		v1.Post("/auth/refresh", middleware.JWTRefreshProtected, c.Refresh)
	}
}

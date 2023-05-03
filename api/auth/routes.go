package authRoutes

import (
	"github.com/gofiber/fiber/v2"
	authControllers "github.com/sarawutwn/gofiber-template/api/auth/controllers"
	"github.com/sarawutwn/gofiber-template/middleware"
)

func SetupAuthRoutes(router fiber.Router) {
	app := router.Group("auth")
	app.Get("/get-all", middleware.AuthorizationRequired(), authControllers.GetAll)
	app.Post("/sign-up", authControllers.SignUp)
	app.Post("/sign-in", authControllers.SignIn)
	app.Get("/get-profile", middleware.AuthorizationRequired(), authControllers.GetProfile)
	app.Get("/GetUserIdByToken", middleware.AuthorizationRequired(), authControllers.GetUserIdByToken)
}

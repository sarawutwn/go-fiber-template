package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	authRoutes "github.com/sarawutwn/gofiber-template/api/auth"
	permissionRoutes "github.com/sarawutwn/gofiber-template/api/permission"
	roleRoutes "github.com/sarawutwn/gofiber-template/api/role"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("api", logger.New())

	// authRoutes.SetupAuthRoutes(api)
	authRoutes.SetupAuthRoutes(api)
	roleRoutes.SetupRoleRoutes(api)
	permissionRoutes.SetupPermissionRoutes(api)
}

package roleRoutes

import (
	"github.com/gofiber/fiber/v2"
	roleController "github.com/sarawutwn/gofiber-template/api/role/controllers"
	"github.com/sarawutwn/gofiber-template/middleware"
)

func SetupRoleRoutes(router fiber.Router) {
	app := router.Group("role")
	app.Get("/get-by-id/:role_id", middleware.AuthorizationRequired(), roleController.GetById)
	app.Get("/get-all", middleware.AuthorizationRequired(), roleController.GetAll)
	app.Post("/create", middleware.AuthorizationRequired(), roleController.CreateRole)
}

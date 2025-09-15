package router

import (
	"github.com/Sambit99/Basic-HRMS-GO-MongoDB/pkg/controller"
	"github.com/gofiber/fiber/v2"
)

var RegisterEmployeeRoutes = func(app *fiber.App) {
	employeeRouter := app.Group("/emp")

	employeeRouter.Get("/", controller.GetAllEmployee)
	employeeRouter.Post("/", controller.NewEmployee)
	employeeRouter.Get("/:id", controller.GetEmployee)
	employeeRouter.Delete("/:id", controller.DeleteEmployee)
	employeeRouter.Put("/:id", controller.UpdateEmployee)
}

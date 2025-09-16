package main

import (
	"fmt"

	"github.com/Sambit99/Basic-HRMS-GO-MongoDB/pkg/router"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	router.RegisterEmployeeRoutes(app)

	fmt.Println("🚀 Server running on port 8080")
	app.Listen(":8080")
}

package controller

import (
	"github.com/Sambit99/Basic-HRMS-GO-MongoDB/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func GetAllEmployee(c *fiber.Ctx) error {
	employees := service.GetEmployees()

	c.Set("Content-Type", "application/json")

	if len(employees) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":     "success",
			"statuscode": 200,
			"message":    "No records found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"statuscode": 200,
		"message":    "All records fetched successfully",
		"data":       employees,
	})
}

func GetEmployee(c *fiber.Ctx) error {
	return nil
}

func NewEmployee(c *fiber.Ctx) error {
	return nil
}

func DeleteEmployee(c *fiber.Ctx) error {
	return nil
}

func UpdateEmployee(c *fiber.Ctx) error {
	return nil
}

package controller

import (
	"github.com/Sambit99/Basic-HRMS-GO-MongoDB/pkg/model"
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
	id := c.Params("id")

	employee := service.GetEmployee(id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"statuscode": 200,
		"message":    "Request completed",
		"data":       employee,
	})
}

func NewEmployee(c *fiber.Ctx) error {
	var employee model.NewEmployeeDto

	if err := c.BodyParser(&employee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": "Error while parsing body",
			"error":   err.Error(),
		})
	}

	id := service.CreateEmployee(employee)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"id": id,
		},
	})
}

func DeleteEmployee(c *fiber.Ctx) error {
	return nil
}

func UpdateEmployee(c *fiber.Ctx) error {
	return nil
}

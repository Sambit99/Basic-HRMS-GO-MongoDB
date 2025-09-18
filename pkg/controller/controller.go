package controller

import (
	"log"

	"github.com/Sambit99/Basic-HRMS-GO-MongoDB/pkg/model"
	"github.com/Sambit99/Basic-HRMS-GO-MongoDB/pkg/service"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
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
	id := c.Params("id")

	success := service.DeleteEmployee(id)

	if success {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"status":     "success",
			"statuscode": 204,
			"message":    "Request completed",
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"status":     "failed",
		"statuscode": 500,
		"message":    "Error while deleting Employee",
	})
}

func UpdateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	var updatedEmployee model.UpdateEmployeeDto

	if err := c.BodyParser(&updatedEmployee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": "Error while parsing body",
			"error":   err.Error(),
		})
	}

	parsedId, err := bson.ObjectIDFromHex(id)

	if err != nil {
		log.Fatal("Error while converting Employee Id, From Hex to Object ID", err.Error())
	}

	updatedEmployee.ID = parsedId

	isUpdated := service.UpdateEmployee(updatedEmployee)

	if !isUpdated {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "failed",
			"statuscode": 500,
			"message":    "Error while updating Employee",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"statuscode": 200,
		"message":    "Employee updated",
	})
}

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

func createCustomer(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse request body into Customer struct
		customer := new(Customer)
		if err := c.BodyParser(customer); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unable to parse request body"})
		}

		// Check if a customer with the same name already exists
		var existingCustomer Customer
		if err := db.Where("name = ?", customer.Name).First(&existingCustomer).Error; err == nil {
			// Customer already exists, return a conflict response
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Customer already exists"})
		}

		// Customer does not exist, create a new record
		db.Create(&customer)

		// Return success response
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "success", "data": customer})
	}
}

func updateCustomer(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		customer := new(Customer)
		if err := c.BodyParser(customer); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unable to parse request body"})
		}

		id := c.Params("ID")
		var existingCustomer Customer
		if err := db.First(&existingCustomer, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Customer not found"})
		}

		db.Model(&existingCustomer).Updates(customer)
		return c.Status(fiber.StatusOK).JSON(existingCustomer)
	}
}

func deleteCustomer(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("ID")
		var customer Customer
		if err := db.First(&customer, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Customer not found"})
		}

		db.Delete(&customer)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Customer deleted successfully"})
	}
}

func getCustomer(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("ID")
		var customer Customer
		if err := db.First(&customer, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Customer not found"})
		}

		return c.Status(fiber.StatusOK).JSON(customer)
	}
}

func getAllCustomers(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Query all customers from the databases
		var customers []Customer
		if err := db.Find(&customers).Error; err != nil {
			// If error occurs during query, return 500 status with JSON error message
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		// Return customers as JSON response
		return c.Status(fiber.StatusOK).JSON(customers)
	}
}

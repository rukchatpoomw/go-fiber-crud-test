package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

func registerRoutes(app *fiber.App, db *gorm.DB) {
	app.Post("/customers", createCustomer(db))
	app.Put("/customers/:id", updateCustomer(db))
	app.Delete("/customers/:id", deleteCustomer(db))
	app.Get("/customers/:id", getCustomer(db))
}

func createCustomer(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		customer := new(Customer)
		if err := c.BodyParser(customer); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unable to parse request body"})
		}

		db.Create(&customer)
		return c.Status(fiber.StatusCreated).JSON(customer)
	}
}

func updateCustomer(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		customer := new(Customer)
		if err := c.BodyParser(customer); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unable to parse request body"})
		}

		id := c.Params("id")
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
		id := c.Params("id")
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
		id := c.Params("id")
		var customer Customer
		if err := db.First(&customer, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Customer not found"})
		}

		return c.Status(fiber.StatusOK).JSON(customer)
	}
}

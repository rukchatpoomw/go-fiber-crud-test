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
	app.Get("/customers", getAllCustomers(db))
}

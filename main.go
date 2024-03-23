package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app := fiber.New()

	// SQLite database connection
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Customer{})

	// Register routes
	registerRoutes(app, db)

	// Start server
	app.Listen(":3000")
}

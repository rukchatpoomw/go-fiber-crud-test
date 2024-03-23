package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// testsssasss
	app := fiber.New()

	// SQLite database connection with GORM (ORM library)
	db, err := initializeDB()

	if err != nil {
		println(err.Error())
	}

	// Migrate the schema
	autoMigrate(db)

	// Create initial customer data
	createInitialData(db)

	// Register routes
	registerRoutes(app, db)

	// Start server
	app.Listen(":3000")
}

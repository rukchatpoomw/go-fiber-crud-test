package main

import "github.com/jinzhu/gorm"

// Initialize database connection
func initializeDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		// Disconnect the database
		defer db.Close()
		return nil, err
	}
	return db, nil
}

// Auto-migrate the database schema
func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Customer{})
}

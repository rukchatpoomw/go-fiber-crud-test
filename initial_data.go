package main

import (
	"log"

	"github.com/jinzhu/gorm"
)

// Create initial customer data in the database
func createInitialData(db *gorm.DB) {
	// Count the total number of customers in the database
	var count int
	db.Model(&Customer{}).Count(&count)

	// Check if there are more than two documents
	if count > 2 {
		log.Println("Skipping initial data creation as there are already documents")
		return
	}

	// Insert initial customer data only if there are less than or equal to two documents
	db.Create(&Customer{Name: "John Doe", Age: 30})
	db.Create(&Customer{Name: "Jane Smith", Age: 25})
	db.Create(&Customer{Name: "Rachel Miki", Age: 18})
	log.Println("Initial data created successfully")
}

package main

type Customer struct {
	ID   uint `gorm:"primary_key"`
	Name string
	Age  int
}

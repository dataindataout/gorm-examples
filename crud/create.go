package main

// using sample data set from https://docs.yugabyte.com/preview/sample-data/retail-analytics/

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	ID        uint
	Category  string
	CreatedAt time.Time
	Ean       string
	Price     float32
	Rating    float32
	Title     string
	Vendor    string
}

func main() {
	dsn := "host=localhost user=yugabyte password=yugabyte dbname=yb_demo port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	product := Product{ID: 205, Category: "Doohicke1", CreatedAt: time.Now(), Ean: "1078766578568", Price: 23.45, Rating: 4.0, Title: "Fancy Blue Hat", Vendor: "Hattier-Braden"}

	result := db.Create(&product) // pass pointer of data to Create

	if result.Error != nil && result.RowsAffected != 1 {
		panic("failed to create record")
	}
}

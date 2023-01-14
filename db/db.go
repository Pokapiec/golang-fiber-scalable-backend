package db

import (
	"todo-server/entities"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDb() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	return db, err
}

func InitialCreate() {
	db, _ := GetDb()
	products := []*entities.Product{
		{Code: "D42", Price: 100},
		{Code: "D43", Price: 20},
		{Code: "D44", Price: 30},
		{Code: "D466", Price: 10},
		{Code: "D452", Price: 1},
		{Code: "D4535", Price: 2121},
	}
	for _, product := range products {
		db.Save(product)
	}
}

func MigrateAll() {
	db, _ := GetDb()
	db.AutoMigrate(&entities.Product{})
}

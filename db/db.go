package db

import (
	"todo-server/entities"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDb() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	return err
}

func MigrateAll() error {
	err := GetDb()
	if err != nil {
		return err
	}
	DB.AutoMigrate(&entities.Product{})
	return nil
}

func InitialCreate() {
	products := []*entities.Product{
		{Code: "D42", Price: 100},
		{Code: "D43", Price: 20},
		{Code: "D44", Price: 30},
		{Code: "D466", Price: 10},
		{Code: "D452", Price: 1},
		{Code: "D4535", Price: 2121},
	}
	for _, product := range products {
		DB.Save(product)
	}
}

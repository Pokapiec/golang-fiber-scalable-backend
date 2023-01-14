package repositories

import (
	"todo-server/db"
	"todo-server/entities"
)

type ProductDTO struct {
	ID    uint   `json:"id"`
	Code  string `json:"code"`
	Price uint   `json:"price"`
}

func GetAllProducts() []ProductDTO {
	db, _ := db.GetDb()
	var products []ProductDTO
	db.Model(&entities.Product{}).Select("code", "price", "id").Find(&products)
	return products
}

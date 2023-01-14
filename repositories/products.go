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
	dbClient, _ := db.GetDb()
	var products []ProductDTO
	dbClient.Model(&entities.Product{}).Select("code", "price", "id").Find(&products)
	return products
}

func GetSingleProduct(productId int) (ProductDTO, error) {
	dbClient, _ := db.GetDb()
	var product ProductDTO
	err := dbClient.Model(&entities.Product{}).Where("id = ?", productId).First(&product).Error
	return product, err
}

func DeleteProduct(productId int) error {
	dbClient, _ := db.GetDb()
	err := dbClient.Delete(&entities.Product{}, productId).Error
	return err
}

func UpdateProduct(productId int, updateValues interface{}) (entities.Product, error) {
	dbClient, _ := db.GetDb()
	var product entities.Product
	err := dbClient.Model(&product).Updates(updateValues).Error
	return product, err
}

func CreateProduct(updateValues interface{}) (entities.Product, error) {
	dbClient, _ := db.GetDb()
	var product entities.Product
	err := dbClient.Model(&product).Create(updateValues).Error
	return product, err
}

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

func GetAllProducts() ([]ProductDTO, error) {
	var products []ProductDTO
	result := db.DB.Model(&entities.Product{}).Find(&products)
	return products, result.Error
}

func GetSingleProduct(productId int) (ProductDTO, error) {
	var product ProductDTO
	result := db.DB.Model(entities.Product{}).First(&product, "id = ?", productId)
	return product, result.Error
}

func DeleteProduct(productId int) error {
	result := db.DB.Delete(&entities.Product{}, productId)
	return result.Error
}

func UpdateProduct(productId int, updateValues ProductDTO) (entities.Product, error) {
	var product entities.Product
	updateData := entities.Product{Code: updateValues.Code, Price: updateValues.Price}
	result := db.DB.Model(&product).Where("id = ?", productId).Updates(updateData)
	return product, result.Error
}

func CreateProduct(updateValues entities.Product) (entities.Product, error) {
	err := db.DB.Create(&updateValues).Error
	return updateValues, err
}

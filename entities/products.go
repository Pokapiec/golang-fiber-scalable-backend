package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code  string `gorm:"uniqueIndex" json:"code"`
	Price uint   `json:"price"`
}

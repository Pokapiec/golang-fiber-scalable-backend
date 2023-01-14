package handlers

import (
	"todo-server/repositories"

	"github.com/gofiber/fiber/v2"
)

func GetAllProductsHandler(c *fiber.Ctx) error {
	products := repositories.GetAllProducts()
	return c.JSON(products)
}

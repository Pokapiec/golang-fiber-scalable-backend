package handlers

import (
	"strconv"
	"todo-server/repositories"

	"github.com/gofiber/fiber/v2"
)

func GetAllProductsHandler(c *fiber.Ctx) error {
	products := repositories.GetAllProducts()
	return c.JSON(products)
}

func GetSingleProductHandler(c *fiber.Ctx) error {
	productIdStr := c.Params("id")
	productId, err := strconv.Atoi(productIdStr)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "bad id"})
	}

	product, err := repositories.GetSingleProduct(productId)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Not found"})
	}

	return c.JSON(product)
}

func DeleteProductHandler(c *fiber.Ctx) error {
	productIdStr := c.Params("id")
	productId, err := strconv.Atoi(productIdStr)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "bad id"})
	}

	err = repositories.DeleteProduct(productId)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Not found"})
	}

	return c.JSON(fiber.Map{"message": "Successfully deleted"})
}

func UpdateProductHandler(c *fiber.Ctx) error {
	productIdStr := c.Params("id")
	productId, err := strconv.Atoi(productIdStr)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "bad id"})
	}

	payload := struct {
		Code  string `json:"code"`
		Price uint   `json:"price"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "bad body"})
	}

	product, err := repositories.UpdateProduct(productId, payload)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Not found"})
	}

	return c.JSON(product)
}

func CreateProductHandler(c *fiber.Ctx) error {
	payload := struct {
		Code  string `json:"code"`
		Price uint   `json:"price"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "bad body"})
	}

	product, err := repositories.CreateProduct(payload)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Not found"})
	}

	return c.JSON(product)
}

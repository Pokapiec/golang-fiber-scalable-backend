package handlers

import (
	"fmt"
	"strconv"
	"todo-server/entities"
	"todo-server/repositories"

	"github.com/gofiber/fiber/v2"
)

//	@Summary		Get a list of products.
//	@Description	fetch products.
//	@Tags			products
//	@Produce		json
//	@Success		200	{object}	[]repositories.ProductDTO
//	@Router			/products [get]
func GetAllProductsHandler(c *fiber.Ctx) error {
	products, err := repositories.GetAllProducts()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "query failed",
		})
	}
	return c.JSON(products)
}

//	@Summary		Get a single product.
//	@Description	fetch a single product.
//	@Tags			products
//	@Param			id	path	int	true	"product ID"
//	@Produce		json
//	@Success		200	{object}	repositories.ProductDTO
//	@Router			/products/{id} [get]
func GetSingleProductHandler(c *fiber.Ctx) error {
	productIdStr := c.Params("id")
	productId, err := strconv.Atoi(productIdStr)

	if err != nil {
		fmt.Printf("Old value %v, converted %v", productIdStr, productId)
		fmt.Println(err)
		return c.Status(400).JSON(fiber.Map{"message": "bad id"})
	}

	product, err := repositories.GetSingleProduct(productId)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Not found"})
	}

	return c.JSON(product)
}

//	@Summary		Delete a single product.
//	@Description	delete a single product by id.
//	@Tags			products
//	@Param			id	path	int	true	"Todo ID"
//	@Produce		json
//	@Success		200	{object}	repositories.ProductDTO
//	@Router			/products/{id} [delete]
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

//	@Summary		Update a product.
//	@Description	Update product.
//	@Tags			products
//	@Param			id	path	int	true	"product ID"
//	@Produce		json
//	@Param			todo	body		repositories.ProductDTO	true	"Product update data"
//	@Success		200		{object}	repositories.ProductDTO
//	@Router			/products/{id} [put]
func UpdateProductHandler(c *fiber.Ctx) error {
	productIdStr := c.Params("id")
	productId, err := strconv.Atoi(productIdStr)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "bad id"})
	}

	var payload repositories.ProductDTO

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "bad body"})
	}

	product, err := repositories.UpdateProduct(productId, payload)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Not found"})
	}

	return c.JSON(product)
}

//	@Summary		Create a product.
//	@Description	Create product.
//	@Tags			products
//	@Produce		json
//	@Param			todo	body		repositories.ProductDTO	true	"Product data"
//	@Success		200		{object}	repositories.ProductDTO
//	@Router			/products [post]
func CreateProductHandler(c *fiber.Ctx) error {
	var payload entities.Product

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "bad body"})
	}

	product, err := repositories.CreateProduct(payload)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Not found"})
	}

	return c.JSON(product)
}

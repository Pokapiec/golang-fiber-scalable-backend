package routers

import (
	"todo-server/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupProductRoutes(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"message": "success",
			"success": true,
		})
	})

	products := app.Group("/products")
	products.Get("/", handlers.GetAllProductsHandler)
	products.Post("/", handlers.CreateProductHandler)
	products.Get("/:id", handlers.GetSingleProductHandler)
	products.Delete("/:id", handlers.DeleteProductHandler)
	products.Put("/:id", handlers.UpdateProductHandler)

}

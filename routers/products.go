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

	todos := app.Group("/products")
	todos.Get("/", handlers.GetAllProductsHandler)
}

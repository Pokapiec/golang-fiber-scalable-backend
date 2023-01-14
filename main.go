package main

import (
	"log"

	"todo-server/db"
	"todo-server/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db.MigrateAll()
	// database.InitialCreate()
	app := fiber.New()
	routers.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}

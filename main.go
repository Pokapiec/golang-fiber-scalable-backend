package main

import (
	"fmt"
	"log"

	"todo-server/db"
	"todo-server/routers"

	_ "todo-server/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/swagger"
)

//	@title			Fiber Example API
//	@version		1.0
//	@description	This is a sample swagger for Fiber
//	@termsOfService	http://swagger.io/terms/
//	@contact.name	API Support
//	@contact.email	fiber@swagger.io
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			localhost:3000
//	@BasePath		/
func main() {
	err := db.MigrateAll()
	if err != nil {
		fmt.Println("Failed to migrate!")
		return
	}
	db.InitialCreate()
	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault) // default swagger
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	routers.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}

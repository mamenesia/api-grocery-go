package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mamenesia/api-grocery/database"
	"github.com/mamenesia/api-grocery/database/migration"
	"github.com/mamenesia/api-grocery/route"
)

func main() {

	// Initiate Database
	database.DatabaseInit()

	// Initiate Migration
	migration.RunMigration()

	app := fiber.New()

	route.RouteInit(app)

	app.Listen(":8080")
}

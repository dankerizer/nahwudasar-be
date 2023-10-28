package main

import (
	"nahwudasar-be/routers"

	"github.com/gofiber/fiber/v2/log"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New(fiber.Config{
		EnablePrintRoutes: false,
		Prefork:           false,
		JSONEncoder:       json.Marshal,
		JSONDecoder:       json.Unmarshal})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Nahwu Dasar nahwu.net")
	})
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Use(cors.New())
	routers.FiilRouter(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen("0.0.0.0:4000"))
}

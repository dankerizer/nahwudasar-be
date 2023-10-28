package routers

import (
	"nahwudasar-be/controllers"

	"github.com/gofiber/fiber/v2"
)

func FiilRouter(app *fiber.App) {
	handler := controllers.NewFiilHandler(conn)
	app.Get("fiils/:page", handler.GetAll)
	app.Get("fiil/:id", handler.GetOne)
}

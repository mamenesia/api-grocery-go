package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mamenesia/api-grocery/handler"
)

func RouteInit(r *fiber.App) {
	r.Get("/user", handler.UserHandlerGetAll)
	r.Post("/user", handler.UserHandlerCreate)

}

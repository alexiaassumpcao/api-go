package router

import (
	handler "github.com/alexiaassumpcao/api-go/internal/handler"
	"github.com/gofiber/fiber/v2"
)

type router struct {
	app             *fiber.App
	messagesHandler handler.Handler
}

func NewRouter(app *fiber.App, messagesHandler handler.Handler) router {
	return router{
		app:             app,
		messagesHandler: messagesHandler,
	}
}

func (r *router) MessagesRoutes() {
	r.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("ola")
	})
	r.app.Get("/messages", r.messagesHandler.GetAll)
}

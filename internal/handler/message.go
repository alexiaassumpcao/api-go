package handler

import (
	presenter "github.com/alexiaassumpcao/api-go/internal/presenter"
	msgUseCases "github.com/alexiaassumpcao/api-go/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type MessagesHandler struct {
	MessageUseCases msgUseCases.UseCases
}

func (h MessagesHandler) GetAll(c *fiber.Ctx) error {
	messages, err := h.MessageUseCases.GetAll()
	if err != nil {
		return c.SendString("Deu ruim")
	}
	response := presenter.MountResponse(*messages)

	return c.JSON(response)
}

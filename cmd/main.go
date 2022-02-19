package main

import (
	handler "github.com/alexiaassumpcao/api-go/internal/handler"
	"github.com/alexiaassumpcao/api-go/internal/repository"
	r "github.com/alexiaassumpcao/api-go/internal/router"
	usecase "github.com/alexiaassumpcao/api-go/internal/usecase"
	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	messageRepo := repository.MessageRepository{}

	usecaseMessage := usecase.MessageUseCases{
		Repository: messageRepo,
	}

	handlerMessages := handler.MessagesHandler{
		MessageUseCases: usecaseMessage,
	}

	router := r.NewRouter(app, handlerMessages)

	router.MessagesRoutes()

	app.Listen(":8080")
}

package mock

import "github.com/alexiaassumpcao/api-go/internal/model"

var Messages = []model.Message{
	{
		ID:         "1",
		UserSender: "Alexia",
		Content:    "Olá, tudo bem?",
	},
	{
		ID:         "2",
		UserSender: "Hantaro",
		Content:    "Quero comida!",
	},
}

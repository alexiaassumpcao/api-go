package presenter

import "github.com/alexiaassumpcao/api-go/internal/model"

type ListAllMessages struct {
	Messages []model.Message `json:"messages"`
}

func MountResponse(messages []model.Message) *ListAllMessages {
	return &ListAllMessages{
		Messages: messages,
	}
}

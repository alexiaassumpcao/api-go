package usecase

import (
	model "github.com/alexiaassumpcao/api-go/internal/model"
	"github.com/alexiaassumpcao/api-go/internal/repository"
)

type MessageUseCases struct {
	Repository repository.Repository
}

func (u MessageUseCases) GetAll() (*[]model.Message, error) {
	messages, err := u.Repository.ListAll()
	if err != nil {
		return nil, err
	}
	return messages, nil
}

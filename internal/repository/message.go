package repository

import (
	"github.com/alexiaassumpcao/api-go/internal/model"
	"github.com/alexiaassumpcao/api-go/internal/repository/mock"
)

type MessageRepository struct{}

func (r MessageRepository) ListAll() (*[]model.Message, error) {
	return &mock.Messages, nil
}

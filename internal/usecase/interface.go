package usecase

import (
	model "github.com/alexiaassumpcao/api-go/internal/model"
)

type UseCases interface {
	GetAll() (*[]model.Message, error)
}

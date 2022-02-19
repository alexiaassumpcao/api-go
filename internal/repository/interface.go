package repository

import "github.com/alexiaassumpcao/api-go/internal/model"

type Repository interface {
	ListAll() (*[]model.Message, error)
}

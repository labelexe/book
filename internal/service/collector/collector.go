package collector

import "github.com/reucot/parser/internal/storage/repository"

type Service struct {
}

func New(repo *repository.Repository) *Service {
	return &Service{}
}

package collector

import (
	"context"

	"github.com/reucot/parser/internal/service/collector/onexstavka"
	"github.com/reucot/parser/internal/storage/repository"
)

type Collector interface {
	Start(ctx context.Context)
}

type Service struct {
	collectors []Collector
}

func New(reps *repository.Repository) *Service {
	c := make([]Collector, 1)

	c[0] = *onexstavka.NewOneXStavka(reps)

	return &Service{
		collectors: c,
	}
}

func (s Service) Run() {
	ctx := context.Background()

	for _, collector := range s.collectors {
		collector.Start(ctx)
	}
}

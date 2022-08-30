package repository

import (
	"context"

	"github.com/reucot/parser/internal/models"
	"github.com/reucot/parser/internal/storage/repository/psql"

	"github.com/jmoiron/sqlx"
)

type Category interface {
	Create(ctx context.Context, item models.Category) (int, error)
	GetByID(ctx context.Context, id int) (models.Category, error)
	GetAll(ctx context.Context) ([]models.Category, error)
	Update(ctx context.Context, id int, input models.Category) error
	Delete(ctx context.Context, id int) error
}

type Country interface {
	Create(ctx context.Context, item models.Country) (int, error)
	GetByID(ctx context.Context, id int) (models.Country, error)
	GetAll(ctx context.Context) ([]models.Country, error)
	Update(ctx context.Context, id int, input models.Country) error
	Delete(ctx context.Context, id int) error
}

type Team interface {
	Create(ctx context.Context, item models.Team) (int, error)
	GetByID(ctx context.Context, id int) (models.Team, error)
	GetAll(ctx context.Context) ([]models.Team, error)
	Update(ctx context.Context, id int, input models.Team) error
	Delete(ctx context.Context, id int) error
}

type Bet interface {
	Create(ctx context.Context, item models.Bet) (int, error)
	GetByID(ctx context.Context, id int) (models.Bet, error)
	GetAll(ctx context.Context) ([]models.Bet, error)
	GetAllByEventID(ctx context.Context, id int) ([]models.Bet, error)
	Update(ctx context.Context, id int, input models.Bet) error
	UpdateCoefficient(ctx context.Context, id, coef int) error
	Delete(ctx context.Context, id int) error
}

type Event interface {
	Create(ctx context.Context, item models.Event) (int, error)
	GetByID(ctx context.Context, id int) (models.Event, error)
	GetAll(ctx context.Context) ([]models.Event, error)
	GetAllByLineID(ctx context.Context, id int) ([]models.Event, error)
	Update(ctx context.Context, id int, input models.Event) error
	Delete(ctx context.Context, id int) error
}

type Repository struct {
	Category
	Country
	Team
	Bet
	Event
}

func NewPsql(db *sqlx.DB) *Repository {
	return &Repository{
		Category: psql.NewCategory(db),
		Country:  psql.NewCountry(db),
		Team:     psql.NewTeam(db),
		Bet:      psql.NewBet(db),
		Event:    psql.NewEvent(db),
	}
}

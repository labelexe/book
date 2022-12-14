package psql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	CategoriesTable  = "categories"
	CountriesTable   = "countries"
	TeamsTable       = "teams"
	GamesTable       = "games"
	LinesTable       = "lines"
	EventsTable      = "events"
	EventsLinesTable = "events_lines"
	BetsTable        = "bets"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func New(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

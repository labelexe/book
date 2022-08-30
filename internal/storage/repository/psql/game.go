package psql

import "github.com/jmoiron/sqlx"

type Game struct {
	db *sqlx.DB
}

func NewGame(db *sqlx.DB) *Game {
	return &Game{
		db: db,
	}
}

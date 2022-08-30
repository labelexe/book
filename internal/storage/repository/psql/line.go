package psql

import "github.com/jmoiron/sqlx"

type Line struct {
	db *sqlx.DB
}

func NewLine(db *sqlx.DB) *Line {
	return &Line{
		db: db,
	}
}

package psql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/reucot/parser/internal/models"
)

type Event struct {
	db *sqlx.DB
}

func NewEvent(db *sqlx.DB) *Event {
	return &Event{
		db: db,
	}
}

func (e Event) Create(ctx context.Context, item models.Event) (int, error) {
	query := fmt.Sprintf(`INSERT INTO %s (name_ru, name_en) VALUES ($1, $2) RETURNING id`, EventsTable)

	id := 0
	if err := e.db.QueryRowContext(ctx, query, item.NameRU, item.NameEN).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (e Event) GetByID(ctx context.Context, id int) (models.Event, error) {
	output := models.Event{}
	query := fmt.Sprintf(`SELECT id, name_ru, name_en FROM %s WHERE id = $1`, EventsTable)

	if err := e.db.GetContext(ctx, &output, query, id); err != nil && err != sql.ErrNoRows {
		return models.Event{}, err
	}

	return output, nil
}

func (e Event) GetAll(ctx context.Context) ([]models.Event, error) {
	output := []models.Event{}
	query := fmt.Sprintf(`SELECT id, name_ru, name_en FROM %s`, EventsTable)

	if err := e.db.SelectContext(ctx, &output, query); err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return output, nil
}

func (e Event) GetAllByLineID(ctx context.Context, id int) ([]models.Event, error) {
	output := []models.Event{}
	query := fmt.Sprintf(`SELECT id, name_ru, name_en FROM %s
	WHERE id IN (SELECT event_id FROM %s WHERE line_id=$1)`, EventsTable, EventsLinesTable)

	if err := e.db.SelectContext(ctx, &output, query); err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return output, nil
}

func (e Event) Update(ctx context.Context, id int, input models.Event) error {
	query := fmt.Sprintf(`UPDATE %s SET name_ru=$2, name_en=$3 WHERE id = $1`, EventsTable)

	if _, err := e.db.ExecContext(ctx, query, id, input.NameRU, input.NameEN); err != nil {
		return err
	}

	return nil
}

func (e Event) Delete(ctx context.Context, id int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, EventsTable)

	if _, err := e.db.ExecContext(ctx, query, id); err != nil {
		return err
	}

	return nil
}

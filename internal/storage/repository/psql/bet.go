package psql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/reucot/parser/internal/models"
)

type Bet struct {
	db *sqlx.DB
}

func NewBet(db *sqlx.DB) *Bet {
	return &Bet{
		db: db,
	}
}

func (b Bet) Create(ctx context.Context, item models.Bet) (int, error) {
	query := fmt.Sprintf(`INSERT INTO %s (name_ru, name_en, event_id, coefficient) VALUES ($1, $2, $3, $4) RETURNING id`, BetsTable)

	id := 0
	if err := b.db.QueryRowContext(ctx, query, item.NameRU, item.NameEN, item.EventID, item.Coefficient).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (b Bet) GetByID(ctx context.Context, id int) (models.Bet, error) {
	output := models.Bet{}
	query := fmt.Sprintf(`SELECT id, name_ru, name_en, event_id, coefficient FROM %s WHERE id = $1`, BetsTable)

	if err := b.db.GetContext(ctx, &output, query, id); err != nil && err != sql.ErrNoRows {
		return models.Bet{}, err
	}

	return output, nil
}

//TODO: Limit offset
func (b Bet) GetAll(ctx context.Context) ([]models.Bet, error) {
	output := []models.Bet{}
	query := fmt.Sprintf(`SELECT id, name_ru, name_en, event_id, coefficient FROM %s`, BetsTable)

	if err := b.db.SelectContext(ctx, &output, query); err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return output, nil
}

//TODO: Limit offset???
func (b Bet) GetAllByEventID(ctx context.Context, id int) ([]models.Bet, error) {
	output := []models.Bet{}
	query := fmt.Sprintf(`SELECT id, name_ru, name_en, event_id, coefficient WHERE event_id=$1 FROM %s`, BetsTable)

	if err := b.db.SelectContext(ctx, &output, query, id); err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return output, nil
}

func (b Bet) Update(ctx context.Context, id int, input models.Bet) error {
	query := fmt.Sprintf(`UPDATE %s SET name_ru=$2, name_en=$3, event_id=$4, coefficient=$5 WHERE id = $1`, BetsTable)

	if _, err := b.db.ExecContext(ctx, query, id, input.NameRU, input.NameEN, input.EventID, input.Coefficient); err != nil {
		return err
	}

	return nil
}

func (b Bet) UpdateCoefficient(ctx context.Context, id, coef int) error {
	query := fmt.Sprintf(`UPDATE %s SET coefficient=$2 WHERE id = $1`, BetsTable)

	if _, err := b.db.ExecContext(ctx, query, id, coef); err != nil {
		return err
	}

	return nil
}

func (b Bet) Delete(ctx context.Context, id int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, BetsTable)

	if _, err := b.db.ExecContext(ctx, query, id); err != nil {
		return err
	}

	return nil
}

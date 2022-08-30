package psql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/reucot/parser/internal/models"
)

type Team struct {
	db *sqlx.DB
}

func NewTeam(db *sqlx.DB) *Team {
	return &Team{
		db: db,
	}
}

func (t Team) Create(ctx context.Context, item models.Team) (int, error) {
	query := fmt.Sprintf(`INSERT INTO %s (category_id, name_ru, name_en, image) VALUES ($1, $2, $3, $4) RETURNING id`, TeamsTable)

	id := 0
	if err := t.db.QueryRowContext(ctx, query, item.CategoryID, item.NameRU, item.NameEN, item.Image).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (t Team) GetByID(ctx context.Context, id int) (models.Team, error) {
	output := models.Team{}
	query := fmt.Sprintf(`SELECT id, category_id, name_ru, name_en, image FROM %s WHERE id = $1`, TeamsTable)

	if err := t.db.GetContext(ctx, &output, query, id); err != nil && err != sql.ErrNoRows {
		return models.Team{}, err
	}

	return output, nil
}

//TODO: limit offset
func (t Team) GetAll(ctx context.Context) ([]models.Team, error) {
	output := []models.Team{}
	query := fmt.Sprintf(`SELECT id, category_id, name_ru, name_en, image FROM %s`, TeamsTable)

	if err := t.db.SelectContext(ctx, &output, query); err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return output, nil
}

func (t Team) Update(ctx context.Context, id int, input models.Team) error {
	query := fmt.Sprintf(`UPDATE %s SET category_id = $2, name_ru=$3, name_en=$4, image=$5 WHERE id = $1`, TeamsTable)

	if _, err := t.db.ExecContext(ctx, query, id, input.CategoryID, input.NameRU, input.NameEN, input.Image); err != nil {
		return err
	}

	return nil
}

func (t Team) Delete(ctx context.Context, id int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, TeamsTable)

	if _, err := t.db.ExecContext(ctx, query, id); err != nil {
		return err
	}

	return nil
}

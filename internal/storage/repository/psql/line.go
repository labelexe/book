package psql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/reucot/parser/internal/models"
)

type Line struct {
	db *sqlx.DB
}

func NewLine(db *sqlx.DB) *Line {
	return &Line{
		db: db,
	}
}

func (l Line) Create(ctx context.Context, item models.Line) (int, error) {
	query := fmt.Sprintf(`INSERT INTO %s 
	(name_ru, name_en, category_id, country_id, tourney, type_line, games_id, api_id, api_src) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`, LinesTable)

	id := 0
	err := l.db.QueryRowContext(ctx,
		query,
		item.NameRU, item.NameEN, item.CategoryID, item.Country,
		item.Tourney, item.TypeLine, item.GameID, item.ApiID, item.ApiSrc).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (l Line) GetByID(ctx context.Context, id int) (models.Line, error) {
	output := models.Line{}
	query := fmt.Sprintf(`
	SELECT 
		name_ru, 
		name_en, 
		category_id, 
		country_id, 
		tourney, 
		type_line, 
		games_id,
		api_id,
		api_src 
	FROM %s 
	WHERE id = $1`, LinesTable)

	if err := l.db.GetContext(ctx, &output, query, id); err != nil && err != sql.ErrNoRows {
		return models.Line{}, err
	}

	return output, nil
}

func (l Line) GetAll(ctx context.Context) ([]models.Line, error) {
	output := []models.Line{}
	query := fmt.Sprintf(`
	SELECT 
		name_ru, 
		name_en, 
		category_id, 
		country_id, 
		tourney, 
		type_line, 
		games_id,
		api_id,
		api_src 
	FROM %s `, LinesTable)

	if err := l.db.SelectContext(ctx, &output, query); err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return output, nil
}

func (l Line) Update(ctx context.Context, id int, input models.Line) error {
	query := fmt.Sprintf(`
	UPDATE %s 
	SET 
		name_ru=$2, name_en=$3, category_id=$4, country_id=$5,
		tourney=$6, type_line=$7, games_id=$8, api_id=$9, api_src=$10  
	WHERE id = $1`, LinesTable)

	_, err := l.db.ExecContext(ctx, query, id,
		input.NameRU, input.NameEN, input.CategoryID, input.Country,
		input.Tourney, input.TypeLine, input.GameID, input.ApiID, input.ApiSrc)

	if err != nil {
		return err
	}

	return nil
}

func (l Line) UpdateTypeLine(ctx context.Context, id int, tl string) error {
	query := fmt.Sprintf(`
	UPDATE %s 
	SET 
		type_line=$2 
	WHERE id = $1`, LinesTable)

	if _, err := l.db.ExecContext(ctx, query, id, tl); err != nil {
		return err
	}

	return nil
}

func (l Line) Delete(ctx context.Context, id int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, LinesTable)

	if _, err := l.db.ExecContext(ctx, query, id); err != nil {
		return err
	}

	return nil
}

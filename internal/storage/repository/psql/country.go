package psql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/reucot/parser/internal/models"
)

type Country struct {
	db *sqlx.DB
}

func NewCountry(db *sqlx.DB) *Country {
	return &Country{
		db: db,
	}
}

func (c Country) Create(ctx context.Context, item models.Country) (int, error) {
	query := fmt.Sprintf(`INSERT INTO %s (name_ru, name_en) VALUES ($1, $2) RETURNING id`, CountriesTable)

	id := 0
	if err := c.db.QueryRowContext(ctx, query, item.NameRU, item.NameEN).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (c Country) GetByID(ctx context.Context, id int) (models.Country, error) {
	output := models.Country{}
	query := fmt.Sprintf(`SELECT id, name_ru, name_en FROM %s WHERE id = $1`, CountriesTable)

	if err := c.db.GetContext(ctx, &output, query, id); err != nil && err != sql.ErrNoRows {
		return models.Country{}, err
	}

	return output, nil
}

func (c Country) GetAll(ctx context.Context) ([]models.Country, error){
	output := []models.Country{}
	query := fmt.Sprintf(`SELECT id, name_ru, name_en FROM %s`, CountriesTable)

	if err := c.db.SelectContext(ctx, &output, query); err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return output, nil
}

func (c Country) Update(ctx context.Context, id int, input models.Country) error {
	query := fmt.Sprintf(`UPDATE %s SET name_ru=$2, name_en=$3 WHERE id = $1`, CountriesTable)

	if _, err := c.db.ExecContext(ctx, query, id, input.NameRU, input.NameEN); err != nil {
		return err
	}

	return nil
}

func (c Country) Delete(ctx context.Context, id int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, CountriesTable)

	if _, err := c.db.ExecContext(ctx, query, id); err != nil {
		return err
	}

	return nil
}

package psql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/reucot/parser/internal/models"
)

type Category struct {
	db *sqlx.DB
}

func NewCategory(db *sqlx.DB) *Category {
	return &Category{
		db: db,
	}
}

func (c Category) Create(ctx context.Context, item models.Category) (int, error) {
	query := fmt.Sprintf(`INSERT INTO %s (name_ru, name_en, api_id, api_src) VALUES ($1, $2, $3, $4) RETURNING id`, CategoriesTable)

	id := 0
	if err := c.db.QueryRowContext(ctx, query, item.NameRU, item.NameEN, item.ApiID, item.ApiSrc).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (c Category) GetByID(ctx context.Context, id int) (models.Category, error) {
	output := models.Category{}
	query := fmt.Sprintf(`SELECT id, name_ru, name_en, api_id, api_src FROM %s WHERE id = $1`, CategoriesTable)

	if err := c.db.GetContext(ctx, &output, query, id); err != nil && err != sql.ErrNoRows {
		return models.Category{}, err
	}

	return output, nil
}

func (c Category) GetByApiID(ctx context.Context, id int, src string) (models.Category, error) {
	output := models.Category{}
	query := fmt.Sprintf(`SELECT id, name_ru, name_en, api_id, api_src FROM %s WHERE api_id = $1 and api_src = $2`, CategoriesTable)

	if err := c.db.GetContext(ctx, &output, query, id, src); err != nil && err != sql.ErrNoRows {
		return models.Category{}, err
	}

	return output, nil
}

func (c Category) GetAll(ctx context.Context) ([]models.Category, error) {
	output := []models.Category{}
	query := fmt.Sprintf(`SELECT id, name_ru, name_en, api_id, api_src FROM %s`, CategoriesTable)

	if err := c.db.SelectContext(ctx, &output, query); err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return output, nil
}

func (c Category) Update(ctx context.Context, id int, input models.Category) error {
	query := fmt.Sprintf(`UPDATE %s SET name_ru=$2, name_en=$3, api_id=$4, api_src=$5 WHERE id = $1`, CategoriesTable)

	if _, err := c.db.ExecContext(ctx, query, id, input.NameRU, input.NameEN, input.ApiID, input.ApiSrc); err != nil {
		return err
	}

	return nil
}

func (c Category) Delete(ctx context.Context, id int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, CategoriesTable)

	if _, err := c.db.ExecContext(ctx, query, id); err != nil {
		return err
	}

	return nil
}

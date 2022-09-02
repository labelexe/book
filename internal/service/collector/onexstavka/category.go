package onexstavka

import (
	"context"
	"fmt"

	"github.com/reucot/parser/internal/models"
	"github.com/reucot/parser/internal/storage/repository"
	"github.com/reucot/parser/pkg/apionexstavka"
)

var Categories map[string]int = map[string]int{
	"Football":   1,
	"Basketball": 3,
	"Volleyball": 5,
	"Tennis":     2,
	"Handball":   20,
	"UFC":        10,
	"Boxing":     16,
	"Wrestling":  17,
	"Esports":    9,
}

type CategoryService struct {
	rep repository.Category
}

func NewCategoryService(rep repository.Category) *CategoryService {
	return &CategoryService{
		rep: rep,
	}
}

func (c CategoryService) LoadCategories(ctx context.Context) ([]models.Category, error) {
	api := apionexstavka.NewCategory()

	cs, err := api.GetLineCategories()
	if err != nil {
		return nil, err
	}

	for _, v := range Categories {
		_, err := c.rep.Create(ctx, models.Category{
			NameRU: cs[v].N,
			NameEN: cs[v].E,
			ApiID:  cs[v].I,
			ApiSrc: ApiOneXStavka,
		})

		if err != nil {
			return nil, err
		}
	}

	ncs, err := c.rep.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, v := range ncs {
		fmt.Println(v)
	}

	return ncs, nil
}

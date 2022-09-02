package onexstavka

import (
	"context"
	"fmt"

	"github.com/reucot/parser/internal/models"
	"github.com/reucot/parser/internal/storage/repository"
	"github.com/reucot/parser/pkg/apionexstavka"

	"github.com/sirupsen/logrus"
)

const ApiOneXStavka = "1x"

type OneXStavkaService struct {
	Category
}

type Category interface {
	LoadCategories(ctx context.Context) ([]models.Category, error)
}

func NewOneXStavka(reps *repository.Repository) *OneXStavkaService {
	return &OneXStavkaService{
		Category: NewCategoryService(reps.Category),
	}
}

// c := onexstavka.NewCategory("1xstavka.ru")
// f, err := c.GetLineCategories()
// if err != nil {
// 	return
// }

// c := apionexstavka.NewMatch()
// f, err := c.GetMatches(4, false)
// if err != nil {
// 	fmt.Println(err.Error())
// 	return
// }

// for _, v := range f {
// 	// fmt.Println(v.I)
// 	r, err := c.GetMatch(v.I, false)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	fmt.Printf("%s - %s\n", r.O1, r.O2)
// }

//TODO: Error chanel
func (o OneXStavkaService) Start(ctx context.Context) {

	cs, err := o.Category.LoadCategories(ctx)
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Info(cs[0].NameRU)
	o.HandleCategory(cs[0].ApiID)
	// for _, c := range cs{
	//go o.HandleCategory(c.OneXStavkaID)
	// }

}

func (o OneXStavkaService) HandleCategory(categoryID int) error {
	api := apionexstavka.NewMatch(categoryID)

	ms, err := api.GetMatches(false)
	if err != nil {
		return err
	}
	//TODO: Создавать мапу где ID матча с апи = наш ID в базе
	for _, m := range ms {
		nm, err := api.GetMatch(m.I, false)
		if err != nil {
			return err
		}
		logrus.Info(fmt.Sprintf("%s - %s", nm.O1, nm.O2))
	}

	return nil
}

package onexstavka

import (
	"context"
	"sync"

	"github.com/reucot/parser/internal/models"
	"github.com/reucot/parser/internal/storage/repository"
	"github.com/reucot/parser/pkg/apionexstavka"
)

type LineService struct {
	repCategory repository.Category
	repLine     repository.Line
	repGame     repository.Game
	repTeam     repository.Team
	linesMap    sync.Map
}

func NewLineService(repLine repository.Line, repGame repository.Game, repTeam repository.Team, repCategory repository.Category) *LineService {
	return &LineService{
		repCategory: repCategory,
		repLine:     repLine,
		repGame:     repGame,
		repTeam:     repTeam,
	}
}

func (l LineService) CreateOrUpdate(ctx context.Context, match apionexstavka.InsteadMatchStruct) error {
	m, ok := l.linesMap.Load(match.I)
	if !ok {
		id, err := l.Create(ctx, match)
		if err != nil {
			return err
		}

		l.linesMap.Store(match.I, id)
	}
}

func (l LineService) Create(ctx context.Context, match apionexstavka.InsteadMatchStruct) (int, error) {

	c, err := l.repCategory.GetByApiID(ctx, match.Si, ApiOneXStavka)
	if err != nil {
		return 0, err
	}

	t1ID, t2ID, err := l.createOrGetTeams(ctx, c.ID, match)
	if err != nil {
		return 0, err
	}

	m := models.Game{
		HomeTeamID: t1ID,
		AwayTeamID: t2ID,
	}
}

//TODO: Img
func (l LineService) createOrGetTeams(ctx context.Context, categoryID int, match apionexstavka.InsteadMatchStruct) (int, int, error) {
	//TODO: Сделать поиск по Api_id и api_src
	t1 := models.Team{
		CategoryID: categoryID,
		NameRU:     match.O1,
		NameEN:     match.O1E,
		// Image: match.O1Img,
	}

	t1ID, err := l.repTeam.Create(ctx, t1)
	if err != nil {
		return 0, 0, err
	}

	t2 := models.Team{
		CategoryID: categoryID,
		NameRU:     match.O2,
		NameEN:     match.O2E,
	}

	t2ID, err := l.repTeam.Create(ctx, t2)
	if err != nil {
		return 0, 0, err
	}

	return t1ID, t2ID, nil
}

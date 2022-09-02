package onexstavka

import (
	"github.com/reucot/parser/internal/storage/repository"
)

type GameService struct {
	rep repository.Game
}

func NewGameService(rep repository.Game) *GameService {
	return &GameService{
		rep: rep,
	}
}

// func (g GameService) Create(ctx context.Context, match apionexstavka.InsteadMatchStruct) error{
// 	// g.rep.
// }

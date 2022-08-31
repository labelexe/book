package psql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/reucot/parser/internal/models"

	"github.com/jmoiron/sqlx"
)

type Game struct {
	db *sqlx.DB
}

func NewGame(db *sqlx.DB) *Game {
	return &Game{
		db: db,
	}
}

func (g Game) Create(ctx context.Context, item models.Game) (int, error) {
	query := fmt.Sprintf(`INSERT INTO %s 
	(home_team_id, away_team_id, home_team_score, away_team_score, start_date, current_event_time, time_events) 
	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`, GamesTable)

	id := 0
	err := g.db.QueryRowContext(ctx,
		query,
		item.HomeTeamID, item.AwayTeamID, item.HomeTeamScore, item.AwayTeamScore,
		item.StartDate, item.CurrentEventTime, item.TimeEvents).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (g Game) GetByID(ctx context.Context, id int) (models.Game, error) {
	output := models.Game{}
	query := fmt.Sprintf(`
	SELECT 
		home_team_id, 
		away_team_id, 
		home_team_score, 
		away_team_score, 
		start_date, 
		current_event_time, 
		time_events 
	FROM %s 
	WHERE id = $1`, GamesTable)

	if err := g.db.GetContext(ctx, &output, query, id); err != nil && err != sql.ErrNoRows {
		return models.Game{}, err
	}

	return output, nil
}

func (g Game) GetAll(ctx context.Context) ([]models.Game, error) {
	output := []models.Game{}
	query := fmt.Sprintf(`
	SELECT 
		home_team_id, 
		away_team_id, 
		home_team_score, 
		away_team_score, 
		start_date, 
		current_event_time, 
		time_events 
	FROM %s`, GamesTable)

	if err := g.db.SelectContext(ctx, &output, query); err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return output, nil
}

func (g Game) Update(ctx context.Context, id int, input models.Game) error {
	query := fmt.Sprintf(`
	UPDATE %s 
	SET 
		home_team_id=$2, away_team_id=$3, home_team_score=$4, away_team_score=$5,
		start_date=$6, current_event_time=$7, time_events=$8 
	WHERE id = $1`, GamesTable)

	_, err := g.db.ExecContext(ctx, query, id,
		input.HomeTeamID, input.AwayTeamID, input.HomeTeamScore, input.AwayTeamScore,
		input.StartDate, input.CurrentEventTime, input.TimeEvents)

	if err != nil {
		return err
	}

	return nil
}

func (g Game) UpdateScore(ctx context.Context, id, scores int, home bool) error {
	team := "away_team_score=away_team_score+$2"
	if home {
		team = "home_team_score=home_team_score+$2"
	}

	query := fmt.Sprintf(`
	UPDATE %s 
	SET`+team+` 
	WHERE id = $1`, GamesTable)

	if _, err := g.db.ExecContext(ctx, query, id,
		scores); err != nil {
		return err
	}

	return nil
}

func (g Game) UpdateCurrentEventTime(ctx context.Context, id int, cet time.Time) error {
	query := fmt.Sprintf(`
	UPDATE %s 
	SET current_event_time=$2 
	WHERE id = $1`, GamesTable)

	if _, err := g.db.ExecContext(ctx, query, id, cet); err != nil {
		return err
	}

	return nil
}

func (g Game) Delete(ctx context.Context, id int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, GamesTable)

	if _, err := g.db.ExecContext(ctx, query, id); err != nil {
		return err
	}

	return nil
}

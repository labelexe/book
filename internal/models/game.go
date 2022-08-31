package models

import "time"

type Game struct {
	ID               int         `json:"id" db:"id"`
	HomeTeamID       int         `json:"home_team_id" db:"home_team_id"`
	AwayTeamID       int         `json:"away_team_id" db:"away_team_id"`
	HomeTeamScore    int         `json:"home_team_score" db:"home_team_score"`
	AwayTeamScore    int         `json:"away_team_score" db:"away_team_score"`
	StartDate        time.Time   `json:"start_date" db:"start_date"`
	CurrentEventTime time.Time   `json:"current_event_time" db:"current_event_time"`
	TimeEvents       []time.Time `json:"time_events" db:"time_events"`
}

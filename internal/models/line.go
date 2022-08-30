package models

type Line struct {
	ID         int    `json:"id" db:"id"`
	NameRU     string `json:"name_ru" db:"name_ru"`
	NameEN     string `json:"name_en" db:"name_en"`
	CategoryID int    `json:"category_id" db:"category_id"`
	CountryID  int    `json:"country_id" db:"country_id"`
	Tourney    string `json:"tourney" db:"tourney"`
	TypeLine   string `json:"type_line" db:"type_line"`
	GameID     int    `json:"game_id" db:"game_id"`
}

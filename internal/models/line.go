package models

type Line struct {
	ID         int    `json:"id" db:"id"`
	NameRU     string `json:"name_ru" db:"name_ru"`
	NameEN     string `json:"name_en" db:"name_en"`
	CategoryID int    `json:"category_id" db:"category_id"`
	Country    string `json:"country" db:"country"`
	Tourney    string `json:"tourney" db:"tourney"`
	TypeLine   string `json:"type_line" db:"type_line"`
	GameID     int    `json:"game_id" db:"game_id"`
	ApiID      int    `json:"api_id" db:"api_id"`
	ApiSrc     string `json:"api_src" db:"api_src"`
}

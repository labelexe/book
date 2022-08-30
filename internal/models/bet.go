package models

type Bet struct {
	ID          int    `json:"id" db:"id"`
	NameRU      string `json:"name_ru" db:"name_ru"`
	NameEN      string `json:"name_en" db:"name_en"`
	EventID     int    `json:"event_id" db:"event_id"`
	Coefficient int    `json:"coefficient" db:"coefficient"`
}

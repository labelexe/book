package models

type Event struct {
	ID     int    `json:"id" db:"id"`
	NameRU string `json:"name_ru" db:"name_ru"`
	NameEN string `json:"name_en" db:"name_en"`
}

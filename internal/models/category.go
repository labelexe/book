package models

type Category struct {
	ID     int    `json:"id" db:"id"`
	NameRU string `json:"name_ru" db:"name_ru"`
	NameEN string `json:"name_en" db:"name_en"`
	ApiID  int    `json:"api_id" db:"api_id"`
	ApiSrc string `json:"api_src" db:"api_src"`
}

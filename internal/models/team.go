package models

type Team struct {
	ID         int    `json:"id" db:"id"`
	CategoryID int    `json:"category_id" db:"category_id"`
	NameRU     string `json:"name_ru" db:"name_ru"`
	NameEN     string `json:"name_en" db:"name_en"`
	Image      byte   `json:"image" db:"image"`
}

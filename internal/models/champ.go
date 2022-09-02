package models

type Champ struct {
	ID           int    `json:"id" db:"id"`
	NameRU       string `json:"name_ru" db:"name_ru"`
	NameEN       string `json:"name_en" db:"name_en"`
	CountMatch   int    `json:"count_match" db:"count_match"`
	OneXStavkaID int    `json:"one_x_stavka_id" db:"one_x_stavka_id"`
}

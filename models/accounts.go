package models

type Account struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	NIK        string  `json:"nik"`
	NoHP       string  `json:"no_hp"`
	NoRekening string  `json:"no_rekening"`
	Balance    float64 `json:"balance"`
}

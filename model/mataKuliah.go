package model

type MataKuliah struct {
	ID        int    `json:"id" gorm:"primary_key"`
	KodeMk    string `json:"kode_mk"`
	Nama      string `json:"nama"`
	Deskripsi string `json:"deskripsi"`
	Sks       int    `json:"sks"`
	OBEId     int    `json:"obe_id"`
}

func (MataKuliah) TableName() string {
	return "mata_kuliah"
}

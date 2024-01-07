package model

type PLO struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Nama      string `json:"nama"`
	Deskripsi string `json:"deskripsi"`
	OBEId     int    `json:"obe_id"`
}

func (PLO) TableName() string {
	return "plo"
}

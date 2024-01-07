package model

type Dosen struct {
	ID        int    `json:"id" gorm:"primary_key"`
	KodeDosen string `json:"kode_dosen"`
	Nama      string `json:"nama"`
	UserId    int    `json:"user_id"`
}

func (Dosen) TableName() string {
	return "dosen"
}

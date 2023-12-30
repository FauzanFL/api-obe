package model

type Dosen struct {
	KodeDosen string `json:"kode_dosen" gorm:"primary_key"`
}

func (Dosen) TableName() string {
	return "dosen"
}

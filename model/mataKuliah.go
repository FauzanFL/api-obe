package model

type MataKuliah struct {
	ID int `json:"id" gorm:"primary_key"`
}

func (MataKuliah) TableName() string {
	return "mata_kuliah"
}

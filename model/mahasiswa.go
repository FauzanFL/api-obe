package model

type Mahasiswa struct {
	NIM string `json:"nim" gorm:"primary_key"`
}

func (Mahasiswa) TableName() string {
	return "mahasiswa"
}

package model

type MkMahasiswa struct {
	ID int `json:"id" gorm:"primary_key"`
}

func (MkMahasiswa) TableName() string {
	return "mk_mahasiswa"
}

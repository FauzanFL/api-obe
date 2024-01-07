package model

type Mahasiswa struct {
	ID      int    `json:"id" gorm:"primary_key"`
	NIM     string `json:"nim"`
	Nama    string `json:"nama"`
	KelasId int    `json:"kelas_id"`
}

func (Mahasiswa) TableName() string {
	return "mahasiswa"
}

package model

type MkMahasiswa struct {
	ID      int `json:"id" gorm:"primary_key"`
	MKId    int `json:"mk_id"`
	MhsId   int `json:"mhs_id"`
	KelasId int `json:"kelas_id"`
}

func (MkMahasiswa) TableName() string {
	return "mk_mahasiswa"
}

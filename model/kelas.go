package model

type Kelas struct {
	ID        int    `json:"id" gorm:"primary_key"`
	KodeKelas string `json:"kode_kelas"`
}

func (Kelas) TableName() string {
	return "kelas"
}

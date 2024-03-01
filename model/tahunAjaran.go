package model

type TahunAjaran struct {
	ID           int    `json:"id" gorm:"primary_key"`
	Tahun        string `json:"tahun"`
	Semester     string `json:"semester"`
	BulanMulai   string `json:"bulan_mulai"`
	BulanSelesai string `json:"bulan_selesai"`
}

func (TahunAjaran) TableName() string {
	return "tahun_ajaran"
}

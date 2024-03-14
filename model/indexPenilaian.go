package model

type IndexPenilaian struct {
	ID         int     `json:"id" gorm:"primary_key"`
	Grade      string  `json:"grade"`
	BatasAwal  float64 `json:"batas_awal"`
	BatasAkhir float64 `json:"batas_akhir"`
}

func (IndexPenilaian) TableName() string {
	return "index_penilaian"
}

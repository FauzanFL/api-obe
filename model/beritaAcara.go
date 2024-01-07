package model

type BeritaAcara struct {
	ID          int     `json:"id" gorm:"primary_key"`
	TahunAjaran string  `json:"tahun_ajaran"`
	Dosen       string  `json:"dosen"`
	NIM         string  `json:"nim"`
	Assessment  string  `json:"assessment"`
	Nilai       float64 `json:"nilai"`
}

func (BeritaAcara) TableName() string {
	return "berita_acara"
}

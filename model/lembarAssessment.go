package model

type LembarAssessment struct {
	ID        int     `json:"id" gorm:"primary_key"`
	Nama      string  `json:"nama"`
	Deskripsi string  `json:"deskripsi"`
	Bobot     float64 `json:"bobot"`
	CLOId     int     `json:"clo_id"`
	JenisId   int     `json:"jenis_id"`
}

type LembarAssessmentWithJenis struct {
	ID        int     `json:"id" gorm:"primary_key"`
	Nama      string  `json:"nama"`
	Deskripsi string  `json:"deskripsi"`
	Bobot     float64 `json:"bobot"`
	CLOId     int     `json:"clo_id"`
	JenisId   int     `json:"jenis_id"`
	Jenis     string  `json:"jenis"`
}

func (LembarAssessment) TableName() string {
	return "lembar_assessment"
}

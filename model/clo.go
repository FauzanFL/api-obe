package model

type CLO struct {
	ID        int     `json:"id" gorm:"primary_key"`
	PLOId     int     `json:"plo_id"`
	Nama      string  `json:"nama"`
	Deskripsi string  `json:"deskripsi"`
	Bobot     float64 `json:"bobot"`
	MkId      int     `json:"mk_id"`
}

type CLOWithAssessment struct {
	ID          int                         `json:"id" gorm:"primary_key"`
	PLOId       int                         `json:"plo_id"`
	Nama        string                      `json:"nama"`
	Deskripsi   string                      `json:"deskripsi"`
	Bobot       float64                     `json:"bobot"`
	MkId        int                         `json:"mk_id"`
	Assessments []LembarAssessmentWithJenis `json:"assessment"`
}

func (CLO) TableName() string {
	return "clo"
}

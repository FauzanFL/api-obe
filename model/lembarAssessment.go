package model

type LembarAssessment struct {
	ID      int     `json:"id" gorm:"primary_key"`
	Nama    string  `json:"nama"`
	Bobot   float64 `json:"bobot"`
	CLOId   int     `json:"clo_id"`
	JenisId int     `json:"jenis_id"`
}

func (LembarAssessment) TableName() string {
	return "lembar_assessment"
}

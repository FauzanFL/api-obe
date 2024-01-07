package model

type JenisAssessment struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Nama string `json:"nama"`
}

func (JenisAssessment) TableName() string {
	return "jenis_assessment"
}

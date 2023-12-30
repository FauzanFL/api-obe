package model

type JenisAssessment struct {
	ID int `json:"id" gorm:"primary_key"`
}

func (JenisAssessment) TableName() string {
	return "jenis_assessment"
}

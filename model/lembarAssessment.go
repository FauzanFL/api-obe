package model

type LembarAssessment struct {
	ID int `json:"id" gorm:"primary_key"`
}

func (LembarAssessment) TableName() string {
	return "lembar_assessment"
}

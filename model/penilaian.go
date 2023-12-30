package model

type Penilaian struct {
	ID int `json:"id" gorm:"primary_key"`
}

func (Penilaian) TableName() string {
	return "penilaian"
}

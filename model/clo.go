package model

type CLO struct {
	ID int `json:"id" gorm:"primary_key"`
}

func (CLO) TableName() string {
	return "clo"
}

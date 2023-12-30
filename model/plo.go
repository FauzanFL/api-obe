package model

type PLO struct {
	ID int `json:"id" gorm:"primary_key"`
}

func (PLO) TableName() string {
	return "plo"
}

package model

type Kelas struct {
	ID int `json:"id" gorm:"primary_key"`
}

func (Kelas) TableName() string {
	return "kelas"
}

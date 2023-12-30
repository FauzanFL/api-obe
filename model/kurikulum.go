package model

type Kurikulum struct {
	ID int `json:"id" gorm:"primary_key"`
}

func (Kurikulum) TableName() string {
	return "kurikulum"
}

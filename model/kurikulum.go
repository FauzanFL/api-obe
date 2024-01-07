package model

type Kurikulum struct {
	ID int `json:"id" gorm:"primary_key"`
	Nama string `json:"nama"`
}

func (Kurikulum) TableName() string {
	return "kurikulum"
}

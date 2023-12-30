package model

type TahunAjaran struct {
	ID int `json:"id" gorm:"primary_key"`
}

func (TahunAjaran) TableName() string {
	return "tahun_ajaran"
}

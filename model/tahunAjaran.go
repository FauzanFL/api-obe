package model

type TahunAjaran struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Tahun    string `json:"tahun"`
	Semester string `json:"semester"`
}

func (TahunAjaran) TableName() string {
	return "tahun_ajaran"
}

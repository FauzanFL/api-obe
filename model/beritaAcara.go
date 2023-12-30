package model

type BeritaAcara struct {
	ID int `json:"id" gorm:"primary_key"`
}

func (BeritaAcara) TableName() string {
	return "berita_acara"
}

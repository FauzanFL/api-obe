package model

type PLO_CLO struct {
	ID    int `json:"id" gorm:"primary_key"`
	PLOId int `json:"plo_id"`
	CLOId int `json:"clo_id"`
}

func (PLO_CLO) TableName() string {
	return "plo_clo"
}

package model

type PerancanganObe struct {
	ID int `json:"id" gorm:"primary_key"`
}

func (PerancanganObe) TableName() string {
	return "perancangan_obe"
}

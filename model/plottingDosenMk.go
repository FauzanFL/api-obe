package model

type PlottingDosenMk struct {
	ID int `json:"id" gorm:"primary_key"`
}

func (PlottingDosenMk) TableName() string {
	return "plotting_dosen_mk"
}

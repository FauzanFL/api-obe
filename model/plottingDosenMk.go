package model

type PlottingDosenMk struct {
	ID      int `json:"id" gorm:"primary_key"`
	MKId    int `json:"mk_id"`
	DosenId int `json:"dosen_id"`
}

func (PlottingDosenMk) TableName() string {
	return "plotting_dosen_mk"
}

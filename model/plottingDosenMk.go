package model

type PlottingDosenMk struct {
	ID      int `json:"id" gorm:"primary_key"`
	MKId    int `json:"mk_id"`
	DosenId int `json:"dosen_id"`
	KelasId int `json:"kelas_id"`
}

type PlotData struct {
	ID         int    `json:"id"`
	MataKuliah string `json:"mata_kuliah"`
	Dosen      string `json:"dosen"`
	Kelas      string `json:"kelas"`
}

func (PlottingDosenMk) TableName() string {
	return "plotting_dosen_mk"
}

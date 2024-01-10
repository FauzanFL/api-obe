package model

type MataKuliah struct {
	ID        int    `json:"id" gorm:"primary_key"`
	KodeMk    string `json:"kode_mk"`
	Nama      string `json:"nama"`
	Deskripsi string `json:"deskripsi"`
	Sks       int    `json:"sks"`
	Semester  int    `json:"semester"`
	Prasyarat string `json:"prasyarat"`
	OBEId     int    `json:"obe_id"`
}

type RPS struct {
	NamaMataKuliah      string  `json:"nama_mata_kuliah"`
	KodeMataKuliah      string  `json:"kode_mata_kuliah"`
	SKS                 int     `json:"sks"`
	Semester            int     `json:"semester"`
	Prodi               string  `json:"prodi"`
	Prasyarat           string  `json:"prasyarat"`
	DeskripsiMataKuliah string  `json:"deskripsi_mata_kuliah"`
	DosenPengampu       []Dosen `json:"dosen_pengampu"`
	PLO                 []PLO   `json:"plo"`
	CLO                 []CLO   `json:"clo"`
}

func (MataKuliah) TableName() string {
	return "mata_kuliah"
}

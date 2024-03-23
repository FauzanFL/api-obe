package model

type BeritaAcara struct {
	ID          int    `json:"id" gorm:"primary_key"`
	MataKuliah  string `json:"mata_kuliah"`
	Dosen       string `json:"dosen"`
	Kelas       string `json:"kelas"`
	Nilai       string `json:"nilai"`
	PenilaianId int    `json:"penilaian_id"`
}

type BeritaAcaraResp struct {
	ID          int              `json:"id" gorm:"primary_key"`
	MataKuliah  MataKuliah       `json:"mata_kuliah"`
	Dosen       Dosen            `json:"dosen"`
	Kelas       Kelas            `json:"kelas"`
	Nilai       []NilaiMahasiswa `json:"nilai"`
	PenilaianId int              `json:"penilaian_id"`
}

func (BeritaAcara) TableName() string {
	return "berita_acara"
}

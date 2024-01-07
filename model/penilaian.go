package model

type Penilaian struct {
	ID            int     `json:"id" gorm:"primary_key"`
	Nilai         float64 `json:"nilai"`
	AssessmentId  int     `json:"assessment_id"`
	MhsId         int     `json:"mhs_id"`
	DosenId       int     `json:"dosen_id"`
	TahunAjaranId int     `json:"tahun_ajaran_id"`
}

func (Penilaian) TableName() string {
	return "penilaian"
}

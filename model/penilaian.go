package model

type Penilaian struct {
	ID      int    `json:"id" gorm:"primary_key"`
	Nilai   string `json:"nilai"`
	Status  string `json:"status"`
	MkId    int    `json:"mk_id"`
	KelasId int    `json:"kelas_id"`
}

type PenilaianResp struct {
	ID      int              `json:"id" gorm:"primary_key"`
	Nilai   []NilaiMahasiswa `json:"nilai"`
	Status  string           `json:"status"`
	MkId    int              `json:"mk_id"`
	KelasId int              `json:"kelas_id"`
}

type NilaiMahasiswa struct {
	NIM             string            `json:"nim"`
	Nama            string            `json:"nama"`
	NilaiAssessment []NilaiAssessment `json:"nilai_assessment"`
}

type NilaiAssessment struct {
	AssessmentId int     `json:"assessment_id"`
	Nilai        float64 `json:"nilai"`
}

type PenilaianData struct {
	CLOAsessment []CLOWithAssessment `json:"clo_assessment"`
	Penilaian    PenilaianResp       `json:"penilaian"`
}

func (Penilaian) TableName() string {
	return "penilaian"
}

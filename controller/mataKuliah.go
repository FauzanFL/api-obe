package controller

import (
	"api-obe/model"
	repo "api-obe/repository"
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MataKuliahController interface {
	GetMataKuliah(c *gin.Context)
	GetMataKuliahById(c *gin.Context)
	GetMataKuliahActiveByTahunId(c *gin.Context)
	SearchMataKuliahActiveByTahunId(c *gin.Context)
	CreateMataKuliah(c *gin.Context)
	UpdateMataKuliah(c *gin.Context)
	DeleteMataKuliah(c *gin.Context)
	GetRPS(c *gin.Context)
}

type mataKuliahController struct {
	mataKuliahRepo  repo.MataKuliahRepository
	plottingRepo    repo.PlottingDosenMkRepository
	dosenRepo       repo.DosenRepository
	assessmentRepo  repo.LembarAssessmentRepository
	perancanganRepo repo.PerancanganObeRepository
}

func NewMataKuliahController(mataKuliahRepo repo.MataKuliahRepository, plottingRepo repo.PlottingDosenMkRepository, dosenRepo repo.DosenRepository, assessmentRepo repo.LembarAssessmentRepository, perancanganRepo repo.PerancanganObeRepository) MataKuliahController {
	return &mataKuliahController{
		mataKuliahRepo,
		plottingRepo,
		dosenRepo,
		assessmentRepo,
		perancanganRepo,
	}
}

func (m *mataKuliahController) GetMataKuliah(c *gin.Context) {
	mataKuliah, err := m.mataKuliahRepo.GetMataKuliah()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mataKuliah)
}

func (m *mataKuliahController) GetMataKuliahById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mataKuliah, err := m.mataKuliahRepo.GetMataKuliahById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mataKuliah)
}

func (m *mataKuliahController) GetMataKuliahActiveByTahunId(c *gin.Context) {
	tahunId, err := strconv.Atoi(c.Param("tahunId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	obe, err := m.perancanganRepo.GetActivePerancanganObe()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	mataKuliah, err := m.mataKuliahRepo.GetMataKuliahByObeIdAndTahunId(obe.ID, tahunId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mataKuliah)
}

func (m *mataKuliahController) SearchMataKuliahActiveByTahunId(c *gin.Context) {
	keyword := c.Query("keyword")
	tahunId, err := strconv.Atoi(c.Param("tahunId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	obe, err := m.perancanganRepo.GetActivePerancanganObe()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	mataKuliah, err := m.mataKuliahRepo.GetMataKuliahByObeIdTahunIdAndKeyword(obe.ID, tahunId, keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mataKuliah)
}

func (m *mataKuliahController) CreateMataKuliah(c *gin.Context) {
	var body struct {
		KodeMk        string `json:"kode_mk" binding:"required"`
		Nama          string `json:"nama" binding:"required"`
		Deskripsi     string `json:"deskripsi" binding:"required"`
		Sks           int    `json:"sks" binding:"required"`
		Semester      int    `json:"semester" binding:"required"`
		Prasyarat     string `json:"prasyarat"`
		TahunAjaranId int    `json:"tahun_ajaran_id" binding:"required"`
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if body.KodeMk == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "kode_mk can't be empty"})
		return
	}
	if body.Nama == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nama can't be empty"})
		return
	}
	if body.Deskripsi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "deskripsi can't be empty"})
		return
	}
	if body.Sks == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sks can't be empty"})
		return
	}
	if body.Semester == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "semester can't be empty"})
		return
	}

	obe, err := m.perancanganRepo.GetActivePerancanganObe()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var mataKuliah model.MataKuliah
	mataKuliah.KodeMk = body.KodeMk
	mataKuliah.Nama = body.Nama
	mataKuliah.Deskripsi = body.Deskripsi
	mataKuliah.Sks = body.Sks
	mataKuliah.Semester = body.Semester
	mataKuliah.Prasyarat = body.Prasyarat
	mataKuliah.OBEId = obe.ID
	mataKuliah.TahunAjaranId = body.TahunAjaranId
	if err := m.mataKuliahRepo.CreateMataKuliah(mataKuliah); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Mata kuliah added successfully"})
}

func (m *mataKuliahController) UpdateMataKuliah(c *gin.Context) {
	var body struct {
		KodeMk        string `json:"kode_mk" binding:"required"`
		Nama          string `json:"nama" binding:"required"`
		Deskripsi     string `json:"deskripsi" binding:"required"`
		Sks           int    `json:"sks" binding:"required"`
		Semester      int    `json:"semester" binding:"required"`
		Prasyarat     string `json:"prasyarat"`
		TahunAjaranId int    `json:"tahun_ajaran_id" binding:"required"`
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if body.KodeMk == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "kode_mk can't be empty"})
		return
	}
	if body.Nama == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nama can't be empty"})
		return
	}
	if body.Deskripsi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "deskripsi can't be empty"})
		return
	}
	if body.Sks == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sks can't be empty"})
		return
	}
	if body.Semester == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "semester can't be empty"})
		return
	}

	obe, err := m.perancanganRepo.GetActivePerancanganObe()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var mataKuliah model.MataKuliah
	mataKuliah.KodeMk = body.KodeMk
	mataKuliah.Nama = body.Nama
	mataKuliah.Deskripsi = body.Deskripsi
	mataKuliah.Sks = body.Sks
	mataKuliah.Semester = body.Semester
	mataKuliah.Prasyarat = body.Prasyarat
	mataKuliah.OBEId = obe.ID
	mataKuliah.TahunAjaranId = body.TahunAjaranId
	mataKuliah.ID = id
	if err := m.mataKuliahRepo.UpdateMataKuliah(mataKuliah); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Mata kuliah updated successfully"})
}

func (m *mataKuliahController) DeleteMataKuliah(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := m.mataKuliahRepo.DeleteMataKuliah(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Mata kuliah deleted successfully"})
}

func (m *mataKuliahController) GetRPS(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mk, err := m.mataKuliahRepo.GetMataKuliahById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	dosenPlotting, err := m.plottingRepo.GetPlottingDosenByMkId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var dosenMk []model.Dosen
	for _, v := range dosenPlotting {
		dosen, err := m.dosenRepo.GetDosenById(v.DosenId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		dosenMk = append(dosenMk, dosen)
	}

	plo, err := m.mataKuliahRepo.GetPLOFromMataKuliah(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	clo, err := m.mataKuliahRepo.GetCLOFromMataKuliah(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	sort.Slice(dosenMk, func(i, j int) bool {
		return dosenMk[i].ID < dosenMk[j].ID
	})

	for i := 0; i < len(dosenMk)-1; i++ {
		if dosenMk[i] == dosenMk[i+1] {
			dosenMk = dosenMk[:i+1]
			break
		}
	}

	var assessments []model.LembarAssessmentWithJenis
	for _, v := range clo {
		assessment, err := m.assessmentRepo.GetLembarAssessmentByCloId(v.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		assessments = append(assessments, assessment...)
	}

	var rps model.RPS
	rps.NamaMataKuliah = mk.Nama
	rps.KodeMataKuliah = mk.KodeMk
	rps.SKS = mk.Sks
	rps.Semester = mk.Semester
	rps.Prasyarat = mk.Prasyarat
	rps.DeskripsiMataKuliah = mk.Deskripsi
	rps.DosenPengampu = dosenMk
	rps.Prodi = "Rekayasa Perangkat Lunak"
	rps.PLO = plo
	rps.CLO = clo
	rps.LembarAssessment = assessments

	c.JSON(http.StatusOK, rps)
}

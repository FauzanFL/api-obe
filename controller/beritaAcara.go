package controller

import (
	"api-obe/model"
	repo "api-obe/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BeritaAcaraController interface {
	GetBeritaAcara(c *gin.Context)
	GetBeritaAcaraByPenilaian(c *gin.Context)
	CreateBeritaAcara(c *gin.Context)
	DeleteBeritaAcara(c *gin.Context)
}

type beritaAcaraController struct {
	beritaAcaraRepo repo.BeritaAcaraRepository
	dosenRepo       repo.DosenRepository
	tahunAjaranRepo repo.TahunAjaranRepository
	assessmentRepo  repo.LembarAssessmentRepository
	penilaianRepo   repo.PenilaianRepository
}

func NewBeritaAcaraController(beritaAcaraRepo repo.BeritaAcaraRepository, dosenRepo repo.DosenRepository,
	tahunAjaranRepo repo.TahunAjaranRepository,
	assessmentRepo repo.LembarAssessmentRepository, penilaianRepo repo.PenilaianRepository) BeritaAcaraController {
	return &beritaAcaraController{beritaAcaraRepo, dosenRepo,
		tahunAjaranRepo,
		assessmentRepo,
		penilaianRepo}
}

func (b *beritaAcaraController) GetBeritaAcara(c *gin.Context) {
	beritaAcara, err := b.beritaAcaraRepo.GetBeritaAcara()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, beritaAcara)
}

func (b *beritaAcaraController) GetBeritaAcaraByPenilaian(c *gin.Context) {
	penilaianId, err := strconv.Atoi(c.Param("penilaianId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	beritaAcara, err := b.beritaAcaraRepo.GetBeritaAcaraByPenilaianId(penilaianId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	kelas := model.Kelas{}
	matakuliah := model.MataKuliah{}
	dosen := model.Dosen{}
	nilai := []model.NilaiMahasiswa{}

	err = json.Unmarshal([]byte(beritaAcara.Kelas), &kelas)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = json.Unmarshal([]byte(beritaAcara.MataKuliah), &matakuliah)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = json.Unmarshal([]byte(beritaAcara.Dosen), &dosen)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = json.Unmarshal([]byte(beritaAcara.Nilai), &nilai)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	beritaAcaraResp := model.BeritaAcaraResp{}
	beritaAcaraResp.ID = beritaAcara.ID
	beritaAcaraResp.Kelas = kelas
	beritaAcaraResp.Dosen = dosen
	beritaAcaraResp.MataKuliah = matakuliah
	beritaAcaraResp.Nilai = nilai
	beritaAcaraResp.PenilaianId = beritaAcara.PenilaianId

	c.JSON(http.StatusOK, beritaAcaraResp)
}

func (b *beritaAcaraController) CreateBeritaAcara(c *gin.Context) {
	var body struct {
		MataKuliah  model.MataKuliah       `json:"mata_kuliah" binding:"required"`
		Kelas       model.Kelas            `json:"kelas" binding:"required"`
		Nilai       []model.NilaiMahasiswa `json:"nilai" binding:"required"`
		PenilaianId int                    `json:"penilaian_id" binding:"required"`
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if body.MataKuliah == (model.MataKuliah{}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "matakuliah can't be empty"})
		return
	}
	if body.Kelas == (model.Kelas{}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "kelas can't be empty"})
		return
	}
	if len(body.Nilai) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nilai can't be empty"})
		return
	}
	if body.PenilaianId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "penilaian_id can't be empty"})
		return
	}

	matakuliah, err := json.Marshal(body.MataKuliah)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userGet, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not exist"})
		return
	}

	user := userGet.(model.User)

	dosenGet, err := b.dosenRepo.GetDosenByUserId(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	dosen, err := json.Marshal(dosenGet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	kelas, err := json.Marshal(body.Kelas)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	nilai, err := json.Marshal(body.Nilai)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var beritaAcara model.BeritaAcara
	beritaAcara.MataKuliah = string(matakuliah)
	beritaAcara.Dosen = string(dosen)
	beritaAcara.Kelas = string(kelas)
	beritaAcara.Nilai = string(nilai)
	beritaAcara.PenilaianId = body.PenilaianId
	if err := b.beritaAcaraRepo.CreateBeritaAcara(beritaAcara); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	penilaian, err := b.penilaianRepo.GetPenilaianById(body.PenilaianId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := b.penilaianRepo.UpdateStatusPenilaian(penilaian.ID, "final"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Berita acara added successfully"})
}

func (b *beritaAcaraController) DeleteBeritaAcara(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	beritaAcara, err := b.beritaAcaraRepo.GetBeritaAcaraById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := b.penilaianRepo.UpdateStatusPenilaian(beritaAcara.PenilaianId, "draft"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = b.beritaAcaraRepo.DeleteBeritaAcara(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Berita acara deleted successfully"})
}

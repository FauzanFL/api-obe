package controller

import (
	"api-obe/model"
	repo "api-obe/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BeritaAcaraController interface {
	GetBeritaAcara(c *gin.Context)
	CreateBeritaAcara(c *gin.Context)
	CreateManyBeritaAcara(c *gin.Context)
	DeleteBeritaAcara(c *gin.Context)
}

type beritaAcaraController struct {
	beritaAcaraRepo repo.BeritaAcaraRepository
	dosenRepo       repo.DosenRepository
	tahunAjaranRepo repo.TahunAjaranRepository
	assessmentRepo  repo.LembarAssessmentRepository
	mahasiswaRepo   repo.MahasiswaRepository
	penilaianRepo   repo.PenilaianRepository
}

func NewBeritaAcaraController(beritaAcaraRepo repo.BeritaAcaraRepository, dosenRepo repo.DosenRepository,
	tahunAjaranRepo repo.TahunAjaranRepository,
	assessmentRepo repo.LembarAssessmentRepository,
	mahasiswaRepo repo.MahasiswaRepository, penilaianRepo repo.PenilaianRepository) BeritaAcaraController {
	return &beritaAcaraController{beritaAcaraRepo, dosenRepo,
		tahunAjaranRepo,
		assessmentRepo,
		mahasiswaRepo,
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

func (b *beritaAcaraController) CreateBeritaAcara(c *gin.Context) {
	var body struct {
		TahunAjaran string  `json:"tahun_ajaran" binding:"required"`
		Dosen       string  `json:"dosen" binding:"required"`
		NIM         string  `json:"nim" binding:"required"`
		Assessment  string  `json:"assessment" binding:"required"`
		Nilai       float64 `json:"nilai" binding:"required"`
		PenilaianId int     `json:"penilaian_id" binding:"required"`
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if body.TahunAjaran == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tahun_ajaran can't be empty"})
		return
	}
	if body.Dosen == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dosen can't be empty"})
		return
	}
	if body.NIM == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nim can't be empty"})
		return
	}
	if body.Assessment == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "assessment can't be empty"})
		return
	}
	if body.PenilaianId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "penilaian_id can't be empty"})
		return
	}

	var beritaAcara model.BeritaAcara
	beritaAcara.TahunAjaran = body.TahunAjaran
	beritaAcara.Dosen = body.Dosen
	beritaAcara.NIM = body.NIM
	beritaAcara.Assessment = body.Assessment
	beritaAcara.Nilai = body.Nilai
	if err := b.beritaAcaraRepo.CreateBeritaAcara(beritaAcara); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Berita acara added successfully"})
}

func (b *beritaAcaraController) CreateManyBeritaAcara(c *gin.Context) {
	var body []struct {
		PenilaianId   int     `json:"id" binding:"required"`
		Nilai         float64 `json:"nilai" binding:"required"`
		AssessmentId  int     `json:"assessment_id" binding:"required"`
		MhsId         int     `json:"mhs_id" binding:"required"`
		TahunAjaranId int     `json:"tahun_ajaran_id" binding:"required"`
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(body) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No data provided"})
		return
	}

	listBeritaAcara := []model.BeritaAcara{}

	for _, v := range body {
		if v.TahunAjaranId == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "tahun_ajaran can't be empty"})
			return
		}
		if v.Nilai == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "nim can't be empty"})
			return
		}
		if v.AssessmentId == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "assessment can't be empty"})
			return
		}
		if v.MhsId == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "assessment can't be empty"})
			return
		}
		if v.PenilaianId == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "penilaian_id can't be empty"})
			return
		}

		userGet, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not exist"})
			return
		}

		user := userGet.(model.User)

		dosen, err := b.dosenRepo.GetDosenByUserId(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		mhs, err := b.mahasiswaRepo.GetMahasiswaById(v.MhsId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		tahun, err := b.tahunAjaranRepo.GetTahunAjaranById(v.TahunAjaranId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		assessment, err := b.assessmentRepo.GetLembarAssessmentById(v.AssessmentId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := b.penilaianRepo.UpdateStatusToFinal(v.PenilaianId); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var beritaAcara model.BeritaAcara
		beritaAcara.TahunAjaran = fmt.Sprintf("%s %s", tahun.Tahun, tahun.Semester)
		beritaAcara.Dosen = dosen.KodeDosen
		beritaAcara.NIM = mhs.NIM
		beritaAcara.Assessment = assessment.Nama
		beritaAcara.Nilai = v.Nilai
		beritaAcara.PenilaianId = v.PenilaianId

		beritaAcaraCheck, _ := b.beritaAcaraRepo.GetBeritaAcaraByPenialainId(v.PenilaianId)

		if beritaAcaraCheck.ID == 0 {
			listBeritaAcara = append(listBeritaAcara, beritaAcara)
		}
	}

	if err := b.beritaAcaraRepo.CreateManyBeritaAcara(listBeritaAcara); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Many berita acara added successfully"})
}

func (b *beritaAcaraController) DeleteBeritaAcara(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = b.beritaAcaraRepo.DeleteBeritaAcara(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "Berita acara deleted successfully"})
}

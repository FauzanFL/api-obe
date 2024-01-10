package controller

import (
	"api-obe/model"
	repo "api-obe/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BeritaAcaraController interface {
	GetBeritaAcara(c *gin.Context)
	CreateBeritaAcara(c *gin.Context)
	DeleteBeritaAcara(c *gin.Context)
}

type beritaAcaraController struct {
	beritaAcaraRepo repo.BeritaAcaraRepository
}

func NewBeritaAcaraController(beritaAcaraRepo repo.BeritaAcaraRepository) BeritaAcaraController {
	return &beritaAcaraController{beritaAcaraRepo}
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

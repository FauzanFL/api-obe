package controller

import (
	"api-obe/model"
	repo "api-obe/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LembarAssessmentController interface {
	GetLembarAssessment(c *gin.Context)
	GetLembarAssessmentById(c *gin.Context)
	GetLembarAssessmentByCloId(c *gin.Context)
	SearchLembarAssessment(c *gin.Context)
	CreateLembarAssessment(c *gin.Context)
	UpdateLembarAssessment(c *gin.Context)
	DeleteLembarAssessment(c *gin.Context)
}

type lembarAssessmentController struct {
	lembarAssessmentRepo repo.LembarAssessmentRepository
}

func NewLembarAssessmentController(lembarAssessmentRepo repo.LembarAssessmentRepository) LembarAssessmentController {
	return &lembarAssessmentController{
		lembarAssessmentRepo,
	}
}

func (l *lembarAssessmentController) GetLembarAssessment(c *gin.Context) {
	lembarAssessment, err := l.lembarAssessmentRepo.GetLembarAssessment()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lembarAssessment)
}

func (l *lembarAssessmentController) GetLembarAssessmentById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lembarAssessment, err := l.lembarAssessmentRepo.GetLembarAssessmentById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lembarAssessment)
}

func (l *lembarAssessmentController) GetLembarAssessmentByCloId(c *gin.Context) {
	cloId, err := strconv.Atoi(c.Param("cloId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lembarAssessment, err := l.lembarAssessmentRepo.GetLembarAssessmentByCloId(cloId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lembarAssessment)
}

func (l *lembarAssessmentController) SearchLembarAssessment(c *gin.Context) {
	keyword := c.Query("keyword")
	cloId, err := strconv.Atoi(c.Param("cloId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lembarAssessment, err := l.lembarAssessmentRepo.GetLembarAssessmentByCloIdAndKeyword(cloId, keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lembarAssessment)
}

func (l *lembarAssessmentController) CreateLembarAssessment(c *gin.Context) {
	var body struct {
		Nama      string  `json:"nama" binding:"required"`
		Deskripsi string  `json:"deskripsi" binding:"required"`
		Bobot     float64 `json:"bobot" binding:"required"`
		CLOId     int     `json:"clo_id" binding:"required"`
		JenisId   int     `json:"jenis_id" binding:"required"`
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if body.Nama == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nama tidak boleh kosong"})
		return
	}
	if body.Bobot == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bobot tidak boleh kosong"})
		return
	}
	if body.CLOId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "clo_id tidak boleh kosong"})
		return
	}
	if body.JenisId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "jenis_id tidak boleh kosong"})
		return
	}

	var lembarAssessment model.LembarAssessment
	lembarAssessment.Nama = body.Nama
	lembarAssessment.Deskripsi = body.Deskripsi
	lembarAssessment.Bobot = body.Bobot
	lembarAssessment.CLOId = body.CLOId
	lembarAssessment.JenisId = body.JenisId

	if err := l.lembarAssessmentRepo.CreateLembarAssessment(lembarAssessment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Lembar assessment added successfully"})
}

func (l *lembarAssessmentController) UpdateLembarAssessment(c *gin.Context) {
	var body struct {
		Nama      string  `json:"nama" binding:"required"`
		Deskripsi string  `json:"deskripsi" binding:"required"`
		Bobot     float64 `json:"bobot" binding:"required"`
		CLOId     int     `json:"clo_id" binding:"required"`
		JenisId   int     `json:"jenis_id" binding:"required"`
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
	if body.Nama == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nama can't be empty"})
		return
	}
	if body.Bobot == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bobot can't be empty"})
		return
	}
	if body.CLOId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "clo_id can't be empty"})
		return
	}
	if body.JenisId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "jenis_id can't be empty"})
		return
	}

	var lembarAssessment model.LembarAssessment
	lembarAssessment.Nama = body.Nama
	lembarAssessment.Deskripsi = body.Deskripsi
	lembarAssessment.Bobot = body.Bobot
	lembarAssessment.CLOId = body.CLOId
	lembarAssessment.JenisId = body.JenisId
	lembarAssessment.ID = id
	if err := l.lembarAssessmentRepo.UpdateLembarAssessment(lembarAssessment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Lembar assessment updated successfully"})
}

func (l *lembarAssessmentController) DeleteLembarAssessment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = l.lembarAssessmentRepo.DeleteLembarAssessment(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Lembar assessment delted successfully"})
}

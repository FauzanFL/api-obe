package controller

import (
	"api-obe/model"
	repo "api-obe/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JenisAssessmentController interface {
	GetJenisAssessment(c *gin.Context)
	CreateJenisAssessment(c *gin.Context)
	UpdateJenisAssessment(c *gin.Context)
	DeleteJenisAssessment(c *gin.Context)
}

type jenisAssessmentController struct {
	jenisAssessmentRepo repo.JenisAssessmentRepository
}

func NewJenisAssessmentController(jenisAssessmentRepo repo.JenisAssessmentRepository) JenisAssessmentController {
	return &jenisAssessmentController{
		jenisAssessmentRepo,
	}
}

func (ja *jenisAssessmentController) GetJenisAssessment(c *gin.Context) {
	jenisAssessment, err := ja.jenisAssessmentRepo.GetJenisAssessment()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, jenisAssessment)
}

func (ja *jenisAssessmentController) CreateJenisAssessment(c *gin.Context) {
	var body struct {
		Nama string `json:"nama" binding:"required"`
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if body.Nama == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nama can't be empty"})
		return
	}

	var jenisAssessment model.JenisAssessment
	jenisAssessment.Nama = body.Nama
	if err := ja.jenisAssessmentRepo.CreateJenisAssessment(jenisAssessment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Jenis assessment added successfully"})
}

func (ja *jenisAssessmentController) UpdateJenisAssessment(c *gin.Context) {
	var body struct {
		Nama string `json:"nama" binding:"required"`
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

	var jenisAssessment model.JenisAssessment
	jenisAssessment.ID = id
	jenisAssessment.Nama = body.Nama
	if err := ja.jenisAssessmentRepo.UpdateJenisAssessment(jenisAssessment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Jenis assessment updated successfully"})
}

func (ja *jenisAssessmentController) DeleteJenisAssessment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = ja.jenisAssessmentRepo.DeleteJenisAssessment(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Jenis assessment deleted successfully"})
}

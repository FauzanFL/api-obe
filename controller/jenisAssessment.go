package controller

import (
	repo "api-obe/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JenisAssessmentController interface {
	GetJenisAssessment(c *gin.Context)
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

package controller

import (
	repo "api-obe/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LembarAssessmentController interface {
	GetLembarAssessment(c *gin.Context)
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

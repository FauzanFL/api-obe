package controller

import (
	repo "api-obe/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PerancanganObeController interface {
	GetPerancanganObe(c *gin.Context)
}

type perancanganObeController struct {
	perancanganObeRepo repo.PerancanganObeRepository
}

func NewPerancanganObeController(perancanganObeRepo repo.PerancanganObeRepository) PerancanganObeController {
	return &perancanganObeController{
		perancanganObeRepo,
	}
}

func (m *perancanganObeController) GetPerancanganObe(c *gin.Context) {
	perancanganObe, err := m.perancanganObeRepo.GetPerancanganObe()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, perancanganObe)
}

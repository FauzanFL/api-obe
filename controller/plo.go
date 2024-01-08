package controller

import (
	repo "api-obe/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PloController interface {
	GetPlo(c *gin.Context)
}

type ploController struct {
	ploRepo repo.PloRepository
}

func NewPloController(ploRepo repo.PloRepository) PloController {
	return &ploController{
		ploRepo,
	}
}

func (m *ploController) GetPlo(c *gin.Context) {
	plo, err := m.ploRepo.GetPlo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, plo)
}

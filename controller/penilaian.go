package controller

import (
	repo "api-obe/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PenilaianController interface {
	GetPenilaian(c *gin.Context)
}

type penilaianController struct {
	penilaianRepo repo.PenilaianRepository
}

func NewPenilaianController(penilaianRepo repo.PenilaianRepository) PenilaianController {
	return &penilaianController{
		penilaianRepo,
	}
}

func (m *penilaianController) GetPenilaian(c *gin.Context) {
	penilaian, err := m.penilaianRepo.GetPenilaian()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, penilaian)
}

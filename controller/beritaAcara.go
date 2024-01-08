package controller

import (
	repo "api-obe/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BeritaAcaraController interface {
	GetBeritaAcara(c *gin.Context)
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

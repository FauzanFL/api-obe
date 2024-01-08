package controller

import (
	repo "api-obe/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TahunAjaranController interface {
	GetTahunAjaran(c *gin.Context)
}

type tahunAjaranController struct {
	tahunAjaranRepo repo.TahunAjaranRepository
}

func NewTahunAjaranController(tahunAjaranRepo repo.TahunAjaranRepository) TahunAjaranController {
	return &tahunAjaranController{
		tahunAjaranRepo,
	}
}

func (t *tahunAjaranController) GetTahunAjaran(c *gin.Context) {
	tahunAjaran, err := t.tahunAjaranRepo.GetTahunAjaran()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tahunAjaran)
}

package controller

import (
	repo "api-obe/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type KelasController interface {
	GetKelas(c *gin.Context)
}

type kelasController struct {
	kelasRepo repo.KelasRepository
}

func NewKelasController(kelasRepo repo.KelasRepository) KelasController {
	return &kelasController{
		kelasRepo,
	}
}

func (k *kelasController) GetKelas(c *gin.Context) {
	kelas, err := k.kelasRepo.GetKelas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, kelas)
}

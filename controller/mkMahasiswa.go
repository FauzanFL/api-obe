package controller

import (
	repo "api-obe/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MkMahasiswaController interface {
	GetMkMahasiswa(c *gin.Context)
}

type mkMahasiswaController struct {
	mkMahasiswaRepo repo.MkMahasiswaRepository
}

func NewMkMahasiswaController(mkMahasiswaRepo repo.MkMahasiswaRepository) MkMahasiswaController {
	return &mkMahasiswaController{
		mkMahasiswaRepo,
	}
}

func (m *mkMahasiswaController) GetMkMahasiswa(c *gin.Context) {
	mkMahasiswa, err := m.mkMahasiswaRepo.GetMkMahasiswa()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mkMahasiswa)
}

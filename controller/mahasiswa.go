package controller

import (
	repo "api-obe/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MahasiswaController interface {
	GetMahasiswa(c *gin.Context)
}

type mahasiswaController struct {
	mahasiswaRepo repo.MahasiswaRepository
}

func NewMahasiswaController(mahasiswaRepo repo.MahasiswaRepository) MahasiswaController {
	return &mahasiswaController{
		mahasiswaRepo,
	}
}

func (m *mahasiswaController) GetMahasiswa(c *gin.Context) {
	mahasiswa, err := m.mahasiswaRepo.GetMahasiswa()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mahasiswa)
}

package controller

import (
	repo "api-obe/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MahasiswaController interface {
	GetMahasiswa(c *gin.Context)
	GetMahasiswaByMataKuliah(c *gin.Context)
	GetMahasiswaByKelasMataKuliah(c *gin.Context)
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

func (m *mahasiswaController) GetMahasiswaByMataKuliah(c *gin.Context) {
	mkId, err := strconv.Atoi(c.Param("mkId"))
	mahasiswa, err := m.mahasiswaRepo.GetMahasiswaByMataKuliah(mkId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mahasiswa)
}

func (m *mahasiswaController) GetMahasiswaByKelasMataKuliah(c *gin.Context) {
	kelasId, err := strconv.Atoi(c.Param("kelasId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mkId, err := strconv.Atoi(c.Param("mkId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mahasiswa, err := m.mahasiswaRepo.GetMahasiswaByKelasMataKuliah(mkId, kelasId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mahasiswa)
}

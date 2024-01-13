package controller

import (
	"api-obe/model"
	repo "api-obe/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MkMahasiswaController interface {
	GetMkMahasiswa(c *gin.Context)
	CreateMkMahasiswa(c *gin.Context)
	DeleteMkMahasiswa(c *gin.Context)
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

func (m *mkMahasiswaController) CreateMkMahasiswa(c *gin.Context) {
	var body struct {
		MKId    int `json:"mk_id" binding:"required"`
		MhsId   int `json:"mhs_id" binding:"required"`
		KelasId int `json:"kelas_id" binding:"required"`
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if body.MKId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "mk_id can't be empty"})
	}
	if body.MhsId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "mhs_id can't be empty"})
	}

	var mkMahasiswa model.MkMahasiswa
	mkMahasiswa.MKId = body.MKId
	mkMahasiswa.MhsId = body.MhsId
	mkMahasiswa.KelasId = body.KelasId
	if err := m.mkMahasiswaRepo.CreateMkMahasiswa(mkMahasiswa); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Mk mahasiswa added successfully"})
}

func (m *mkMahasiswaController) DeleteMkMahasiswa(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = m.mkMahasiswaRepo.DeleteMkMahasiswa(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Mk mahasiswa deleted successfully"})
}

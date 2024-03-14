package controller

import (
	"api-obe/model"
	repo "api-obe/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type KelasController interface {
	GetKelas(c *gin.Context)
	GetKelasById(c *gin.Context)
	SearchKelas(c *gin.Context)
	CreateKelas(c *gin.Context)
	UpdateKelas(c *gin.Context)
	DeleteKelas(c *gin.Context)
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

func (k *kelasController) GetKelasById(c *gin.Context) {
	kelasId, err := strconv.Atoi(c.Param("kelasId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	kelas, err := k.kelasRepo.GetKelasById(kelasId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, kelas)
}

func (k *kelasController) SearchKelas(c *gin.Context) {
	keyword := c.Query("keyword")

	kelas, err := k.kelasRepo.GetKelasByKeyword(keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, kelas)
}

func (k *kelasController) CreateKelas(c *gin.Context) {
	var body struct {
		KodeKelas string `json:"kode_kelas" binding:"required"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if body.KodeKelas == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "kode kelas can't be empty"})
		return
	}

	var kelas model.Kelas
	kelas.KodeKelas = body.KodeKelas
	if err := k.kelasRepo.CreateKelas(kelas); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Kelas added successfully"})
}

func (k *kelasController) UpdateKelas(c *gin.Context) {
	var body struct {
		KodeKelas string `json:"kode_kelas" binding:"required"`
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if body.KodeKelas == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "kode kelas can't be empty"})
		return
	}

	var kelas model.Kelas
	kelas.ID = id
	kelas.KodeKelas = body.KodeKelas
	if err := k.kelasRepo.UpdateKelas(kelas); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Kelas updated successfully"})
}

func (k *kelasController) DeleteKelas(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := k.kelasRepo.DeleteKelas(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Kelas deleted successfully"})
}

package controller

import (
	"api-obe/model"
	repo "api-obe/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IndexPenilaianController interface {
	GetIndexPenilaian(c *gin.Context)
	GetIndexPenilaianById(c *gin.Context)
	GetIndexPenilaianByNilai(c *gin.Context)
	CreateIndexPenilaian(c *gin.Context)
	UpdateIndexPenilaian(c *gin.Context)
	DeleteIndexPenilaian(c *gin.Context)
}

type indexPenilaianController struct {
	indexPenilaianRepo repo.IndexPenilaianRepository
}

func NewIndexPenilaianController(indexPenilaianRepo repo.IndexPenilaianRepository) IndexPenilaianController {
	return &indexPenilaianController{indexPenilaianRepo}
}
func (i *indexPenilaianController) GetIndexPenilaian(c *gin.Context) {
	indexPenilaian, err := i.indexPenilaianRepo.GetIndexPenilaian()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, indexPenilaian)
}

func (i *indexPenilaianController) GetIndexPenilaianById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	indexPenilaian, err := i.indexPenilaianRepo.GetIndexPenilaianById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, indexPenilaian)
}

func (i *indexPenilaianController) GetIndexPenilaianByNilai(c *gin.Context) {
	nilai, err := strconv.ParseFloat(c.Query("nilai"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid nilai"})
	}

	indexPenilaian, err := i.indexPenilaianRepo.GetIndexPenilaianByNilai(nilai)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, indexPenilaian)
}

func (i *indexPenilaianController) CreateIndexPenilaian(c *gin.Context) {
	var body struct {
		Grade      string  `json:"grade" binding:"required"`
		BatasAwal  float64 `json:"batas_awal" binding:"required"`
		BatasAkhir float64 `json:"batas_akhir" binding:"required"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if body.Grade == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "grade can't be empty"})
		return
	}
	if body.BatasAwal == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "batas awal can't be empty"})
		return
	}
	if body.BatasAkhir == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "batas akhir can't be empty"})
		return
	}

	var indexPenilaian model.IndexPenilaian
	indexPenilaian.Grade = body.Grade
	indexPenilaian.BatasAwal = body.BatasAwal
	indexPenilaian.BatasAkhir = body.BatasAkhir

	err := i.indexPenilaianRepo.CreateIndexPenilaian(indexPenilaian)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Index penilaian added successfully"})
}

func (i *indexPenilaianController) UpdateIndexPenilaian(c *gin.Context) {
	var body struct {
		Grade      string  `json:"grade" binding:"required"`
		BatasAwal  float64 `json:"batas_awal" binding:"required"`
		BatasAkhir float64 `json:"batas_akhir" binding:"required"`
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if body.Grade == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "grade can't be empty"})
		return
	}
	if body.BatasAwal == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "batas awal can't be empty"})
		return
	}
	if body.BatasAkhir == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "batas akhir can't be empty"})
		return
	}

	var indexPenilaian model.IndexPenilaian
	indexPenilaian.ID = id
	indexPenilaian.Grade = body.Grade
	indexPenilaian.BatasAwal = body.BatasAwal
	indexPenilaian.BatasAkhir = body.BatasAkhir

	err = i.indexPenilaianRepo.UpdateIndexPenilaian(indexPenilaian)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Index penilaian updated successfully"})
}

func (i *indexPenilaianController) DeleteIndexPenilaian(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	err = i.indexPenilaianRepo.DeleteIndexPenilaian(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Index penilaian deleted successfully"})
}

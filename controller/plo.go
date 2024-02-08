package controller

import (
	"api-obe/model"
	repo "api-obe/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PloController interface {
	GetPlo(c *gin.Context)
	GetPloById(c *gin.Context)
	GetPloByObeId(c *gin.Context)
	SearchPloByObeId(c *gin.Context)
	CreatePlo(c *gin.Context)
	UpdatePlo(c *gin.Context)
	DeletePlo(c *gin.Context)
}

type ploController struct {
	ploRepo repo.PloRepository
}

func NewPloController(ploRepo repo.PloRepository) PloController {
	return &ploController{
		ploRepo,
	}
}

func (m *ploController) GetPlo(c *gin.Context) {
	plo, err := m.ploRepo.GetPlo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, plo)
}

func (m *ploController) GetPloById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plo, err := m.ploRepo.GetPloById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, plo)
}

func (m *ploController) GetPloByObeId(c *gin.Context) {
	obeId, err := strconv.Atoi(c.Param("obeId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plo, err := m.ploRepo.GetPloByObeId(obeId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, plo)
}

func (m *ploController) SearchPloByObeId(c *gin.Context) {
	keyword := c.Query("keyword")
	obeId, err := strconv.Atoi(c.Param("obeId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plo, err := m.ploRepo.GetPloByObeIdAndKeyword(obeId, keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, plo)
}

func (m *ploController) CreatePlo(c *gin.Context) {
	var body struct {
		Nama      string `json:"nama" binding:"required"`
		Deskripsi string `json:"deskripsi" binding:"required"`
		OBEId     int    `json:"obe_id" binding:"required"`
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if body.Nama == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nama can't be empty"})
		return
	}
	if body.Deskripsi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "deskripsi can't be empty"})
		return
	}
	if body.OBEId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "obe_id can't be empty"})
		return
	}

	var plo model.PLO
	plo.Nama = body.Nama
	plo.Deskripsi = body.Deskripsi
	plo.OBEId = body.OBEId

	if err := m.ploRepo.CreatePlo(plo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "PLO added successfully"})
}

func (m *ploController) UpdatePlo(c *gin.Context) {
	var body struct {
		Nama      string `json:"nama" binding:"required"`
		Deskripsi string `json:"deskripsi" binding:"required"`
		OBEId     int    `json:"obe_id" binding:"required"`
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
	if body.Nama == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nama can't be empty"})
		return
	}
	if body.Deskripsi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "deskripsi can't be empty"})
		return
	}
	if body.OBEId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "obe_id can't be empty"})
		return
	}

	var plo model.PLO
	plo.Nama = body.Nama
	plo.Deskripsi = body.Deskripsi
	plo.OBEId = body.OBEId
	plo.ID = id
	if err := m.ploRepo.UpdatePlo(plo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "PLO updated successfully"})
}

func (m *ploController) DeletePlo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := m.ploRepo.DeletePlo(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "PLO deleted successfully"})
}

package controller

import (
	"api-obe/model"
	repo "api-obe/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PerancanganObeController interface {
	GetPerancanganObe(c *gin.Context)
	SearchPerancanganObe(c *gin.Context)
	GetActivePerancanganObe(c *gin.Context)
	GetPerancanganObeById(c *gin.Context)
	CreatePerancanganObe(c *gin.Context)
	UpdatePerancanganObe(c *gin.Context)
	DeletePerancanganObe(c *gin.Context)
	ActivatePerancangan(c *gin.Context)
}

type perancanganObeController struct {
	perancanganObeRepo repo.PerancanganObeRepository
}

func NewPerancanganObeController(perancanganObeRepo repo.PerancanganObeRepository) PerancanganObeController {
	return &perancanganObeController{
		perancanganObeRepo,
	}
}

func (m *perancanganObeController) GetPerancanganObe(c *gin.Context) {
	perancanganObe, err := m.perancanganObeRepo.GetPerancanganObe()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, perancanganObe)
}

func (m *perancanganObeController) SearchPerancanganObe(c *gin.Context) {
	keyword := c.Query("keyword")
	perancanganObe, err := m.perancanganObeRepo.GetPerancanganObeByKeyword(keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, perancanganObe)
}

func (m *perancanganObeController) GetActivePerancanganObe(c *gin.Context) {
	perancanganObe, err := m.perancanganObeRepo.GetActivePerancanganObe()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, perancanganObe)
}

func (m *perancanganObeController) GetPerancanganObeById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	perancanganObe, err := m.perancanganObeRepo.GetPerancanganObeById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, perancanganObe)
}

func (m *perancanganObeController) CreatePerancanganObe(c *gin.Context) {
	var body struct {
		Nama        string `json:"nama" binding:"required"`
		KurikulumID int    `json:"kurikulum_id" binding:"required"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if body.Nama == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nama can't be empty"})
		return
	}
	if body.KurikulumID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "kurikulum_id can't be empty"})
		return
	}

	var perancanganObe model.PerancanganObe
	perancanganObe.Nama = body.Nama
	perancanganObe.KurikulumID = body.KurikulumID
	perancanganObe.Status = "inactive"
	if err := m.perancanganObeRepo.CreatePerancanganObe(perancanganObe); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Perancangan OBE added successfully"})
}

func (m *perancanganObeController) UpdatePerancanganObe(c *gin.Context) {
	var body struct {
		Nama        string `json:"nama" binding:"required"`
		KurikulumID int    `json:"kurikulum_id" binding:"required"`
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
	if body.KurikulumID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "kurikulum_id can't be empty"})
		return
	}

	var perancanganObe model.PerancanganObe
	perancanganObe.Nama = body.Nama
	perancanganObe.KurikulumID = body.KurikulumID
	perancanganObe.ID = id
	if err := m.perancanganObeRepo.UpdatePerancanganObe(perancanganObe); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Perancangan OBE updated successfully"})
}

func (m *perancanganObeController) DeletePerancanganObe(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	perancangan, err := m.perancanganObeRepo.GetPerancanganObeById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if perancangan.Status == "active" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't delete active perancangan"})
		return
	}

	if err := m.perancanganObeRepo.DeletePerancanganObe(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Perancangan OBE deleted successfully"})
}

func (m *perancanganObeController) ActivatePerancangan(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	perancangan, err := m.perancanganObeRepo.GetPerancanganObeById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if perancangan.Status == "active" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Perancangan already active"})
		return
	}

	if err := m.perancanganObeRepo.DiactivatePerancangan(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := m.perancanganObeRepo.UpdateStatusPerancangan(id, "active"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Perancangan OBE activated successfully"})
}

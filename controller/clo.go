package controller

import (
	"api-obe/model"
	repo "api-obe/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CloController interface {
	GetClo(c *gin.Context)
	GetCloById(c *gin.Context)
	CreateClo(c *gin.Context)
	UpdateClo(c *gin.Context)
	DeleteClo(c *gin.Context)
}

type cloController struct {
	cloRepo repo.CloRepository
}

func NewCloController(cloRepo repo.CloRepository) CloController {
	return &cloController{cloRepo}
}

func (cl *cloController) GetClo(c *gin.Context) {
	clo, err := cl.cloRepo.GetClo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, clo)
}

func (cl *cloController) GetCloById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clo, err := cl.cloRepo.GetCloById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, clo)
}

func (cl *cloController) CreateClo(c *gin.Context) {
	var body struct {
		PLOId     int     `json:"plo_id" binding:"required"`
		Nama      string  `json:"nama" binding:"required"`
		Deskripsi string  `json:"deskripsi" binding:"required"`
		Bobot     float64 `json:"bobot" binding:"required"`
		MkId      int     `json:"mk_id" binding:"required"`
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
	if body.Bobot == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bobot can't be empty"})
		return
	}
	if body.MkId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "mk_id can't be empty"})
		return
	}

	var clo model.CLO
	clo.PLOId = body.PLOId
	clo.Nama = body.Nama
	clo.Deskripsi = body.Deskripsi
	clo.Bobot = body.Bobot
	clo.MkId = body.MkId

	if err := cl.cloRepo.CreateClo(clo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "CLO added successfully"})
}

func (cl *cloController) UpdateClo(c *gin.Context) {
	var body struct {
		PLOId     int     `json:"plo_id" binding:"required"`
		Nama      string  `json:"nama" binding:"required"`
		Deskripsi string  `json:"deskripsi" binding:"required"`
		Bobot     float64 `json:"bobot" binding:"required"`
		MkId      int     `json:"mk_id" binding:"required"`
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
	if body.Bobot == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bobot can't be empty"})
		return
	}
	if body.MkId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "mk_id can't be empty"})
		return
	}

	var clo model.CLO
	clo.Nama = body.Nama
	clo.Deskripsi = body.Deskripsi
	clo.Bobot = body.Bobot
	clo.MkId = body.MkId
	clo.PLOId = body.PLOId
	clo.ID = id
	if err := cl.cloRepo.UpdateClo(clo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "CLO updated successfully"})
}

func (cl *cloController) DeleteClo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cl.cloRepo.DeleteClo(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "CLO deleted successfully"})
}

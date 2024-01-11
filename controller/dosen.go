package controller

import (
	repo "api-obe/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DosenController interface {
	GetDosen(c *gin.Context)
	GetMataKuliahObe(c *gin.Context)
}

type dosenController struct {
	dosenRepo      repo.DosenRepository
	mataKuliahRepo repo.MataKuliahRepository
}

func NewDosenController(dosenRepo repo.DosenRepository, mataKuliahRepo repo.MataKuliahRepository) DosenController {
	return &dosenController{dosenRepo, mataKuliahRepo}
}

func (d *dosenController) GetDosen(c *gin.Context) {
	dosen, err := d.dosenRepo.GetDosen()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dosen)
}

func (d *dosenController) GetMataKuliahObe(c *gin.Context) {
	id, err := strconv.Atoi("id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	obeId, err := strconv.Atoi("obeId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mataKuliah, err := d.mataKuliahRepo.GetMataKuliahByDosenObeId(obeId, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mataKuliah)
}

package controller

import (
	"api-obe/model"
	repo "api-obe/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DosenController interface {
	GetDosen(c *gin.Context)
	GetMataKuliah(c *gin.Context)
}

type dosenController struct {
	dosenRepo       repo.DosenRepository
	mataKuliahRepo  repo.MataKuliahRepository
	perancanganRepo repo.PerancanganObeRepository
}

func NewDosenController(dosenRepo repo.DosenRepository, mataKuliahRepo repo.MataKuliahRepository, perancanganRepo repo.PerancanganObeRepository) DosenController {
	return &dosenController{dosenRepo, mataKuliahRepo, perancanganRepo}
}

func (d *dosenController) GetDosen(c *gin.Context) {
	dosen, err := d.dosenRepo.GetDosen()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dosen)
}

func (d *dosenController) GetMataKuliah(c *gin.Context) {
	user, _ := c.Get("user")

	userExist := user.(model.User)

	dosen, err := d.dosenRepo.GetDosenByUserId(userExist.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	perancanganObe, err := d.perancanganRepo.GetActivePerancanganObe()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	mataKuliah, err := d.mataKuliahRepo.GetMataKuliahByDosenObeId(perancanganObe.ID, dosen.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mataKuliah)
}

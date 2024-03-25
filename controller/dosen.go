package controller

import (
	"api-obe/model"
	repo "api-obe/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DosenController interface {
	GetDosen(c *gin.Context)
	GetMataKuliahByTahun(c *gin.Context)
	SearchMataKuliahByTahun(c *gin.Context)
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

func (d *dosenController) GetMataKuliahByTahun(c *gin.Context) {
	tahunId, err := strconv.Atoi(c.Param("tahunId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
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

	mataKuliah, err := d.mataKuliahRepo.GetMataKuliahByDosenObeIdAndTahunId(perancanganObe.ID, dosen.ID, tahunId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mataKuliah)
}

func (d *dosenController) SearchMataKuliahByTahun(c *gin.Context) {
	tahunId, err := strconv.Atoi(c.Param("tahunId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	keyword := c.Query("keyword")
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

	mataKuliah, err := d.mataKuliahRepo.GetMataKuliahByDosenObeIdAndTahunIdAndKeyword(perancanganObe.ID, dosen.ID, tahunId, keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mataKuliah)
}

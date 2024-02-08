package controller

import (
	"api-obe/model"
	repo "api-obe/repository"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type PlottingDosenMkController interface {
	GetPlottingDosenMk(c *gin.Context)
	SearchPlottingDosenMk(c *gin.Context)
	CreatePlottingDosenMk(c *gin.Context)
	DeletePlottingDosenMk(c *gin.Context)
}

type plottingDosenMkController struct {
	plottingDosenMkRepo repo.PlottingDosenMkRepository
	perancanganRepo     repo.PerancanganObeRepository
	dosenRepo           repo.DosenRepository
	kelasRepo           repo.KelasRepository
	mataKuliahRepo      repo.MataKuliahRepository
}

func NewPlottingDosenMkController(plottingDosenMkRepo repo.PlottingDosenMkRepository, perancanganRepo repo.PerancanganObeRepository, dosenRepo repo.DosenRepository, kelasRepo repo.KelasRepository, mataKuliahRepo repo.MataKuliahRepository) PlottingDosenMkController {
	return &plottingDosenMkController{
		plottingDosenMkRepo,
		perancanganRepo,
		dosenRepo,
		kelasRepo,
		mataKuliahRepo,
	}
}

func (m *plottingDosenMkController) GetPlottingDosenMk(c *gin.Context) {
	obe, err := m.perancanganRepo.GetActivePerancanganObe()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	plottingDosenMk, err := m.plottingDosenMkRepo.GetPlottingDosenMkByObeId(obe.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var plotDatas []model.PlotData
	for _, v := range plottingDosenMk {
		var plot model.PlotData
		plot.ID = v.ID
		dosen, _ := m.dosenRepo.GetDosenById(v.DosenId)
		kelas, _ := m.kelasRepo.GetKelasById(v.KelasId)
		mk, _ := m.mataKuliahRepo.GetMataKuliahById(v.MKId)
		plot.Dosen = dosen.KodeDosen
		plot.Kelas = kelas.KodeKelas
		plot.MataKuliah = mk.KodeMk
		plotDatas = append(plotDatas, plot)
	}

	c.JSON(http.StatusOK, plotDatas)
}

func (m *plottingDosenMkController) SearchPlottingDosenMk(c *gin.Context) {
	keyword := strings.ToUpper(c.Query("keyword"))
	obe, err := m.perancanganRepo.GetActivePerancanganObe()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	plottingDosenMk, err := m.plottingDosenMkRepo.GetPlottingDosenMkByObeId(obe.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	plotDatas := []model.PlotData{}
	for _, v := range plottingDosenMk {
		var plot model.PlotData
		dosen, _ := m.dosenRepo.GetDosenById(v.DosenId)
		kelas, _ := m.kelasRepo.GetKelasById(v.KelasId)
		mk, _ := m.mataKuliahRepo.GetMataKuliahById(v.MKId)
		if strings.Contains(mk.KodeMk, keyword) || strings.Contains(dosen.KodeDosen, keyword) {
			plot.ID = v.ID
			plot.Dosen = dosen.KodeDosen
			plot.Kelas = kelas.KodeKelas
			plot.MataKuliah = mk.KodeMk
			plotDatas = append(plotDatas, plot)
		}
	}

	c.JSON(http.StatusOK, plotDatas)
}

func (m *plottingDosenMkController) CreatePlottingDosenMk(c *gin.Context) {
	var body struct {
		MKId    int `json:"mk_id" binding:"required"`
		DosenId int `json:"dosen_id" binding:"required"`
		KelasId int `json:"kelas_id" binding:"required"`
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if body.MKId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "mk_id can't be empty"})
		return
	}
	if body.DosenId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dosen_id can't be empty"})
		return
	}
	var plottingDosenMk model.PlottingDosenMk
	plottingDosenMk.MKId = body.MKId
	plottingDosenMk.DosenId = body.DosenId
	plottingDosenMk.KelasId = body.KelasId
	err := m.plottingDosenMkRepo.CreatePlottingDosenMk(plottingDosenMk)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Plotting dosen added successfully"})
}

func (m *plottingDosenMkController) DeletePlottingDosenMk(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = m.plottingDosenMkRepo.DeletePlottingDosenMk(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Plotting dosen deleted successfully"})
}

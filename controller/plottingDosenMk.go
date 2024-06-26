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
	GetPlottingDosenMkByTahun(c *gin.Context)
	SearchPlottingDosenMkByTahun(c *gin.Context)
	GetKelasDosenByMk(c *gin.Context)
	GetKelasMkByMk(c *gin.Context)
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

func (m *plottingDosenMkController) GetPlottingDosenMkByTahun(c *gin.Context) {
	tahunId, err := strconv.Atoi(c.Param("tahunId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	obe, err := m.perancanganRepo.GetActivePerancanganObe()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	plottingDosenMk, err := m.plottingDosenMkRepo.GetPlottingDosenMkByObeIdAndTahunId(obe.ID, tahunId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	plotDatas := []model.PlotData{}
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

func (m *plottingDosenMkController) SearchPlottingDosenMkByTahun(c *gin.Context) {
	keyword := strings.ToUpper(c.Query("keyword"))
	obe, err := m.perancanganRepo.GetActivePerancanganObe()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tahunId, err := strconv.Atoi(c.Param("tahunId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plottingDosenMk, err := m.plottingDosenMkRepo.GetPlottingDosenMkByObeIdAndTahunId(obe.ID, tahunId)
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

func (m *plottingDosenMkController) GetKelasDosenByMk(c *gin.Context) {
	user, _ := c.Get("user")

	userExist := user.(model.User)

	dosen, err := m.dosenRepo.GetDosenByUserId(userExist.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	mkId, err := strconv.Atoi(c.Param("mkId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plotting, err := m.plottingDosenMkRepo.GetPlottingDosenByMkIdAndDosenId(mkId, dosen.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	kelasList := []model.Kelas{}
	for _, v := range plotting {
		kelas, err := m.kelasRepo.GetKelasById(v.KelasId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		kelasList = append(kelasList, kelas)
	}

	c.JSON(http.StatusOK, kelasList)
}

func (m *plottingDosenMkController) GetKelasMkByMk(c *gin.Context) {
	mkId, err := strconv.Atoi(c.Param("mkId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plotting, err := m.plottingDosenMkRepo.GetPlottingDosenByMkId(mkId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	kelasList := []model.Kelas{}
	for _, v := range plotting {
		kelas, err := m.kelasRepo.GetKelasById(v.KelasId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		kelasList = append(kelasList, kelas)
	}

	c.JSON(http.StatusOK, kelasList)
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

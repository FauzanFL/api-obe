package controller

import (
	"api-obe/model"
	repo "api-obe/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlottingDosenMkController interface {
	GetPlottingDosenMk(c *gin.Context)
	CreatePlottingDosenMk(c *gin.Context)
	DeletePlottingDosenMk(c *gin.Context)
}

type plottingDosenMkController struct {
	plottingDosenMkRepo repo.PlottingDosenMkRepository
}

func NewPlottingDosenMkController(plottingDosenMkRepo repo.PlottingDosenMkRepository) PlottingDosenMkController {
	return &plottingDosenMkController{
		plottingDosenMkRepo,
	}
}

func (m *plottingDosenMkController) GetPlottingDosenMk(c *gin.Context) {
	plottingDosenMk, err := m.plottingDosenMkRepo.GetPlottingDosenMk()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, plottingDosenMk)
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

package controller

import (
	repo "api-obe/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlottingDosenMkController interface {
	GetPlottingDosenMk(c *gin.Context)
}

type plottingDosenMkController struct {
	plottingDosenMkRepo repo.PlottingDosenMkRepository
}

func NewPlottingDosenMkController(plottingDosenMkRepo repo.PlottingDosenMkRepository) PlottingDosenMkController {
	return &plottingDosenMkController{
		plottingDosenMkRepo,
	}
}

func (

	m *plottingDosenMkController) GetPlottingDosenMk(c *gin.Context) {
	plottingDosenMk, err := m.plottingDosenMkRepo.GetPlottingDosenMk()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, plottingDosenMk)
}

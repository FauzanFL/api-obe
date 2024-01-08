package controller

import (
	repo "api-obe/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DosenController interface {
	GetDosen(c *gin.Context)
}

type dosenController struct {
	dosenRepo repo.DosenRepository
}

func NewDosenController(dosenRepo repo.DosenRepository) DosenController {
	return &dosenController{dosenRepo}
}

func (d *dosenController) GetDosen(c *gin.Context) {
	dosen, err := d.dosenRepo.GetDosen()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dosen)
}

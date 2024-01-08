package controller

import (
	repo "api-obe/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CloController interface {
	GetClo(c *gin.Context)
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

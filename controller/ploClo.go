package controller

import (
	"api-obe/model"
	repo "api-obe/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PloCloController interface {
	GetPloClo(c *gin.Context)
	CreatePloClo(c *gin.Context)
	DeletePloClo(c *gin.Context)
}

type ploCloController struct {
	ploCloRepository repo.PloCloRepository
}

func NewPloCloController(ploCloRepository repo.PloCloRepository) PloCloController {
	return &ploCloController{ploCloRepository}
}

func (p *ploCloController) GetPloClo(c *gin.Context) {
	result, err := p.ploCloRepository.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (p *ploCloController) CreatePloClo(c *gin.Context) {
	ploClo := model.PLO_CLO{}
	if err := c.ShouldBindJSON(&ploClo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := p.ploCloRepository.Create(ploClo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ploClo)
}

func (p *ploCloController) DeletePloClo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = p.ploCloRepository.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

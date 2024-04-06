package controller

import (
	"api-obe/model"
	repo "api-obe/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type KurikulumController interface {
	GetKurikulum(c *gin.Context)
	CreateKurikulum(c *gin.Context)
	DeleteKurikulum(c *gin.Context)
}

type kurikulumController struct {
	kurikulumRepo repo.KurikulumRepository
}

func NewKurikulumController(kurikulumRepo repo.KurikulumRepository) KurikulumController {
	return &kurikulumController{
		kurikulumRepo,
	}
}

func (k *kurikulumController) GetKurikulum(c *gin.Context) {
	kurikulum, err := k.kurikulumRepo.GetKurikulum()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, kurikulum)
}

func (k *kurikulumController) CreateKurikulum(c *gin.Context) {
	var body struct {
		Nama string `json:"nama" binding:"required"`
	}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	kurikulum := model.Kurikulum{
		Nama: body.Nama,
	}
	err = k.kurikulumRepo.CreateKurikulum(kurikulum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Kurikulum created successfully"})
}

func (k *kurikulumController) DeleteKurikulum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = k.kurikulumRepo.DeleteKurikulum(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Kurikulum deleted successfully"})
}

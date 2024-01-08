package controller

import (
	repo "api-obe/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type KurikulumController interface {
	GetKurikulum(c *gin.Context)
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

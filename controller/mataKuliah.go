package controller

import (
	repo "api-obe/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MataKuliahController interface {
	GetMataKuliah(c *gin.Context)
}

type mataKuliahController struct {
	mataKuliahRepo repo.MataKuliahRepository
}

func NewMataKuliahController(mataKuliahRepo repo.MataKuliahRepository) MataKuliahController {
	return &mataKuliahController{
		mataKuliahRepo,
	}
}

func (m *mataKuliahController) GetMataKuliah(c *gin.Context) {
	mataKuliah, err := m.mataKuliahRepo.GetMataKuliah()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mataKuliah)
}

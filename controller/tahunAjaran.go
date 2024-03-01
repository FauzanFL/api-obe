package controller

import (
	"api-obe/model"
	repo "api-obe/repository"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type TahunAjaranController interface {
	GetTahunAjaran(c *gin.Context)
	GetTahunAjaranNow(c *gin.Context)
}

type tahunAjaranController struct {
	tahunAjaranRepo repo.TahunAjaranRepository
}

func NewTahunAjaranController(tahunAjaranRepo repo.TahunAjaranRepository) TahunAjaranController {
	return &tahunAjaranController{
		tahunAjaranRepo,
	}
}

func (t *tahunAjaranController) GetTahunAjaran(c *gin.Context) {
	tahunAjaran, err := t.tahunAjaranRepo.GetTahunAjaran()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tahunAjaran)
}

func (t *tahunAjaranController) GetTahunAjaranNow(c *gin.Context) {
	currentTime := time.Now()
	currentMonth := currentTime.Month()
	currentYear := currentTime.Year()

	tahunAjaran, err := t.tahunAjaranRepo.GetTahunAjaranByMonth(int(currentMonth))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tahunAjarNow := model.TahunAjaran{}
	for _, v := range tahunAjaran {
		tahunArr := strings.Split(v.Tahun, "/")
		tahunStart, err := strconv.Atoi(tahunArr[0])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tahunEnd, err := strconv.Atoi(tahunArr[1])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// masih belum betul
		if v.Semester == "Ganjil" && tahunStart == currentYear {
			tahunAjarNow = v
			break
		} else if v.Semester == "Genap" && tahunEnd == currentYear {
			tahunAjarNow = v
			break
		}
	}

	c.JSON(http.StatusOK, tahunAjarNow)
}

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
	SearchTahunAjaran(c *gin.Context)
	CreateTahunAjaran(c *gin.Context)
	UpdateTahunAjaran(c *gin.Context)
	DeleteTahunAjaran(c *gin.Context)
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

func (t *tahunAjaranController) SearchTahunAjaran(c *gin.Context) {
	keyword := c.Query("keyword")

	tahunAjar, err := t.tahunAjaranRepo.GetTahunAjaranByKeyword(keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tahunAjar)
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

func (t *tahunAjaranController) CreateTahunAjaran(c *gin.Context) {
	var body struct {
		Tahun        string `json:"tahun" binding:"required"`
		Semester     string `json:"semester" binding:"required"`
		BulanMulai   int    `json:"bulan_mulai" binding:"required"`
		BulanSelesai int    `json:"bulan_selesai" binding:"required"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if body.Tahun == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tahun can't be empty"})
		return
	}
	if body.Semester == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "semester can't be empty"})
		return
	}
	if body.BulanMulai == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bulan mulai can't be empty"})
		return
	}
	if body.BulanSelesai == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bulan selesai can't be empty"})
		return
	}

	var tahunAjaran model.TahunAjaran
	tahunAjaran.Tahun = body.Tahun
	tahunAjaran.Semester = body.Semester
	tahunAjaran.BulanMulai = body.BulanMulai
	tahunAjaran.BulanSelesai = body.BulanSelesai
	if err := t.tahunAjaranRepo.CreateTahunAjaran(tahunAjaran); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Tahun ajaran added successfully"})
}

func (t *tahunAjaranController) UpdateTahunAjaran(c *gin.Context) {
	var body struct {
		Tahun        string `json:"tahun" binding:"required"`
		Semester     string `json:"semester" binding:"required"`
		BulanMulai   int    `json:"bulan_mulai" binding:"required"`
		BulanSelesai int    `json:"bulan_selesai" binding:"required"`
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if body.Tahun == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tahun can't be empty"})
		return
	}
	if body.Semester == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "semester can't be empty"})
		return
	}
	if body.BulanMulai == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bulan mulai can't be empty"})
		return
	}
	if body.BulanSelesai == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bulan selesai can't be empty"})
		return
	}

	var tahunAjaran model.TahunAjaran
	tahunAjaran.ID = id
	tahunAjaran.Tahun = body.Tahun
	tahunAjaran.Semester = body.Semester
	tahunAjaran.BulanMulai = body.BulanMulai
	tahunAjaran.BulanSelesai = body.BulanSelesai
	if err := t.tahunAjaranRepo.UpdateTahunAjaran(tahunAjaran); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Tahun ajaran updated successfully"})
}

func (t *tahunAjaranController) DeleteTahunAjaran(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := t.tahunAjaranRepo.DeleteTahunAjaran(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Tahun ajaran deleted successfully"})
}

package controller

import (
	"api-obe/db"
	"api-obe/model"
	repo "api-obe/repository"
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type PenilaianController interface {
	GetPenilaian(c *gin.Context)
	GetPenilaianById(c *gin.Context)
	CreatePenilaian(c *gin.Context)
	UpdatePenilaian(c *gin.Context)
	DeletePenilaian(c *gin.Context)
	UploadEvidence(c *gin.Context)
}

type penilaianController struct {
	penilaianRepo repo.PenilaianRepository
}

func NewPenilaianController(penilaianRepo repo.PenilaianRepository) PenilaianController {
	return &penilaianController{
		penilaianRepo,
	}
}

func (p *penilaianController) GetPenilaian(c *gin.Context) {
	penilaian, err := p.penilaianRepo.GetPenilaian()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, penilaian)
}

func (p *penilaianController) GetPenilaianById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	penilaian, err := p.penilaianRepo.GetPenilaianById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, penilaian)
}

func (p *penilaianController) CreatePenilaian(c *gin.Context) {
	var body struct {
		Nilai         float64 `json:"nilai" binding:"required"`
		AssessmentId  int     `json:"assessment_id" binding:"required"`
		MhsId         int     `json:"mhs_id" binding:"required"`
		DosenId       int     `json:"dosen_id" binding:"required"`
		TahunAjaranId int     `json:"tahun_ajaran_id" binding:"required"`
	}
	if body.Nilai == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nilai can't be empty"})
		return
	}
	if body.AssessmentId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "assessment_id can't be empty"})
		return
	}
	if body.MhsId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "mhs_id can't be empty"})
		return
	}
	if body.DosenId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dosen_id can't be empty"})
		return
	}
	if body.TahunAjaranId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tahun_ajaran_id can't be empty"})
		return
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var penilaian model.Penilaian
	penilaian.Nilai = body.Nilai
	penilaian.AssessmentId = body.AssessmentId
	penilaian.MhsId = body.MhsId
	penilaian.DosenId = body.DosenId
	penilaian.TahunAjaranId = body.TahunAjaranId

	if err := p.penilaianRepo.CreatePenilaian(penilaian); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Penilaian added successfully"})
}

func (p *penilaianController) UpdatePenilaian(c *gin.Context) {
	var body struct {
		Nilai         float64 `json:"nilai" binding:"required"`
		AssessmentId  int     `json:"assessment_id" binding:"required"`
		MhsId         int     `json:"mhs_id" binding:"required"`
		DosenId       int     `json:"dosen_id" binding:"required"`
		TahunAjaranId int     `json:"tahun_ajaran_id" binding:"required"`
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
	if body.Nilai == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nilai can't be empty"})
		return
	}
	if body.AssessmentId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "assessment_id can't be empty"})
		return
	}
	if body.MhsId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "mhs_id can't be empty"})
		return
	}
	if body.DosenId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dosen_id can't be empty"})
		return
	}
	if body.TahunAjaranId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tahun_ajaran_id can't be empty"})
		return
	}

	var penilaian model.Penilaian
	penilaian.Nilai = body.Nilai
	penilaian.AssessmentId = body.AssessmentId
	penilaian.MhsId = body.MhsId
	penilaian.DosenId = body.DosenId
	penilaian.TahunAjaranId = body.TahunAjaranId
	penilaian.ID = id
	if err := c.Bind(&penilaian); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := p.penilaianRepo.UpdatePenilaian(penilaian); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Penilaian updated successfully"})
}

func (p *penilaianController) DeletePenilaian(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = p.penilaianRepo.DeletePenilaian(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Penilaian deleted successfully"})
}

func (p *penilaianController) UploadEvidence(c *gin.Context) {
	ctx := context.Background()

	client := db.ServiceAccount("client_secret.json")

	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer f.Close()

	fmt.Println(file.Filename)

	fileMetadata := &drive.File{
		MimeType: "application/octet-stream",
		Name:     file.Filename,
		Parents:  []string{"1nFGLiS9VIdTN5KIZn-Cfxg5s6rEymsET"},
	}

	_, err = srv.Files.Create(fileMetadata).Media(f).Do()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}

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
	"github.com/maxbeatty/golang-book/chapter11/math"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type PenilaianController interface {
	GetPenilaian(c *gin.Context)
	GetDataPenilaian(c *gin.Context)
	GetDataPenilaianCLOPLOByMk(c *gin.Context)
	GetDataPenilaianPLO(c *gin.Context)
	GetPenilaianById(c *gin.Context)
	GetPenilaianByKelas(c *gin.Context)
	CreatePenilaian(c *gin.Context)
	UpdatePenilaian(c *gin.Context)
	DeletePenilaian(c *gin.Context)
	UploadEvidence(c *gin.Context)
}

type penilaianController struct {
	penilaianRepo   repo.PenilaianRepository
	cloRepo         repo.CloRepository
	assessmentRepo  repo.LembarAssessmentRepository
	mahasiswaRepo   repo.MahasiswaRepository
	dosenRepo       repo.DosenRepository
	ploRepo         repo.PloRepository
	perancanganRepo repo.PerancanganObeRepository
}

func NewPenilaianController(penilaianRepo repo.PenilaianRepository, cloRepo repo.CloRepository, assessmentRepo repo.LembarAssessmentRepository, mahasiswaRepo repo.MahasiswaRepository, dosenRepo repo.DosenRepository, ploRepo repo.PloRepository, perancanganRepo repo.PerancanganObeRepository) PenilaianController {
	return &penilaianController{
		penilaianRepo,
		cloRepo,
		assessmentRepo,
		mahasiswaRepo,
		dosenRepo,
		ploRepo,
		perancanganRepo,
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

func (p *penilaianController) GetDataPenilaian(c *gin.Context) {
	kelasId, err := strconv.Atoi(c.Param("kelasId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tahunId, err := strconv.Atoi(c.Param("tahunId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mkId, err := strconv.Atoi(c.Param("mkId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clo, err := p.cloRepo.GetCLOByMkId(mkId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cloWithAssessment := []model.CLOWithAssessment{}
	assessmentsMhs := []model.LembarAssessmentWithJenis{}
	for _, v := range clo {
		assessments := []model.LembarAssessmentWithJenis{}
		result, err := p.assessmentRepo.GetLembarAssessmentByCloId(v.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		assessments = append(assessments, result...)
		assessmentsMhs = append(assessmentsMhs, assessments...)
		cloWithAssessment = append(cloWithAssessment, model.CLOWithAssessment{
			ID:          v.ID,
			PLOId:       v.PLOId,
			Nama:        v.Nama,
			Deskripsi:   v.Deskripsi,
			Bobot:       v.Bobot,
			MkId:        v.MkId,
			Assessments: assessments,
		})
	}

	mahasiswa, err := p.mahasiswaRepo.GetMahasiswaByKelasMataKuliah(mkId, kelasId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	mahasiswaWithNilai := []model.MahasiswaWithPenilaian{}
	for _, v := range mahasiswa {
		nilaiAssessments := []model.Penilaian{}
		for _, val := range assessmentsMhs {
			nilai, _ := p.penilaianRepo.GetPenilaianByMhsIdAndAssessmentId(v.ID, val.ID, tahunId)
			if nilai.ID != 0 {
				nilaiAssessments = append(nilaiAssessments, nilai)
			}
		}
		mahasiswaWithNilai = append(mahasiswaWithNilai, model.MahasiswaWithPenilaian{
			ID:        v.ID,
			NIM:       v.NIM,
			Nama:      v.Nama,
			KelasId:   v.KelasId,
			Penilaian: nilaiAssessments,
		})
	}

	dataPenilaian := model.PenilaianData{
		CLOAsessment:   cloWithAssessment,
		MahasiswaNilai: mahasiswaWithNilai,
	}
	c.JSON(http.StatusOK, dataPenilaian)
}

func (p *penilaianController) GetDataPenilaianCLOPLOByMk(c *gin.Context) {
	mkId, err := strconv.Atoi(c.Param("mkId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tahunId, err := strconv.Atoi(c.Param("tahunId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clo, err := p.cloRepo.GetCLOByMkId(mkId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	listCloWithNilai := make([]model.CLOWithNilai, len(clo))
	listPloWithNilai := []model.PLOWithNilai{}
	listPlo := map[int]model.PLO{}

	for i, v := range clo {
		assessments, _ := p.assessmentRepo.GetLembarAssessmentByCloId(v.ID)
		var nilaiFinal = map[int][]float64{}
		if len(assessments) > 0 {
			for _, val := range assessments {
				penilaian, _ := p.penilaianRepo.GetPenilaianByAssessmentId(val.ID, tahunId)
				if len(penilaian) > 0 {
					for _, value := range penilaian {
						nilaiFinal[value.MhsId] = append(nilaiFinal[value.MhsId], value.Nilai)
					}
				} else {
					nilaiFinal[0] = []float64{0}
				}
			}
		} else {
			nilaiFinal = map[int][]float64{
				0: {0},
			}
		}

		plo, _ := p.ploRepo.GetPloById(v.PLOId)

		var total float64 = 0.0
		for _, numbers := range nilaiFinal {
			total += math.Average(numbers)
		}

		avg := total / float64(len(nilaiFinal))
		var formattedAvg = formatAvg(avg)
		cloWithNilai := model.CLOWithNilai{
			ID:    v.ID,
			Nama:  v.Nama,
			PLOId: v.PLOId,
			Nilai: formattedAvg,
		}
		listCloWithNilai[i] = cloWithNilai

		listPlo[plo.ID] = plo
	}

	for _, v := range listPlo {
		var nilaiFinal = []float64{}
		for _, val := range listCloWithNilai {
			if v.ID == val.PLOId {
				nilaiFinal = append(nilaiFinal, val.Nilai)
			}
		}
		avg := math.Average(nilaiFinal)
		var formattedAvg = formatAvg(avg)
		ploWithNilai := model.PLOWithNilai{
			ID:    v.ID,
			Nama:  v.Nama,
			Nilai: formattedAvg,
		}
		listPloWithNilai = append(listPloWithNilai, ploWithNilai)
	}

	var response struct {
		Clo []model.CLOWithNilai `json:"clo"`
		Plo []model.PLOWithNilai `json:"plo"`
	}

	response.Clo = listCloWithNilai
	response.Plo = listPloWithNilai

	c.JSON(http.StatusOK, response)
}

func (p *penilaianController) GetDataPenilaianPLO(c *gin.Context) {
	tahunId, err := strconv.Atoi(c.Param("tahunId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	obe, err := p.perancanganRepo.GetActivePerancanganObe()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plo, err := p.ploRepo.GetPloByObeId(obe.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	listPloWithNilai := []model.PLOWithNilai{}
	for _, v := range plo {
		clo, _ := p.cloRepo.GetCLOByPLOId(v.ID)
		var nilaiFinal float64
		if len(clo) > 0 {
			var nilaiClo []float64
			for _, c := range clo {
				assessments, _ := p.assessmentRepo.GetLembarAssessmentByCloId(c.ID)
				for _, val := range assessments {
					penilaian, _ := p.penilaianRepo.GetPenilaianByAssessmentId(val.ID, tahunId)
					nilai := make(map[int][]float64)
					if len(penilaian) > 0 {
						for _, value := range penilaian {
							nilai[value.MhsId] = append(nilai[value.MhsId], value.Nilai)
						}
					} else {
						nilai[0] = []float64{0}
					}
					var total float64 = 0.0
					for _, numbers := range nilai {
						total += math.Average(numbers)
					}
					avg := total / float64(len(nilai))
					nilaiClo = append(nilaiClo, avg)
				}

			}

			fmt.Println(nilaiClo)

			nilaiFinal = math.Average(nilaiClo)
		} else {
			nilaiFinal = 0
		}

		var formattedAvg = formatAvg(nilaiFinal)

		ploWithNilai := model.PLOWithNilai{
			ID:    v.ID,
			Nama:  v.Nama,
			Nilai: formattedAvg,
		}
		listPloWithNilai = append(listPloWithNilai, ploWithNilai)
	}

	c.JSON(http.StatusOK, listPloWithNilai)
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

func (p *penilaianController) GetPenilaianByKelas(c *gin.Context) {
	kelasId, err := strconv.Atoi(c.Param("kelasId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tahunId, err := strconv.Atoi(c.Param("tahunId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	penilaian, err := p.penilaianRepo.GetPenilaianByKelas(kelasId, tahunId)
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
		TahunAjaranId int     `json:"tahun_ajaran_id" binding:"required"`
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
	if body.TahunAjaranId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tahun_ajaran_id can't be empty"})
		return
	}

	userGet, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not exist"})
		return
	}

	user := userGet.(model.User)

	dosen, err := p.dosenRepo.GetDosenByUserId(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var penilaian model.Penilaian
	penilaian.Nilai = body.Nilai
	penilaian.AssessmentId = body.AssessmentId
	penilaian.MhsId = body.MhsId
	penilaian.DosenId = dosen.ID
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
	if body.TahunAjaranId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tahun_ajaran_id can't be empty"})
		return
	}

	userGet, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not exist"})
		return
	}

	user := userGet.(model.User)

	dosen, err := p.dosenRepo.GetDosenByUserId(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var penilaian model.Penilaian
	penilaian.Nilai = body.Nilai
	penilaian.AssessmentId = body.AssessmentId
	penilaian.MhsId = body.MhsId
	penilaian.DosenId = dosen.ID
	penilaian.TahunAjaranId = body.TahunAjaranId
	penilaian.ID = id

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

func formatAvg(avg float64) float64 {
	if avg == float64(int64(avg)) {
		return avg
	} else {
		parsedNumber, err := strconv.ParseFloat(fmt.Sprintf("%.2f", avg), 64)
		if err != nil {
			return avg
		}
		return parsedNumber
	}
}

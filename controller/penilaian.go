package controller

import (
	"api-obe/db"
	"api-obe/model"
	repo "api-obe/repository"
	"context"
	"encoding/json"
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
	CreatePenilaian(c *gin.Context)
	UpdatePenilaian(c *gin.Context)
	DeletePenilaian(c *gin.Context)
	UploadEvidence(c *gin.Context)
}

type penilaianController struct {
	penilaianRepo   repo.PenilaianRepository
	cloRepo         repo.CloRepository
	assessmentRepo  repo.LembarAssessmentRepository
	dosenRepo       repo.DosenRepository
	ploRepo         repo.PloRepository
	perancanganRepo repo.PerancanganObeRepository
	plottingRepo    repo.PlottingDosenMkRepository
}

func NewPenilaianController(penilaianRepo repo.PenilaianRepository, cloRepo repo.CloRepository, assessmentRepo repo.LembarAssessmentRepository, dosenRepo repo.DosenRepository, ploRepo repo.PloRepository, perancanganRepo repo.PerancanganObeRepository, plottingRepo repo.PlottingDosenMkRepository) PenilaianController {
	return &penilaianController{
		penilaianRepo,
		cloRepo,
		assessmentRepo,
		dosenRepo,
		ploRepo,
		perancanganRepo,
		plottingRepo,
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
	var listCLoWithAssessment []model.CLOWithAssessment

	for _, v := range clo {
		assessments, err := p.assessmentRepo.GetLembarAssessmentByCloId(v.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		cloWithAssessment := model.CLOWithAssessment{
			ID:          v.ID,
			PLOId:       v.PLOId,
			Nama:        v.Nama,
			Deskripsi:   v.Deskripsi,
			Bobot:       v.Bobot,
			MkId:        v.MkId,
			Assessments: assessments,
		}
		listCLoWithAssessment = append(listCLoWithAssessment, cloWithAssessment)
	}

	penilaianRes := model.PenilaianResp{}
	penilaian, _ := p.penilaianRepo.GetPenilaianByKelasIdAndMkId(kelasId, mkId)

	nilaiMahasiswa := []model.NilaiMahasiswa{}
	err = json.Unmarshal([]byte(penilaian.Nilai), &nilaiMahasiswa)
	if err != nil {
		fmt.Println("Error:", err)
		penilaianRes.Nilai = []model.NilaiMahasiswa{}
	}
	penilaianRes.ID = penilaian.ID
	penilaianRes.Status = penilaian.Status
	penilaianRes.Nilai = nilaiMahasiswa
	penilaianRes.MkId = penilaian.MkId
	penilaianRes.KelasId = penilaian.KelasId

	dataPenilaian := model.PenilaianData{
		CLOAsessment: listCLoWithAssessment,
		Penilaian:    penilaianRes,
	}
	c.JSON(http.StatusOK, dataPenilaian)
}

func (p *penilaianController) GetDataPenilaianCLOPLOByMk(c *gin.Context) {
	mkId, err := strconv.Atoi(c.Param("mkId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	plotting, err := p.plottingRepo.GetPlottingDosenByMkIdAndDosenId(mkId, dosen.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	clo, err := p.cloRepo.GetCLOByMkId(mkId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cloWithNilai := []model.CLOWithNilai{}
	ploWithNilai := []model.PLOWithNilai{}

	listPLO := map[int]model.PLO{}

	for _, v := range clo {
		plo, _ := p.ploRepo.GetPloById(v.PLOId)
		listPLO[plo.ID] = plo

		nilaiArr := []float64{}

		assessments, _ := p.assessmentRepo.GetLembarAssessmentByCloId(v.ID)
		for _, plot := range plotting {
			penilaian, _ := p.penilaianRepo.GetPenilaianByKelasIdAndMkId(plot.KelasId, mkId)
			if penilaian.ID != 0 {
				nilaiMahasiswa := []model.NilaiMahasiswa{}
				err = json.Unmarshal([]byte(penilaian.Nilai), &nilaiMahasiswa)
				if err != nil {
					fmt.Println("Error:", err)
				}

				mhsNilaiArr := []float64{}
				for _, mhs := range nilaiMahasiswa {
					mhsTotal := 0.0
					for _, nilai := range mhs.NilaiAssessment {
						for _, assessment := range assessments {
							if nilai.AssessmentId == assessment.ID {
								mhsTotal += nilai.Nilai
							}
						}

					}
					nt := mhsTotal / float64(len(assessments))
					mhsNilaiArr = append(mhsNilaiArr, nt)
				}

				nilaiArr = append(nilaiArr, math.Average(mhsNilaiArr))
			}

		}
		avgCloNilai := math.Average(nilaiArr)
		cloWithNilai = append(cloWithNilai, model.CLOWithNilai{
			ID:    v.ID,
			Nama:  v.Nama,
			PLOId: v.PLOId,
			Nilai: formatAvg(avgCloNilai),
		})
	}
	for _, v := range listPLO {
		total := 0.0
		count := 0
		for _, clo := range cloWithNilai {
			if clo.PLOId == v.ID {
				total += clo.Nilai
				count++
			}
		}

		avgTotal := total / float64(count)
		ploWithNilai = append(ploWithNilai, model.PLOWithNilai{
			ID:    v.ID,
			Nama:  v.Nama,
			Nilai: formatAvg(avgTotal),
		})
	}

	result := model.CloPloWithNilai{}
	result.Clo = cloWithNilai
	result.Plo = ploWithNilai

	c.JSON(http.StatusOK, result)
}

func (p *penilaianController) GetDataPenilaianPLO(c *gin.Context) {
	perancangan, err := p.perancanganRepo.GetActivePerancanganObe()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	plo, err := p.ploRepo.GetPloByObeId(perancangan.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	penilaian, err := p.penilaianRepo.GetPenilaian()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ploWithNilai := []model.PLOWithNilai{}

	for _, v := range plo {
		clo, _ := p.cloRepo.GetCLOByPLOId(v.ID)
		cloNilai := []float64{}
		if len(clo) > 0 {
			for _, c := range clo {
				nilaiArr := []float64{}
				assessments, _ := p.assessmentRepo.GetLembarAssessmentByCloId(c.ID)
				for _, p := range penilaian {
					nilaiMahasiswa := []model.NilaiMahasiswa{}
					err = json.Unmarshal([]byte(p.Nilai), &nilaiMahasiswa)
					if err != nil {
						fmt.Println("Error:", err)
					}

					mhsNilaiArr := []float64{}
					for _, mhs := range nilaiMahasiswa {
						mhsTotal := 0.0
						for _, nilai := range mhs.NilaiAssessment {
							for _, assessment := range assessments {
								if nilai.AssessmentId == assessment.ID {
									mhsTotal += nilai.Nilai
								}
							}
						}
						nt := mhsTotal / float64(len(assessments))
						mhsNilaiArr = append(mhsNilaiArr, nt)
					}

					nt := math.Average(mhsNilaiArr)
					nilaiArr = append(nilaiArr, nt)
				}
				nt := math.Average(nilaiArr)
				cloNilai = append(cloNilai, nt)
			}
		}
		total := 0.0
		if len(cloNilai) > 0 {
			total = math.Average(cloNilai)
		}
		ploWithNilai = append(ploWithNilai, model.PLOWithNilai{
			ID:    v.ID,
			Nama:  v.Nama,
			Nilai: formatAvg(total),
		})
	}

	c.JSON(http.StatusOK, ploWithNilai)
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
		Nilai   []model.NilaiMahasiswa `json:"nilai" binding:"required"`
		MkId    int                    `json:"mk_id" binding:"required"`
		KelasId int                    `json:"kelas_id" binding:"required"`
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if body.Nilai == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nilai can't be empty"})
		return
	}
	if body.MkId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "mk_id can't be empty"})
		return
	}
	if body.KelasId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "kelas_id can't be empty"})
		return
	}

	nilai, err := json.Marshal(body.Nilai)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var penilaian model.Penilaian
	penilaian.Nilai = string(nilai)
	penilaian.Status = "draft"
	penilaian.MkId = body.MkId
	penilaian.KelasId = body.KelasId

	if err := p.penilaianRepo.CreatePenilaian(penilaian); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Penilaian added successfully"})
}

func (p *penilaianController) UpdatePenilaian(c *gin.Context) {
	var body struct {
		Nilai   []model.NilaiMahasiswa `json:"nilai" binding:"required"`
		MkId    int                    `json:"mk_id" binding:"required"`
		KelasId int                    `json:"kelas_id" binding:"required"`
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
	if body.Nilai == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nilai can't be empty"})
		return
	}
	if body.MkId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "mk_id can't be empty"})
		return
	}
	if body.KelasId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "kelas_id can't be empty"})
		return
	}

	nilai, err := json.Marshal(body.Nilai)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var penilaian model.Penilaian
	penilaian.Nilai = string(nilai)
	penilaian.Status = "draft"
	penilaian.MkId = body.MkId
	penilaian.KelasId = body.KelasId
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
